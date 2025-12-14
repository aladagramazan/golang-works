package main

import (
	"fmt"
	"log"
	"myproject/pkg/handler"
	"net/http"
)

const portNumber = ":8080"

func main() {
	log.Println("Starting application...")

	http.HandleFunc("/", handler.Home)
	http.HandleFunc("/about", handler.About)

	log.Println("Handlers registered")
	fmt.Printf("Server is running on http://localhost%s\n", portNumber)

	err := http.ListenAndServe(portNumber, nil)
	if err != nil {
		log.Fatal("Error starting server:", err)
	}
}
