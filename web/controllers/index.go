package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/zenazn/goji/web"

	"github.com/indraniel/srasearch/searchdb"
	"github.com/indraniel/srasearch/web/render"
)

func Home(c web.C, w http.ResponseWriter, r *http.Request) {
	q := r.URL.Query()
	if _, exists := q["q"]; exists {
		url := fmt.Sprintf("%s?%s", "/search", r.URL.RawQuery)
		http.Redirect(w, r, url, http.StatusSeeOther)
	}

	templates := render.BaseTemplates()
	templates = append(templates, "web/views/index.html")

	data := make(map[string]string)
	data["Title"] = "Home"

	err := render.RenderHTML(w, templates, "base", data)
	if err != nil {
		render.RenderError(w, err, http.StatusInternalServerError)
	}
}

func Search(c web.C, w http.ResponseWriter, r *http.Request) {
	q := r.URL.Query()
	term, exists := q["q"]
	if exists == false || term[0] == "" {
		url := "/"
		http.Redirect(w, r, url, http.StatusSeeOther)
	}

	pageNum := 1
	if page, ok := q["page"]; ok {
		pageNum, _ = strconv.Atoi(page[0])
	}

	searchResults, err := searchdb.Query(term[0], pageNum)
	if err != nil {
		render.RenderError(w, err, http.StatusInternalServerError)
		return
	}
	jsonStr, _ := json.MarshalIndent(searchResults, "", "    ")

	templates := render.BaseTemplates()
	templates = append(templates, "web/views/search.html")

	data := make(map[string]interface{})
	data["Title"] = "Search"
	data["Query"] = term[0]
	data["JsonStr"] = string(jsonStr)
	data["searchResults"] = searchResults

	err = render.RenderHTML(w, templates, "base", data)
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
	accession := c.URLParams["id"]

	si, err := searchdb.GetSRAItem(accession)
	if err != nil {
		render.RenderError(w, err, http.StatusNotFound)
		return
	}

	templates := render.BaseTemplates()
	templates = append(templates, "web/views/accession.html")

	data := make(map[string]interface{})
	data["Title"] = fmt.Sprintf("%s : %s", "Accession", accession)
	data["SRAItem"] = si

	err = render.RenderHTML(w, templates, "base", data)
	if err != nil {
		render.RenderError(w, err, http.StatusInternalServerError)
	}
}

func NotFound(c web.C, w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(404)
	render.RenderNotFound(w, r)
}
