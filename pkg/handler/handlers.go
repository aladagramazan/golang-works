package handler

import (
	"log"
	"myproject/pkg/config"
	"myproject/pkg/render"
	"net/http"
)

// Repository holds the application config
type Repository struct {
	App *config.AppConfig
}

// NewRepository creates a new repository
func NewRepository(a *config.AppConfig) *Repository {
	return &Repository{
		App: a,
	}
}

// Home is the home page handler
func (repo *Repository) Home(w http.ResponseWriter, r *http.Request) {
	log.Printf("Home handler called - Method: %s, URL: %s", r.Method, r.URL.Path)
	render.Template(w, "home.page.tmpl", repo.App.TemplateCache, repo.App.UseCache)
}

// About is the about page handler
func (repo *Repository) About(w http.ResponseWriter, r *http.Request) {
	log.Printf("About handler called - Method: %s, URL: %s", r.Method, r.URL.Path)
	render.Template(w, "about.page.tmpl", repo.App.TemplateCache, repo.App.UseCache)
}
