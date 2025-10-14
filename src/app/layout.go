package app

import (
	"context"
	"os"

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

type Metadata struct {
	URL         string
	Title       string
	Description string
	Image       string
}

func RootLayout(server *server.Server, children Node) Node {
	specs := Speculation{
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

	base := os.Getenv("WEBSITE_URL")
	image := base + "/og-image.png"

	meta := Metadata{
		URL:         base,
		Title:       "The SSR Library for Go Developers.",
		Description: "Write type-safe templates with IntelliSense support, compose reusable components that make sense. When you're ready to ship, compile everything down to one binary and deploy anywhere. Performance is built-in, caching is opt-in, and streaming is automatic.",
		Image:       image,
	}

	return Fragment(
		Doctype,
		Head(
			Meta(Charset("UTF-8")),
			Meta(Name("viewport"), Content("width=device-width, initial-scale=1.0")),
			Link(Rel("icon"), Type("image/x-icon"), Href("/favicon.ico")),

			Title(Text("Welcome to Pacis")),
			Meta(Name("description"), Content(meta.Description)),
			Meta(Name("robots"), Content("index, follow")),

			// Open Graph
			Meta(PropertyAttr("og:title"), Content(meta.Title)),
			Meta(PropertyAttr("og:description"), Content(meta.Description)),
			Meta(PropertyAttr("og:url"), Content(meta.URL)),
			Meta(PropertyAttr("og:image"), Content(meta.Image)),

			// Twitter Cards
			Meta(Name("twitter:card"), Content("summary_large_image")),
			Meta(Name("twitter:title"), Content(meta.Title)),
			Meta(Name("twitter:description"), Content(meta.Description)),
			Meta(Name("twitter:image"), Content(meta.Image)),

			Link(Rel("stylesheet"), Href(server.Asset("/src/web/style.css"))),
			Script(Type("module"), Src(server.Asset("/src/web/main.ts"))),
			Script(Type("speculationrules"), JSON(specs)),
			Script(Defer, Src("https://analytics.formkitt.com/script.js"), Data("website-id", "78b36665-a529-4cb2-9297-938657f8b692")),
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
