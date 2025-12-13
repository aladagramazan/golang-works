package main

import (
	"errors"
	"fmt"
	"net/http"
)

const portNumber = ":8080"

func Home(w http.ResponseWriter, r *http.Request) {
	_, err := fmt.Fprintf(w, "This is the home page!")
	if err != nil {
		return
	}
}

func About(w http.ResponseWriter, r *http.Request) {
	sum := addValue(5, 10)
	_, _ = fmt.Fprintf(w, "This is the about page! Sum: %d", sum)
}

func addValue(x, y int) int {
	return x + y
}

func Divide(w http.ResponseWriter, r *http.Request) {
	f, err := divideValues(100.0, 20.0)
	if err != nil {
		_, _ = fmt.Fprintf(w, "cannot divide by zero ")
		return
	}

	_, _ = fmt.Fprintf(w, fmt.Sprintf("%f divide by %f is %f", 100.0, 10.0, f))
}

func divideValues(x, y float32) (float32, error) {
	if y == 0 {
		return 0, errors.New("division by zero is not allowed")
	}
	result := x / y
	return result, nil
}

func main() {

	http.HandleFunc("/", Home)
	http.HandleFunc("/about", About)
	http.HandleFunc("/divide", Divide)

	fmt.Println("Server is running on port 8080")
	err := http.ListenAndServe(portNumber, nil)
	if err != nil {
		fmt.Println("Error starting server:", err)
	}
}
