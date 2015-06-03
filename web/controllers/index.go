package controllers

import (
	"github.com/indraniel/srasearch/searchdb"
	"github.com/indraniel/srasearch/utils"
	"github.com/indraniel/srasearch/web/render"

	"github.com/zenazn/goji/web"

	"encoding/json"
	"fmt"
	"net/http"
	"path/filepath"
	"strconv"
)

type Settings struct {
	Debug     bool
	IndexPath string
}

var setup Settings

func Init(debug bool, indexPath string) {
	setup.Debug = debug
	setup.IndexPath = indexPath
}

func Hello(c web.C, w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, %s!", c.URLParams["name"])
}

func Home(c web.C, w http.ResponseWriter, r *http.Request) {
	q := r.URL.Query()
	if _, exists := q["q"]; exists {
		url := fmt.Sprintf("%s?%s", "/search", r.URL.RawQuery)
		http.Redirect(w, r, url, http.StatusSeeOther)
	}

	templates := render.BaseTemplates()
	templates = append(templates, "web/views/index.html")

	data := make(map[string]string)
	data["Title"] = "Beaker: A NCBI SRA Search Assistant"

	err := render.RenderHTML(w, templates, "base", data)
	if err != nil {
		render.RenderError(w, err, http.StatusInternalServerError)
	}
}

func Search(c web.C, w http.ResponseWriter, r *http.Request) {
	q := r.URL.Query()
	query, exists := q["q"]
	if exists == false || query[0] == "" {
		url := "/"
		http.Redirect(w, r, url, http.StatusSeeOther)
	}

	start, exists := q["start"]
	if exists == false || start[0] == "" {
		err := fmt.Errorf("Please provide a valid start date!")
		render.RenderError(w, err, http.StatusInternalServerError)
		return
	}

	end, exists := q["end"]
	if exists == false || end[0] == "" {
		err := fmt.Errorf("Please provide a valid end date!")
		render.RenderError(w, err, http.StatusInternalServerError)
		return
	}

	// default query size
	querySize := 25

	pageNum := 1
	if page, ok := q["page"]; ok {
		pageNum, _ = strconv.Atoi(page[0])
	}


	searchResults, err := searchdb.Query(
		query[0],
		start[0],
		end[0],
		pageNum,
		querySize,
	)

	if err != nil {
		render.RenderError(w, err, http.StatusInternalServerError)
		return
	}
	jsonStr, _ := json.MarshalIndent(searchResults, "", "    ")
	pagination := NewPagination(pageNum, querySize, int(searchResults.Total))

	templates := render.BaseTemplates()
	templates = append(templates, "web/views/search.html")

	data := make(map[string]interface{})
	data["Title"] = "Beaker Search"
	data["Query"] = query[0]
	data["JsonStr"] = string(jsonStr)
	data["searchResults"] = searchResults
	data["pagination"] = pagination
	data["Debug"] = setup.Debug
	data["Start"] = start[0]
	data["End"] = end[0]

	err = render.RenderHTML(w, templates, "base", data)
	if err != nil {
		render.RenderError(w, err, http.StatusInternalServerError)
	}
}

func About(c web.C, w http.ResponseWriter, r *http.Request) {
	templates := render.BaseTemplates()
	templates = append(templates, "web/views/about.html")

	pattern := "index_meta.json"
	fullPattern := filepath.Join(setup.IndexPath, pattern)
	file, err := utils.FindFile(fullPattern)
	if err != nil {
		render.RenderError(w, err, http.StatusNotFound)
		return
	}
	mtime, err := utils.FileModificationTime(file)
	if err != nil {
		render.RenderError(w, err, http.StatusNotFound)
		return
	}

	upFile, err := getUploadsFile()
	if err != nil {
		upFile = "-"
	}

	data := make(map[string]interface{})
	data["Title"] = fmt.Sprintf("Beaker: Examples & Documentation")
	data["IndexPath"] = setup.IndexPath
	data["CreationTime"] = mtime
	data["UploadsFile"] = upFile

	err = render.RenderHTML(w, templates, "base", data)
	if err != nil {
		render.RenderError(w, err, http.StatusInternalServerError)
	}
}

func Examples(c web.C, w http.ResponseWriter, r *http.Request) {
	templates := render.BaseTemplates()
	templates = append(templates, "web/views/examples.html")

	data := make(map[string]interface{})
	data["Title"] = fmt.Sprintf("Beaker: Examples & Documentation")

	err := render.RenderHTML(w, templates, "base", data)
	if err != nil {
		render.RenderError(w, err, http.StatusInternalServerError)
	}
}

func Uploads(c web.C, w http.ResponseWriter, r *http.Request) {
	file, err := getUploadsFile()
	if err != nil {
		render.RenderError(w, err, http.StatusNotFound)
		return
	}

	basename := filepath.Base(file)
	contentDisposition := fmt.Sprintf("inline; filename=\"%s\"", basename)
	w.Header().Set("Content-Type", "text/tab-separated-values")
	w.Header().Set("Content-Disposition", contentDisposition)
	http.ServeFile(w, r, file)
}

func Accession(c web.C, w http.ResponseWriter, r *http.Request) {
	accession := c.URLParams["id"]

	q := r.URL.Query()

	start, sOk := q["start"]
	end, eOk := q["end"]

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

	if sOk && eOk && start[0] != "" && end[0] != "" {
		data["Start"] = start[0]
		data["End"] = end[0]
	}

	err = render.RenderHTML(w, templates, "base", data)
	if err != nil {
		render.RenderError(w, err, http.StatusInternalServerError)
	}
}

func NotFound(c web.C, w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(404)
	render.RenderNotFound(w, r)
}

func getUploadsFile() (string, error) {
	pattern := "recent-*-sra-uploads-*.tsv"
	fullPattern := filepath.Join(setup.IndexPath, pattern)
	file, err := utils.FindFile(fullPattern)
	return file, err
}
