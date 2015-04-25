package controllers

import (
	"fmt"
	"net/http"

	"github.com/zenazn/goji/web"
)

func Home(c web.C, w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Howdy!")
}

func Hello(c web.C, w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, %s!", c.URLParams["name"])
}

func Accession(c web.C, w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Got accession: %s!", c.URLParams["id"])
}
