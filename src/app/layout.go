package app

import (
	"context"

	. "github.com/canpacis/pacis/html"
	"github.com/canpacis/pacis/server"
	"github.com/canpacis/pacis/server/font"
	"github.com/canpacis/pacis/server/middleware"
)

func RootLayout(server *server.Server, children Node) Node {
	return Fragment(
		Doctype,
		Head(
			Meta(Charset("UTF-8")),
			Meta(Name("viewport"), Content("width=device-width, initial-scale=1.0")),
			Title(Text("Welcome to Pacis")),
			Link(Rel("icon"), Type("image/x-icon"), Href("/favicon.ico")),

			Link(Rel("stylesheet"), Href(server.Asset("/src/web/style.css"))),
			Script(Type("module"), Src(server.Asset("/src/web/main.ts"))),
			server.HMR(),
			font.Head(
				font.New("Inter", font.WeightList{font.W100, font.W900}, font.Auto, font.Latin, font.LatinExt),
			),
		),
		Body(
			DeferredAttr("class", func(ctx context.Context) string {
				return middleware.GetColorScheme(ctx)
			}),

			children,
			Script(Src(server.Asset("/src/web/stream.ts"))),
		),
	)
}
