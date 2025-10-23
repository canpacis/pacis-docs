package app

import (
	"context"
	"fmt"
	"os"
	"path"

	"components/ui/sheet"

	"github.com/canpacis/pacis-docs/src/icons"
	. "github.com/canpacis/pacis/html"
	"github.com/canpacis/pacis/lucide"
	"github.com/canpacis/pacis/server"
	"github.com/canpacis/pacis/server/metadata"
	"github.com/canpacis/pacis/x"
	parser "github.com/sivukhin/godjot/djot_parser"
)

type DocLink struct {
	Label string
	Href  string
	File  string
}

type DocTitle struct {
	Label  string
	Links  []DocLink
	Folder string
	Href   string
}

var Docs = []DocTitle{
	{
		Label:  "Getting Started",
		Folder: "getting-started",
		Href:   "/getting-started",
		Links: []DocLink{
			{Label: "Introduction", Href: "/getting-started/introduction/", File: "introduction"},
			{Label: "Installation", Href: "/getting-started/installation/", File: "installation"},
			{Label: "Quick Start", Href: "/getting-started/quick-start/", File: "quick-start"},
			{Label: "Templating", Href: "/getting-started/templating/", File: "templating"},
			{Label: "Conventions", Href: "/getting-started/conventions/", File: "conventions"},
		},
	},
	{
		Label:  "Core Concepts",
		Folder: "core-concepts",
		Href:   "/core-concepts",
		Links: []DocLink{
			{Label: "Rendering", Href: "/core-concepts/rendering/", File: "rendering"},
			{Label: "Request Data", Href: "/core-concepts/request-data/", File: "request-data"},
			{Label: "Streaming", Href: "/core-concepts/streaming/", File: "streaming"},
			{Label: "Routing", Href: "/core-concepts/routing/", File: "routing"},
			{Label: "Middlewares", Href: "/core-concepts/middlewares/", File: "middlewares"},
			{Label: "Assets", Href: "/core-concepts/assets/", File: "assets"},
			{Label: "Deploying", Href: "/core-concepts/deploying/", File: "deploying"},
		},
	},
	{
		Label:  "API Reference",
		Folder: "api-reference",
		Href:   "/api-reference",
		Links: []DocLink{
			{Label: "Nodes", Href: "/api-reference/nodes/", File: "nodes"},
			{Label: "Attributes", Href: "/api-reference/attributes/", File: "attributes"},
			{Label: "Server", Href: "/api-reference/server/", File: "server"},
			{Label: "Middleware", Href: "/api-reference/middleware/", File: "middleware"},
			{Label: "Request Data", Href: "/api-reference/request-data/", File: "request-data"},
			{Label: "Metadata", Href: "/api-reference/metadata/", File: "metadata"},
		},
	},
	// {
	// 	Label:  "Components",
	// 	Folder: "components",
	// 	Href:   "/components",
	// 	Links: []DocLink{
	// 		{Label: "Alert", Href: "/components/alert/", File: "alert"},
	// 	},
	// },
}

