package app

import (
	"context"
	"embed"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/canpacis/pacis-docs/src/icons"
	. "github.com/canpacis/pacis/html"
	"github.com/canpacis/pacis/lucide"
	"github.com/canpacis/pacis/server"
	"github.com/canpacis/pacis/x"
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
	Href  string
}

var Docs = []DocTitle{
	{
		Label: "Getting Started",
		Href:  "/getting-started",
		Links: []DocLink{
			{Label: "Introduction", Href: "/getting-started/introduction/", File: "introduction"},
			{Label: "Installation", Href: "/getting-started/installation/", File: "installation"},
			{Label: "Quick Start", Href: "/getting-started/quick-start/", File: "quick-start"},
			{Label: "Templating", Href: "/getting-started/templating/", File: "templating"},
			{Label: "Conventions", Href: "/getting-started/conventions/", File: "conventions"},
		},
	},
	{
		Label: "Core Concepts",
		Href:  "/core-concepts",
		Links: []DocLink{
			{Label: "Rendering", Href: "/core-concepts/rendering/", File: "rendering"},
			{Label: "Request Data", Href: "/core-concepts/request-data/", File: "request-data"},
			{Label: "Streaming", Href: "/core-concepts/streaming/", File: "streaming"},
			{Label: "Middlewares", Href: "/core-concepts/middlewares/", File: "middlewares"},
			{Label: "Caching", Href: "/core-concepts/caching/", File: "caching"},
			{Label: "Assets", Href: "/core-concepts/assets/", File: "assets"},
			{Label: "Deploying", Href: "/core-concepts/deploying/", File: "deploying"},
		},
	},
}

