package main

import (
	"github.com/gobuffalo/packr/v2"
	spafs "github.com/koron/go-spafs"
	"gopkg.in/macaron.v1"
	"net/http"
	"os"
)

func main() {

	os.Setenv("MACARON_ENV", "production")
	dist := packr.New("static", "apps/admin/dist/admin")
	m := macaron.Classic()

	m.Use(macaron.Renderer())
	m.Get("/api", func() string {
		return "hello api"
	})

	go func() {
		http.ListenAndServe(":3000", spafs.FileServer(dist))
	}()
	m.Run()
}
