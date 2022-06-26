package main

import (
	"html/template"
	"net/http"
)

func renderTemplate(response http.ResponseWriter, tmpl string) {
	t, err := template.ParseFiles("./templates/" + tmpl + ".page.tmpl")
	if err != nil {
		http.Error(response, err.Error(), http.StatusInternalServerError)
		return
	}
	t.Execute(response, nil)
}

func Home(response http.ResponseWriter, request *http.Request) {
	renderTemplate(response, "home")
}

func About(response http.ResponseWriter, request *http.Request) {
	renderTemplate(response, "about")
}
