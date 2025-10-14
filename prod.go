//go:build prod

package main

import (
	"embed"
	"log"

	"github.com/canpacis/pacis/server"
	"github.com/canpacis/pacis/server/middleware"
)

var (
	//go:embed build/*
	build embed.FS
)

func init() {
	AppEnv = server.Prod
	AppPort = ":8080"
}

func Ready(app *server.App) {
	if err := app.SetBuildDir("build", build); err != nil {
		log.Fatal(err)
	}
	app.Use(middleware.DefaultColorScheme)
}
