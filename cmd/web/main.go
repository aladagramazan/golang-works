package main

import (
	"fmt"
	"log"
	"myproject/pkg/config"
	"myproject/pkg/handlers"
	"myproject/pkg/render"
	"net/http"
	"time"

	"github.com/alexedwards/scs/v2"
)

const portNumber = ":8080"

var app config.AppConfig
var session *scs.SessionManager

func main() {
	log.Println("Starting application...")

	// in production, change to true
	app.InProduction = false

	session = scs.New()
	session.Lifetime = 10 * time.Second
	session.Cookie.Persist = true
	session.Cookie.SameSite = http.SameSiteLaxMode
	session.Cookie.Secure = app.InProduction // for development, set to true in production

	app.Session = session

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
