package render

import (
	"bytes"
	"html/template"
	"log"
	"net/http"
	"path/filepath"
)

// InitTemplateCache initializes the template cache on application startup
func InitTemplateCache() (map[string]*template.Template, error) {
	cache, err := createTemplateCache()
	if err != nil {
		return nil, err
	}
	log.Printf("Template cache initialized with %d templates", len(cache))
	return cache, nil
}

// Template renders a template from the given cache or reloads it if UseCache is false
func Template(w http.ResponseWriter, tmplName string, cache map[string]*template.Template, useCache bool) {
	var t *template.Template
	var err error

	if useCache {
		// Production mode: use cache
		var ok bool
		t, ok = cache[tmplName]
		if !ok {
			log.Printf("Could not get template '%s' from template cache", tmplName)
			http.Error(w, "Template not found", http.StatusInternalServerError)
			return
		}
	} else {
		// Development mode: reload template from disk every time
		log.Printf("Development mode: reloading template '%s' from disk", tmplName)
		t, err = createSingleTemplate(tmplName)
		if err != nil {
			log.Printf("Error loading template '%s': %v", tmplName, err)
			http.Error(w, "Error loading template", http.StatusInternalServerError)
			return
		}
	}

	buf := new(bytes.Buffer)
	err = t.Execute(buf, nil)
	if err != nil {
		log.Println("Error executing template:", err)
		http.Error(w, "Error rendering template", http.StatusInternalServerError)
		return
	}
	// render the template
	_, err = buf.WriteTo(w)
	if err != nil {
		log.Println("Error writing template to browser:", err)
	}
}

func createTemplateCache() (map[string]*template.Template, error) {
	myCache := map[string]*template.Template{}

	// get all the files named *.page.tmpl from ./templates
	pages, err := filepath.Glob("./templates/*.page.tmpl")
	if err != nil {
		return myCache, err
	}
	// range through all files ending with *.page.tmpl
	for _, page := range pages {
		name := filepath.Base(page)
		ts, err := template.New(name).ParseFiles(page)
		if err != nil {
			return myCache, err
		}

		matches, err := filepath.Glob("./templates/*.layout.tmpl")
		if err != nil {
			return myCache, err
		}

		if len(matches) > 0 {
			ts, err = ts.ParseGlob("./templates/*.layout.tmpl")
			if err != nil {
				return myCache, err
			}
		}

		myCache[name] = ts
	}
	return myCache, nil
}

// createSingleTemplate loads a single template from disk (for development mode)
func createSingleTemplate(tmplName string) (*template.Template, error) {
	// Parse the page template
	tmplPath := filepath.Join("./templates", tmplName)
	ts, err := template.New(tmplName).ParseFiles(tmplPath)
	if err != nil {
		return nil, err
	}

	// Check for layout templates
	matches, err := filepath.Glob("./templates/*.layout.tmpl")
	if err != nil {
		return nil, err
	}

	// If layouts exist, parse them too
	if len(matches) > 0 {
		ts, err = ts.ParseGlob("./templates/*.layout.tmpl")
		if err != nil {
			return nil, err
		}
	}

	return ts, nil
}
