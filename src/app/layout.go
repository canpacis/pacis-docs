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

func RootLayout(server *server.Server, children Node) Node {
	specs := &Speculation{
		Prerender: []SpeculationRule{
			{
				URLs:      []string{"/getting-started/introduction/", "/getting-started/installation/", "/getting-started/quick-start/"},
				Eagerness: "immediate",
			},
			{
				URLs:      []string{"/getting-started/templating/", "/getting-started/conventions/"},
				Eagerness: "eager",
			},
			{
				URLs: []string{
					"/core-concepts/rendering/",
					"/core-concepts/request-data/",
					"/core-concepts/streaming/",
					"/core-concepts/middlewares/",
					"/core-concepts/caching/",
					"/core-concepts/assets/",
					"/core-concepts/deploying/",
				},
				Eagerness: "conservative",
			},
		},
	}

	return Fragment(
		Doctype,
		Head(
			Meta(Charset("UTF-8")),
			Meta(Name("viewport"), Content("width=device-width, initial-scale=1.0")),
			Title(Text("Welcome to Pacis")),
			Link(Rel("icon"), Type("image/x-icon"), Href("/favicon.ico")),

			Link(Rel("stylesheet"), Href(server.Asset("/src/web/style.css"))),
			Script(Type("module"), Src(server.Asset("/src/web/main.ts"))),
			Script(Type("speculationrules"), JSON(specs)),
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
