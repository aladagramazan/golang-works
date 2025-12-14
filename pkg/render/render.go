package render

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

var templateCache = make(map[string]*template.Template)

// RenderTemplate renders a template from cache or creates it if not exists
func RenderTemplate(w http.ResponseWriter, tmplName string) {
	// Check if template exists in cache
	tmpl, exists := templateCache[tmplName]

	if !exists {
		// Template not in cache, create it
		log.Println("Template not in cache, creating:", tmplName)

		if err := createTemplateCache(tmplName); err != nil {
			log.Printf("Error creating template cache for %s: %v", tmplName, err)
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}

		// Retrieve the newly cached template
		tmpl = templateCache[tmplName]
	}

	// Execute the template
	if err := tmpl.ExecuteTemplate(w, "base", nil); err != nil {
		log.Printf("Error executing template %s: %v", tmplName, err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
}

// createTemplateCache parses template files and stores them in cache
func createTemplateCache(tmplName string) error {
	// Define template files to parse
	templateFiles := []string{
		fmt.Sprintf("./templates/%s", tmplName),
		"./templates/base.layout.tmpl",
	}

	// Parse template files
	tmpl, err := template.ParseFiles(templateFiles...)
	if err != nil {
		return fmt.Errorf("failed to parse template files: %w", err)
	}

	// Store in cache
	templateCache[tmplName] = tmpl
	log.Println("Template cached successfully:", tmplName)

	return nil
}
