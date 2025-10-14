//go:build dev

package main

import (
	"log/slog"
	"net/url"

	"github.com/canpacis/pacis/server"
	"github.com/joho/godotenv"
)

func init() {
	godotenv.Load()
	dev, _ := url.Parse("http://localhost:5173")
	options = &server.Options{
		Env:       server.Dev,
		Port:      ":8081",
		DevServer: dev,
		Logger:    slog.Default(),
	}
}

func Ready(s *server.Server) {
	s.RegisterDevHandlers()
}
