package app

import (
	"context"

	. "github.com/canpacis/pacis/html"
	"github.com/canpacis/pacis/server"
	"github.com/canpacis/pacis/server/font"
	"github.com/canpacis/pacis/server/middleware"
)

type SpeculationRule struct {
	URLs      []string `json:"urls"`
	Eagerness string   `json:"eagerness"`
}

type Speculation struct {
	Prefetch  []SpeculationRule `json:"prefetch,omitempty"`
	Prerender []SpeculationRule `json:"prerender,omitempty"`
}

func RootLayout(s *server.Server, head, children Node) Node {
	specs := Speculation{
		// Prerender: []SpeculationRule{
		// 	{
		// 		URLs:      []string{"/getting-started/introduction/", "/getting-started/installation/", "/getting-started/quick-start/"},
		// 		Eagerness: "immediate",
		// 	},
		// 	{
		// 		URLs:      []string{"/getting-started/templating/", "/getting-started/conventions/"},
		// 		Eagerness: "eager",
		// 	},
		// },
	}

	return Fragment(
		Doctype,
		Html(
			Lang("en"),

			Head(
				Meta(Charset("UTF-8")),
				Meta(Name("viewport"), Content("width=device-width, initial-scale=1.0")),
				Link(Rel("icon"), Type("image/x-icon"), Href("/favicon.ico")),

				Link(Rel("stylesheet"), Href(s.Asset("/src/web/style.css"))),
				Script(Type("module"), Src(s.Asset("/src/web/main.ts"))),
				Script(Type("speculationrules"), JSON(specs)),
				Script(Defer, Src("https://analytics.formkitt.com/script.js"), Data("website-id", "78b36665-a529-4cb2-9297-938657f8b692")),
				head,
				font.Head(
					font.New("Inter", font.WeightList{font.W100, font.W900}, font.Auto, font.Latin, font.LatinExt),
					font.New("JetBrains+Mono", font.WeightList{font.W100, font.W800}, font.Auto, font.Latin, font.LatinExt),
				),
			),
			Body(
				DeferredAttr("class", func(ctx context.Context) string {
					return middleware.GetColorScheme(ctx) + " w-screen overflow-x-hidden"
				}),

				children,
				Script(Src(s.Asset("/src/web/stream.ts"))),
			),
		),
	)
}
