package main

import (
	"fmt"
	"net/http"

	"github.com/felixbouazza/go-web-app/internal/handlers"
)

const port = ":8080"

func main() {
	http.HandleFunc("/", handlers.Home)
	http.HandleFunc("/about", handlers.About)

	fmt.Println("Server start on http://localhost:8080")

	http.ListenAndServe(port, nil)
}
