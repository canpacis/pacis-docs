package main

import (
	"net/http"

	"github.com/canpacis/pacis-docs/src/app"
	"github.com/canpacis/pacis/server"
	"github.com/canpacis/pacis/server/middleware"
)

var options *server.Options

func main() {
	srv := server.New(options)

	Ready(srv)

	srv.Use(middleware.DefaultColorScheme)

	gzip := middleware.DefaultGzip(srv.Env() == server.Dev)
	srv.HandlePage("/", &app.HomePage{}, app.RootLayout, gzip)
	for _, doc := range app.Docs {
		if len(doc.Links) == 0 {
			continue
		}
		srv.Handle("GET "+doc.Href, http.RedirectHandler(doc.Links[0].Href, http.StatusFound))
		for _, link := range doc.Links {
			srv.HandlePage(link.Href, app.NewDocPage(options.Env, doc.Folder, link.File), app.DocsLayout, gzip)
		}
	}
	srv.Handle("GET /docs", http.RedirectHandler("/getting-started", http.StatusFound))
	srv.Handle("GET /docs/", http.RedirectHandler("/getting-started", http.StatusFound))
	srv.Handle("GET /docs/{title}/", http.RedirectHandler("/getting-started", http.StatusFound))
	srv.Handle("GET /docs/{title}/{doc}", http.RedirectHandler("/getting-started", http.StatusFound))
	srv.Handle("GET /docs/{title}/{doc}/", http.RedirectHandler("/getting-started", http.StatusFound))

	srv.Serve()
}
