package main

import (
	"net/http"

	"github.com/canpacis/pacis-app/src/app"
	"github.com/canpacis/pacis/server"
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
	// Serve static assets
	mux.Handle("GET /assets/", http.StripPrefix("/assets/", application.ServeAssets()))

	// Register routes
	application.Register(mux, &app.Home{})

	// Graceful server
	server.Serve(application, mux)
}
