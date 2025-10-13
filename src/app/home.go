package app

import (
	"net/http"

	. "github.com/canpacis/pacis/html"
	"github.com/canpacis/pacis/server"
	"github.com/canpacis/pacis/server/middleware"
)

// Implements server.Route
type Home struct{}

func (*Home) Path() string {
	return "/"
}

func (*Home) Handler(app *server.App) http.Handler {
	// Build page with its layout & middlewares
	return server.HandlerOf(app, HomePage, RootLayout, middleware.DefaultGzip)
}

func HomePage() Node {
	return Div(
		Class("min-h-screen flex flex-col items-center justify-center"),

		H1(
			Class("text-4xl font-bold mb-4"),

			Text("Build With Pacis"),
		),
		P(
			Class("flex items-center gap-2"),

			Text("Update the "),
			Span(
				Class("bg-neutral-200 text-neutral-700 pointer-events-none inline-flex h-5 w-fit min-w-5 items-center justify-center gap-1 rounded-sm px-1 font-sans text-xs font-medium select-none [&_svg:not([class*='size-'])]:size-3 [[data-slot=tooltip-content]_&]:bg-background/20 [[data-slot=tooltip-content]_&]:text-background dark:[[data-slot=tooltip-content]_&]:bg-background/10"),

				Text("src/app/home.go"),
			),
			Text(" file to see changes."),
		),
		A(
			Target("_blank"),
			Href("https://pacis.canpacis.com"),
			Class("underline text-sky-400 text-sm mt-4"),

			Text("Learn More"),
		),
	)
}
