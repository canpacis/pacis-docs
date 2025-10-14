//go:build dev

package main

import (
	"github.com/canpacis/pacis/server"
	"github.com/canpacis/pacis/server/middleware"
)

func init() {
	AppEnv = server.Dev
	AppPort = ":8081"
	DevServer = "http://localhost:5173"
}

func Ready(app *server.App) {
	app.Use(middleware.DefaultColorScheme)
}