func DocsLayout(app *server.Server, head, children Node) Node {
	return RootLayout(app, head, Main(
		Class("min-h-screen flex flex-col"),

		Div(
			Class("flex flex-1 max-w-6xl w-full mx-auto my-0 md:my-2 gap-6 h-full"),

			Nav(
				Class("hidden md:flex w-full flex-col max-w-[180px] h-fit lg:max-w-[224px] p-4 gap-4 sticky top-0"),

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
											detail, err := server.Detail(ctx)
											if err == nil && detail.URL.Path == link.Href {
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
				x.Data(map[string]any{"sheetOpen": false}),
				Class("flex flex-col w-full"),

				Header(
					x.Bind("data-sheet-state", "sheetOpen ? 'open' : 'closed'"),
					Class("pt-3 px-6 flex md:hidden items-center justify-between top-0 bg-background border-b z-40 data-[sheet-state='open']:fixed data-[sheet-state='open']:w-screen data-[sheet-state='closed']:sticky"),

					sheet.New(
						x.On("open", "sheetOpen = true"),
						x.On("closed", "sheetOpen = false"),
						sheet.Trigger(lucide.Menu(Class("size-5"))),
						sheet.Content(
							sheet.Left,
							Class("max-h-screen overflow-y-auto"),

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
															detail, err := server.Detail(ctx)
															if err == nil && detail.URL.Path == link.Href {
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
					),
					Logo(),
				),
				Article(
					x.Bind("data-sheet-state", "sheetOpen ? 'open' : 'closed'"),
					Class("flex-1 py-4 mb-4 md:mb-0 flex flex-col gap-6 md:gap-0 w-screen md:w-auto data-[sheet-state='open']:mt-[69px]"),

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

var Version string = os.Getenv("VERSION")

func Logo() Node {
	return A(
		Class("flex flex-row-reverse md:flex-row gap-3 items-center hover:bg-accent rounded-md px-4 md:px-2 h-14"),
		Href("/"),

		Img(Src("/icon.svg"), Class("h-8")),
		Span(
			Class("flex flex-col items-end md:items-start"),

			Span(Class("uppercase text-base font-light select-none"), Text("Pacis")),
			Span(Class("text-xs text-muted-foreground md:block"), Text(Version)),
		),
	)
}

type DocPage struct {
	File        string
	DocTitle    *DocTitle
	Title       string
	Description string
	Markup      Node
	Prev        *DocLink
	Next        *DocLink
}

func (p *DocPage) Metadata() *metadata.Metadata {
	base := os.Getenv("WEBSITE_URL")

	title := fmt.Sprintf("%s - %s | Pacis Docs", p.Title, p.DocTitle.Label)
	desc := p.Description
	image := base + "/og-image.png"

	return &metadata.Metadata{
		Title:       title,
		Description: desc,
		OpenGraph: &metadata.OpenGraph{
			URL:         base,
			Title:       title,
			Description: desc,
			Images: []metadata.OpenGraphMedia{
				{URL: image},
			},
		},
		Twitter: &metadata.Twitter{
			Card:        "summary_large_image",
			Title:       title,
			Description: desc,
			Images:      []string{image},
		},
	}
}

func (p *DocPage) Page() Node {
	return Div(
		Class("flex flex-col justify-between h-full gap-8 px-6 md:pl-0 md:pr-8"),

		Div(
			Class("flex flex-col gap-8 mb-20"),

			p.Markup,
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
					Href(fmt.Sprintf("https://github.com/canpacis/pacis-docs/edit/main/src/app/docs/%s/%s.md", p.DocTitle.Folder, p.File)),
					Target("_blank"),
					Class("flex flex-nowrap gap-2 items-center text-sky-400"),

					Text("Edit It"),
					lucide.Pen(Class("size-4")),
				),
			),
			Div(
				Class("flex gap-2"),

				IfFn(p.Prev != nil, func() Item {
					return EndButtonLink("Go Back", p.Prev.Label, p.Prev.Href, false)
				}),
				IfFn(p.Next != nil, func() Item {
					return EndButtonLink("Next Up", p.Next.Label, p.Next.Href, true)
				}),
			),
		),
	)
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

func NewDocPage(env server.Environment, doc, name string) *DocPage {
	page := &DocPage{File: name, Title: "Hello", Description: "hello"}

	for _, title := range Docs {
		if title.Folder != doc {
			continue
		}
		for i, link := range title.Links {
			if link.File == name {
				page.DocTitle = &title
				if i > 0 {
					page.Prev = &title.Links[i-1]
				}
				if i < len(title.Links)-1 {
					page.Next = &title.Links[i+1]
				}
			}
		}
	}

	file := GetDocFile(env, path.Join(doc, name))
	node := parser.BuildDjotAst(file)[0]
	page.Title, page.Description = ExtractMetadata(node)
	page.Markup = BuildMarkup(node)
	return page
}
