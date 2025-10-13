//go:build dev

package main

import "github.com/canpacis/pacis/server"

func init() {
	AppEnv = server.Dev
	AppPort = ":8081"
	DevServer = "http://localhost:5173"
}

// Noop
func Ready(*server.App) {}
