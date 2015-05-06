package render

import (
	"html/template"
	"net/http"
	"strconv"
)

func RenderHTML(
	w http.ResponseWriter,
	templates []string,
	name string,
	data interface{}) error {
	t := template.New(name).Funcs(funcMap)

	t, err := t.ParseFiles(templates...)
	if err != nil {
		return err
	}

	if err = t.ExecuteTemplate(w, name, data); err != nil {
		return err
	}

	return nil
}

func BaseTemplates() []string {
	templates := []string{
		"web/views/base/footer.html",
		"web/views/base/header.html",
		"web/views/base/navbar.html",
		"web/views/base/base.html",
	}

	return templates
}

func RenderError(w http.ResponseWriter, e error, status int) {
	w.WriteHeader(status)

	templates := BaseTemplates()
	templates = append(templates, "web/views/error.html")

	data := make(map[string]string)
	data["Title"] = "Error"
	data["Code"] = strconv.Itoa(status)
	data["Error"] = e.Error()

	err := RenderHTML(w, templates, "base", data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func RenderNotFound(w http.ResponseWriter, r *http.Request) {
	templates := BaseTemplates()
	templates = append(templates, "web/views/404.html")

	data := make(map[string]string)
	data["Title"] = "Not Found"
	data["Path"] = r.URL.Path

	err := RenderHTML(w, templates, "base", data)
	if err != nil {
		RenderError(w, err, http.StatusInternalServerError)
	}
}
