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
			Class("flex max-w-6xl flex-1 w-full mx-auto my-2"),

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
			Article(
				Class("flex-1 py-4 pr-8"),

				children,
			),
		),
		Footer(
			Class("mt-auto flex w-full max-w-6xl mx-auto items-center justify-center py-3 border-t text-sm"),

			P(Class("text-muted-foreground"), Text("Made with Pacis")),
		),
	))
}

func BuildMarkup(node parser.TreeNode[parser.DjotNode]) Node {
	var element *Element

	switch node.Type {
	case parser.ParagraphNode:
		element = P(Class("inline"))
	case parser.LinkNode:
		element = A(
			Class("text-sky-400 underline inline"),
			Href(node.Attributes.Get("href")),
			Target("_blank"),
		)
	case parser.UnorderedListNode:
		element = Ul(Class("list-disc list-inside"))
	case parser.OrderedListNode:
		element = Ol(Class("list-decimal list-inside"))
	case parser.ListItemNode:
		element = Li()
	case parser.HeadingNode:
		switch node.Attributes.Get(parser.HeadingLevelKey) {
		case "#":
			element = H1(Class("text-3xl font-semibold mb-2"))
		case "##":
			element = H2(Class("text-2xl font-semibold mb-2"))
		case "###":
			element = H3(Class("text-xl font-semibold mb-2"))
		default:
			element = H4(Class("text-lg font-semibold mb-2"))
		}
	case parser.SectionNode:
		element = Section(Class("flex flex-col gap-4"), ID(node.Attributes.Get("id")))
	case parser.StrongNode:
		element = Span(Class("font-semibold"))
	case parser.SuperscriptNode:
		element = Sup()
	case parser.SubscriptNode:
		element = Sub()
	case parser.FootnoteDefNode:
		element = Span()
	case parser.EmphasisNode:
		element = Span(Class("italic"))
	case parser.ImageNode:
		element = Img(Class("w-full rounded-lg"), Src(node.Attributes.Get("src")))
	case parser.QuoteNode:
		element = Blockquote(Class("border-l-4 border-accent pl-3 py-2"))
	case parser.VerbatimNode:
		element = Span(Class("bg-accent rounded-sm text-sm py-1 px-2"))
	case parser.ThematicBreakNode:
		element = Hr()
	case parser.CodeNode:
		return CodeComponent(node.Attributes.Get("$CodeLangKey"), string(node.FullText()))
	case parser.TextNode:
		return Text(node.Text)
	default:
		nodes := Fragment()
		for _, child := range node.Children {
			nodes = append(nodes, BuildMarkup(child))
		}
		return nodes
	}

	for _, child := range node.Children {
		element.AppendNode(BuildMarkup(child))
	}
	return element
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

	markup := BuildMarkup(parser.BuildDjotAst(file)[0])

	return func(*server.App) Node {
		return Div(
			Class("flex flex-col justify-between h-full gap-8"),

			markup,

			Div(
				Class("flex gap-2 mt-auto"),

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
