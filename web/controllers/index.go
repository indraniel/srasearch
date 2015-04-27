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
		render.RenderError(w, err, http.StatusInternalServerError)
	}
}

func Results(c web.C, w http.ResponseWriter, r *http.Request) {
	templates := render.BaseTemplates()
	templates = append(templates, "web/views/results.html")

	data := make(map[string]string)
	data["Title"] = "Example Results"

	err := render.RenderHTML(w, templates, "base", data)
	if err != nil {
		render.RenderError(w, err, http.StatusInternalServerError)
	}
}

func Hello(c web.C, w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, %s!", c.URLParams["name"])
}

func Accession(c web.C, w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Got accession: %s!", c.URLParams["id"])
}

func NotFound(c web.C, w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(404)
	render.RenderNotFound(w, r)
}
