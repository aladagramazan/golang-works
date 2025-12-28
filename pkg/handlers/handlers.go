package handlers

import (
	"myproject/pkg/config"
	"myproject/pkg/models"
	"myproject/pkg/render"
	"net/http"
)

var Repo *Repository

type Repository struct {
	App *config.AppConfig
}

// NewRepository creates a new repository
func NewRepository(a *config.AppConfig) *Repository {
	return &Repository{
		App: a,
	}
}

func NewHandlers(r *Repository) {
	Repo = r
}
func (repo *Repository) Home(w http.ResponseWriter, r *http.Request) {
	remoteIP := r.RemoteAddr
	repo.App.Session.Put(r.Context(), "remote_ip", remoteIP)
	render.Template(w, "home.page.tmpl", &models.TemplateData{})
}

// About is the about page handler
func (repo *Repository) About(w http.ResponseWriter, r *http.Request) {
	stringMap := make(map[string]string)
	stringMap["test"] = "Hello, again."

	remoteIP := repo.App.Session.GetString(r.Context(), "remote_ip")
	stringMap["remote_ip"] = remoteIP

	render.Template(w, "about.page.tmpl", &models.TemplateData{
		StringMap: stringMap,
	})
}