func DocsLayout(app *server.Server, children Node) Node {
	return RootLayout(app, Main(
		Class("min-h-screen flex flex-col"),

		Div(
			Class("flex flex-1 max-w-6xl w-full mx-auto my-0 md:my-2 gap-6 h-full"),

			Nav(
				Class("hidden md:flex w-full flex-col max-w-[180px] h-fit lg:max-w-[224px] p-4 gap-6 sticky top-0"),

				Logo(),

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
			Div(
				Class("flex flex-col w-full"),

				Header(
					Class("py-3 my-3 px-6 flex md:hidden items-center justify-between sticky top-0 bg-background border-b z-40"),

					Div(
						MobileNav(),
					),
					Logo(),
				),
				Article(
					Class("flex-1 py-4 mb-4 md:mb-0 flex flex-col gap-6 md:gap-0 w-screen md:w-auto"),

					children,
				),
			),
		),
		Footer(
			Class("mt-auto flex w-full max-w-6xl mx-auto items-center justify-between px-8 py-3 border-t text-sm"),

			P(Class("text-muted-foreground"), Text("Made with Pacis")),
			Div(
				A(
					Href("https://github.com/canpacis/pacis"),
					Target("_blank"),

					icons.GithubMark(Class("text-muted-foreground hover:text-foreground transition-colors")),
				),
			),
		),
	))
}

func BuildMarkup(node parser.TreeNode[parser.DjotNode]) Node {
	var element *Element

	switch node.Type {
	case parser.ParagraphNode:
		element = P(Class("inline leading-relaxed"))
	case parser.LinkNode:
		element = A(
			Class("text-sky-400 underline inline"),
			Href(node.Attributes.Get("href")),
		)
		if strings.HasPrefix(node.Attributes.Get("href"), "https://") {
			element.SetAttribute("target", "_blank")
		}
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
		element = Blockquote(Class("border-l-5 border-accent pl-2 py-2"))
	case parser.VerbatimNode:
		element = Span(Class("bg-accent rounded-sm text-sm py-1 px-2 font-mono"))
	case parser.ThematicBreakNode:
		element = Hr()
	case parser.CodeNode:
		lang := node.Attributes.Get("$CodeLangKey")
		noaccessory := node.Attributes.Get("noaccessory")
		return CodeComponent(
			lang,
			string(node.FullText()),
			node.Attributes.Get("file"),
			lang == "go" && noaccessory != "true",
		)
	case parser.TextNode:
		return Text(strings.ReplaceAll(string(node.Text), "&rsquo;", "'"))
	default:
		nodes := Fragment()
		for _, child := range node.Children {
			nodes = append(nodes, BuildMarkup(child))
		}
		return nodes
	}

	class := node.Attributes.Get("class")
	if len(class) > 0 {
		element.AddClass(class)
	}
	for _, child := range node.Children {
		element.AppendNode(BuildMarkup(child))
	}
	return element
}

var docsfs embed.FS

func SetDocsFS(fs embed.FS) {
	docsfs = fs
}

func GetDocFile(env server.Environment, name string) []byte {
	if env == server.Dev {
		file, err := os.ReadFile("src/app/docs/" + name + ".md")
		if err != nil {
			log.Fatal(err)
		}
		return file
	}
	file, err := docsfs.ReadFile("src/app/docs/" + name + ".md")
	if err != nil {
		log.Fatal(err)
	}
	return file
}

func DocPage(env server.Environment, name string) func(*server.Server) Node {
	file := GetDocFile(env, name)

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

	return func(*server.Server) Node {
		return Div(
			Class("flex flex-col justify-between h-full gap-8 px-6 md:pl-0 md:pr-8"),

			Div(
				Class("flex flex-col gap-8 mb-20"),

				markup,
			),

			Div(
				Class("flex flex-col gap-6 mt-auto"),

				Div(
					Class("text-sm w-full flex flex-wrap gap-1 md:gap-3 items-end px-2"),

					P(
						Class("italic"),

						Text("Something wrong with this page?"),
					),
					A(
						Href(fmt.Sprintf("https://github.com/canpacis/pacis-docs/edit/main/src/app/docs/%s.md", name)),
						Target("_blank"),
						Class("flex flex-nowrap gap-2 items-center text-sky-400"),

						Text("Edit It"),
						lucide.Pen(Class("size-4")),
					),
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
			),
		)
	}
}

func EndButtonLink(label, title, href string, forwards bool) Node {
	return A(
		Href(href),
		Class("px-4 py-4 min-h-14 min-w-34 w-full md:w-fit md:min-w-56 flex flex-col justify-center rounded-md border hover:bg-accent hover:border-transparent text-sm md:text-base"),
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

func Logo() Node {
	return A(
		Class("flex flex-row-reverse md:flex-row gap-2 items-center"),
		Href("/"),

		Img(Src("/logo.webp"), Class("h-7")),
		Span(Class("uppercase text-lg font-light select-none"), Text("Pacis")),
	)
}

func MobileNav() Node {
	return Span(
		x.Data(map[string]any{"open": false}),

		Button(x.On("click", "open = !open"), Class("pr-2 py-1"), lucide.Menu(Class("size-5"))),

		Div(
			x.Cloak,
			x.Show("open"),
			x.On("click", "open = false", x.Outside),
			Attr("x-transition:enter-start", "translate-x-[-60vw]"),
			Attr("x-transition:enter-end", "translate-x-0"),
			Attr("x-transition:leave-start", "translate-x-0"),
			Attr("x-transition:leave-end", "translate-x-[-60vw]"),
			Attr("x-trap.noscroll", "open"),
			Class("fixed z-50 top-0 left-0 bg-background border-r h-screen w-[60vw] flex transition-transform duration-500"),

			Nav(
				Class("flex flex-col w-full h-full p-4 gap-6"),

				Map(Docs, func(title DocTitle) Node {
					return Div(
						Class("flex flex-col gap-2"),

						P(Class("font-semibold"), Text(title.Label)),
						Ul(
							Class("flex flex-col gap-0.5"),

							Map(title.Links, func(link DocLink) Node {
								return Li(
									A(
										Class("hover:bg-accent w-full h-8 flex text-sm rounded-md px-3 items-center data-[state=active]:bg-accent"),
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
		),
	)
}
