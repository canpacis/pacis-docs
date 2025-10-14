package main

import (
	"net/http"

	"github.com/canpacis/pacis-app/src/app"
	"github.com/canpacis/pacis/server"
	"github.com/canpacis/pacis/server/middleware"
)

var options *server.Options

func main() {
	server := server.New(options)

	Ready(server)

	server.Use(middleware.DefaultColorScheme)

	server.HandlePage("GET /", app.Home, app.RootLayout, middleware.DefaultGzip)
	for _, doc := range app.Docs {
		for _, link := range doc.Links {
			server.HandlePage("GET "+link.Href, app.DocPage(options.Env, link.File), app.DocsLayout)
		}
	}
	server.Handle("GET /getting-started", http.RedirectHandler("/getting-started/introduction", http.StatusFound))
	server.Handle("GET /getting-started/", http.RedirectHandler("/getting-started/introduction", http.StatusFound))
	server.Handle("GET /core-concepts/", http.RedirectHandler("/core-concepts/components", http.StatusFound))
	server.Handle("GET /core-concepts", http.RedirectHandler("/core-concepts/components", http.StatusFound))

	server.Serve()
}
