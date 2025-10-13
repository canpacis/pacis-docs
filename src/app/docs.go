package app

import (
	"context"
	"log"
	"os"

	. "github.com/canpacis/pacis/html"
	"github.com/canpacis/pacis/lucide"
	"github.com/canpacis/pacis/server"
	parser "github.com/sivukhin/godjot/djot_parser"
)

type DocLink struct {
	Label string
	Href  string
	File  string
}

type DocTitle struct {
	Label string
	Links []DocLink
}

var Docs = []DocTitle{
	{
		Label: "Getting Started",
		Links: []DocLink{
			{Label: "Introduction", Href: "/getting-started/introduction/", File: "introduction"},
			{Label: "Installation", Href: "/getting-started/installation/", File: "installation"},
			{Label: "Qucik Start", Href: "/getting-started/quick-start/", File: "quick-start"},
			{Label: "Tutorial", Href: "/getting-started/tutorial/", File: "tutorial"},
			{Label: "Conventions", Href: "/getting-started/conventions/", File: "conventions"},
		},
	},
	{
		Label: "Core Concepts",
		Links: []DocLink{
			{Label: "Components", Href: "/core-concepts/components/", File: "components"},
			{Label: "Templates", Href: "/core-concepts/templates/", File: "templates"},
			{Label: "Routing", Href: "/core-concepts/routing/", File: "routing"},
			{Label: "Streaming", Href: "/core-concepts/streaming/", File: "streaming"},
			{Label: "Caching", Href: "/core-concepts/caching/", File: "caching"},
			{Label: "Data Fetching", Href: "/core-concepts/data-fetching/", File: "data-fetching"},
		},
	},
}

func DocsLayout(app *server.App, children Node) Node {
	return RootLayout(app, Main(
		Class("min-h-screen flex flex-col"),

		Div(
			Class("flex max-w-6xl w-full mx-auto my-2"),

			Nav(
				Class("flex flex-col w-[200px] h-full p-4 gap-6"),

				A(
					Class("flex gap-2 items-center"),
					Href("/"),

					Img(Src(server.Asset(app, "static/logo.webp")), Class("h-7")),
					Span(Class("uppercase text-lg font-light select-none"), Text("Pacis")),
				),

				Map(Docs, func(title DocTitle) Node {
					return Div(
						Class("flex flex-col gap-2"),

						P(Class("font-semibold text-sm"), Text(title.Label)),
						Ul(
							Class("flex flex-col gap-0.5"),

							Map(title.Links, func(link DocLink) Node {
								return Li(
									A(
										Class("bg-background hover:bg-accent w-full h-7 flex text-sm rounded-md px-3 items-center data-[state=active]:bg-accent"),
										DeferredAttr("data-state", func(ctx context.Context) string {
											detail := server.Detail(ctx)
											if detail.URL.Path == link.Href {
												return "active"
											}
											return ""
										}),
										Href(link.Href),

										Text(link.Label),
									),
								)
							}),
						),
					)
				}),
			),
			Section(
				Class("flex-1 h-full py-4 pr-8"),

				children,
			),
		),
		Footer(
			Class("mt-auto flex w-full max-w-6xl mx-auto items-center justify-center py-3 border-t text-sm"),

			P(Class("text-muted-foreground"), Text("Made with Pacis")),
		),
	))
}

type GettingStartedPayload struct {
	Path string `path:"path"`
}

func DocPage(name string) func(*server.App) Node {
	file, err := os.ReadFile("src/app/docs/" + name + ".md")
	if err != nil {
		log.Fatal(err)
	}

	var prev *DocLink
	var next *DocLink

	for _, title := range Docs {
		for i, link := range title.Links {
			if link.File == name {
				if i > 0 {
					prev = &title.Links[i-1]
				}
				if i < len(title.Links)-1 {
					next = &title.Links[i+1]
				}
			}
		}
	}

	nodes := Frag{}
	parser.BuildDjotAst(file)[0].Traverse(func(node parser.TreeNode[parser.DjotNode]) {
		switch node.Type {
		case parser.HeadingNode:
			nodes = append(nodes, H1(Class("text-3xl font-semibold mb-2"), Text(node.FullText())))
		case parser.ParagraphNode:
			nodes = append(nodes, P(Text(node.FullText())))
		}
	})

	return func(*server.App) Node {
		return Section(
			Class("flex flex-col gap-8"),

			Div(
				Class("flex flex-col gap-4"),

				nodes,
			),

			Div(
				Class("flex gap-2"),

				IfFn(prev != nil, func() Item {
					return EndButtonLink("Go Back", prev.Label, prev.Href, false)
				}),
				IfFn(next != nil, func() Item {
					return EndButtonLink("Next Up", next.Label, next.Href, true)
				}),
			),
		)
	}
}

func EndButtonLink(label, title, href string, forwards bool) Node {
	return A(
		Href(href),
		Class("px-4 py-4 min-h-14 min-w-56 flex flex-col justify-center rounded-md border hover:bg-accent hover:border-transparent"),
		If(forwards, Class("ml-auto")),

		Span(
			Class("text-xs mb-1"),
			If(!forwards, Class("ml-auto")),

			Text(label),
		),
		Span(
			Class("flex justify-between items-center"),

			If(!forwards, lucide.ArrowLeft(Class("size-5"))),
			P(Class("font-semibold"), Text(title)),
			If(forwards, lucide.ArrowRight(Class("size-5"))),
		),
	)
}
