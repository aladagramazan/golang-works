package main

import (
	"fmt"
	"log"
	"myproject/pkg/config"
	"myproject/pkg/handlers"
	"myproject/pkg/render"
	"net/http"
)

const portNumber = ":8080"

func main() {
	log.Println("Starting application...")
	var app config.AppConfig

	// Initialize template cache once at startup
	tc, err := render.InitTemplateCache()
	if err != nil {
		log.Fatal("Cannot create template cache:", err)
	}
	app.TemplateCache = tc
	app.UseCache = true // for development, set to true in production

	// Create handler repository with dependency injection
	repo := handlers.NewRepository(&app)

	handlers.NewHandlers(repo)
	render.SetAppConfig(&app)

	log.Println("Handlers registered")
	fmt.Printf("Server is running on http://localhost%s\n", portNumber)

	srv := &http.Server{
		Addr:    portNumber,
		Handler: routes(&app),
	}
	err = srv.ListenAndServe()
	if err != nil {
		log.Fatal("Error starting server:", err)
	}
}
