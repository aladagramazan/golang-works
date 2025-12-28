package main

import (
	"myproject/pkg/config"
	"myproject/pkg/handlers"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func routes(app *config.AppConfig) http.Handler {
	mux := chi.NewRouter()

	mux.Use(middleware.Recoverer) // handlerda bir panic oluşursa uygulamanın çökmesini engeller
	// panici yakalar ve 500 Internal Server Error döner
	// panic detaylarını loglar
	// sunucunun çalışmaya devam etmesini sağlar

	mux.Use(NoSurf) // CSRF koruması ekler

	mux.Get("/", handlers.Repo.Home)
	mux.Get("/about", handlers.Repo.About)

	return mux
}
