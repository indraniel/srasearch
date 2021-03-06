package routes

import (
	"github.com/indraniel/srasearch/web/controllers"

	"github.com/zenazn/goji"
)

func Include() {
	goji.Get("/", controllers.Home)
	goji.Get("/search", controllers.Search)
	goji.Get("/hello/:name", controllers.Hello)
	goji.Get("/accession/:id", controllers.Accession)
	goji.Get("/examples", controllers.Examples)
	goji.Get("/uploads", controllers.Uploads)
	goji.Get("/about", controllers.About)
	goji.NotFound(controllers.NotFound)
}
