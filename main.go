package main

import (
	"fmt"
	"net/http"
)

const port = ":8080"

func Home(response http.ResponseWriter, request *http.Request) {
	fmt.Fprint(response, "Hello world !")
}

func About(response http.ResponseWriter, request *http.Request) {
	fmt.Fprint(response, "À propos de moi")
}

func main() {
	http.HandleFunc("/", Home)
	http.HandleFunc("/about", About)

	fmt.Println("Server start on http://localhost:8080")

	http.ListenAndServe(port, nil)
}
