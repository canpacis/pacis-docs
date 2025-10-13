package main

import (
	"net/http"

	"github.com/canpacis/pacis-app/src/app"
	"github.com/canpacis/pacis/server"
	"github.com/canpacis/pacis/server/middleware"
)

var (
	AppEnv    server.Environment
	AppPort   string
	DevServer string
)

func main() {
	application := server.NewApp(
		server.WithEnv(AppEnv),
		server.WithPort(AppPort),
		server.WithDevServer(DevServer),
	)
	Ready(application)

	mux := http.NewServeMux()
	mux.Handle("GET /assets/", http.StripPrefix("/assets/", application.ServeAssets()))

	application.Register(mux, server.RouteOf("/", app.Home, app.RootLayout, middleware.DefaultGzip))
	for _, doc := range app.Docs {
		for _, link := range doc.Links {
			application.Register(mux, server.RouteOf(link.Href, app.DocPage(link.File), app.DocsLayout))
		}
	}
	application.Register(mux, server.RedirectRoute("/getting-started", "/getting-started/introduction"))
	application.Register(mux, server.RedirectRoute("/core-concepts/", "/core-concepts/components"))

	server.Serve(application, mux)
}
