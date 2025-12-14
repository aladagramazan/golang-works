package handler

import (
	"log"
	"myproject/pkg/render"
	"net/http"
)

func Home(w http.ResponseWriter, r *http.Request) {
	log.Printf("Home handler called - Method: %s, URL: %s", r.Method, r.URL.Path)
	render.RenderTemplate(w, "home.page.tmpl")
}

func About(w http.ResponseWriter, r *http.Request) {
	log.Printf("About handler called - Method: %s, URL: %s", r.Method, r.URL.Path)
	render.RenderTemplate(w, "about.page.tmpl")
}
