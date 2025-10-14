//go:build prod

package main

import (
	"embed"
	"log"
	"log/slog"

	"github.com/canpacis/pacis-app/src/app"
	"github.com/canpacis/pacis/server"
)

func init() {
	options = &server.Options{
		Env:    server.Prod,
		Port:   ":8080",
		Logger: slog.Default(),
	}
}

var (
	//go:embed build/*
	build embed.FS
	//go:embed build/static/.vite
	vite embed.FS
	//go:embed src/app/docs
	docs embed.FS
)

func Ready(s *server.Server) {
	if err := s.SetBuildDir("build/static", build, vite); err != nil {
		log.Fatal(err)
	}
	app.SetDocsFS(docs)
}
