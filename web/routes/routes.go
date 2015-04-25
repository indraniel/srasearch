package routes

import (
	"github.com/indraniel/srasearch/web/controllers"

	"github.com/zenazn/goji"
)

func Include() {
	goji.Get("/", controllers.Home)
	goji.Get("/hello/:name", controllers.Hello)
	goji.Get("/accession/:id", controllers.Accession)
}
