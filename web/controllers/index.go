package controllers

import (
	"fmt"
	"net/http"

	"github.com/zenazn/goji/web"

	"github.com/indraniel/srasearch/web/render"
)

func Home(c web.C, w http.ResponseWriter, r *http.Request) {
	templates := render.BaseTemplates()
	templates = append(templates, "web/views/index.html")

	data := make(map[string]string)
	data["Title"] = "Home"

	err := render.RenderHTML(w, templates, "base", data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func Hello(c web.C, w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, %s!", c.URLParams["name"])
}

func Accession(c web.C, w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Got accession: %s!", c.URLParams["id"])
}
