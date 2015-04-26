package render

import (
	"html/template"
	"net/http"
)

func RenderHTML(
	w http.ResponseWriter,
	templates []string,
	name string,
	data interface{}) error {
	t := template.New(name)

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
