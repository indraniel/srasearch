package web

import (
	"fmt"
	"log"
	"net"
	"net/http"

	"github.com/zenazn/goji"
	"github.com/zenazn/goji/bind"
	"github.com/zenazn/goji/graceful"
	"github.com/zenazn/goji/web"
)

var indexPath string

func hello(c web.C, w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, %s!", c.URLParams["name"])
}

func Main(port int, host, idxPath string) {
	indexPath = idxPath
	goji.Get("/hello/:name", hello)
	Serve(host, port)
}

func Serve(host string, port int) {
	http.Handle("/", goji.DefaultMux)
	address := fmt.Sprintf("%s:%d", host, port)
	listener, err := net.Listen("tcp", address)
	if err != nil {
		log.Fatalln("Trouble starting up listener on ", address)
	}
	log.Println("Starting Goji on", listener.Addr())
	log.Println("Bleve IndexPath is:", indexPath)
	graceful.HandleSignals()
	bind.Ready()
	graceful.PreHook(func() { log.Printf("Goji received signal, gracefully stopping") })
	graceful.PostHook(func() { log.Printf("Goji stopped") })

	err = graceful.Serve(listener, http.DefaultServeMux)
	if err != nil {
		log.Fatal(err)
	}

	graceful.Wait()
}
