package main

import (
	"fmt"
	"log"
	"myproject/pkg/config"
	"myproject/pkg/handler"
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
	app.UseCache = false // for development, set to true in production

	// Create handler repository with dependency injection
	repo := handler.NewRepository(&app)

	// Register routes
	http.HandleFunc("/", repo.Home)
	http.HandleFunc("/about", repo.About)

	log.Println("Handlers registered")
	fmt.Printf("Server is running on http://localhost%s\n", portNumber)

	err = http.ListenAndServe(portNumber, nil)
	if err != nil {
		log.Fatal("Error starting server:", err)
	}
}
