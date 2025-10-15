package main

import (
	"net/http"

	"github.com/canpacis/pacis-docs/src/app"
	"github.com/canpacis/pacis/server"
	"github.com/canpacis/pacis/server/middleware"
)

var options *server.Options

func main() {
	server := server.New(options)

	Ready(server)

	server.Use(middleware.DefaultColorScheme)

	server.HandlePage("GET /", &app.HomePage{}, app.RootLayout)
	for _, doc := range app.Docs {
		server.Handle("GET "+doc.Href, http.RedirectHandler(doc.Links[0].Href, http.StatusFound))
		server.Handle("GET "+doc.Href+"/", http.RedirectHandler(doc.Links[0].Href, http.StatusFound))
		for _, link := range doc.Links {
			server.HandlePage("GET "+link.Href, app.NewDocPage(options.Env, link.File), app.DocsLayout)
		}
	}

	server.Serve()
}
