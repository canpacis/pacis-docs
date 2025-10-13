package app

import (
	. "github.com/canpacis/pacis/html"
	"github.com/canpacis/pacis/server"
	"github.com/canpacis/pacis/server/font"
)

func RootLayout(app *server.App, children Node) Node {
	return Fragment(
		Doctype,
		Head(
			Meta(Charset("UTF-8")),
			Meta(Name("viewport"), Content("width=device-width, initial-scale=1.0")),
			Title(Text("Welcome to Pacis")),

			Link(Rel("stylesheet"), Href(server.Asset(app, "style.css"))),
			Link(Rel("icon"), Type("image/webp"), Href(server.Asset(app, "static/favicon.webp"))),
			Script(Defer, Type("module"), Src(server.Asset(app, "main.ts"))),
			server.HMR(app),
			font.Head(
				font.New("Inter", font.WeightList{font.W100, font.W900}, font.Auto, font.Latin, font.LatinExt),
			),
		),
		Body(
			Class("dark"),

			children,
			Script(Src(server.Asset(app, "stream.ts"))),
		),
	)
}
