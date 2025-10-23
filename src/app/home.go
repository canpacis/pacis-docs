package app

import (
	"fmt"
	"os"

	"components/ui/button"

	"github.com/canpacis/pacis-docs/src/icons"
	"github.com/canpacis/pacis/components"
	. "github.com/canpacis/pacis/html"
	"github.com/canpacis/pacis/lucide"
	"github.com/canpacis/pacis/server/metadata"
)

type HomePage struct{}

type Metadata struct {
	URL         string
	Title       string
	Description string
	Image       string
}

func (*HomePage) Metadata() *metadata.Metadata {
	base := os.Getenv("WEBSITE_URL")

	title := "Pacis | The Web Dev-Kit for Go Developers."
	desc := "Write type-safe templates with IntelliSense support, compose reusable components that make sense. When you're ready to ship, compile everything down to one binary and deploy anywhere. Performance is built-in, caching is opt-in, and streaming is automatic."
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

func (*HomePage) Page() Node {
	return Main(
		Div(
			Class("relative flex h-[500px] w-full min-h-screen flex-col items-center justify-center overflow-hidden rounded-lg"),

			HeroBackground(
				Class("top-0 bottom-0 m-auto skew-y-12 z-0 h-screen w-screen lg:h-auto lg:w-auto [mask-image:radial-gradient(400px_circle_at_center,white,transparent)] md:[mask-image:radial-gradient(600px_circle_at_center,white,transparent)]"),
			),

			Div(
				Class("relative z-10 flex flex-col items-center gap-8 max-w-xl px-4 mx-auto text-center"),

				Span(
					Class("flex gap-2 items-center"),

					Img(Src("/icon.svg"), Class("h-7")),
					H2(Class("uppercase text-2xl font-light select-none"), Text("Pacis")),
				),
				H1(
					Class("text-3xl md:text-5xl font-black"),

					Text("The Web Dev Kit for Go Developers."),
				),
				P(
					Class("font-medium"),

					Text("Write type-safe templates with IntelliSense support, compose reusable components that make sense. When you're ready to ship, compile everything down to one binary and deploy anywhere. Performance is built-in, caching is opt-in, and streaming is automatic."),
				),

				Div(
					Class("flex justify-center gap-4"),

					button.New(
						components.AsChild, button.Secondary,

						A(
							Href("https://github.com/canpacis/pacis"),
							Target("_blank"),

							icons.GithubMark(),
							Text("Github"),
						),
					),
					button.New(
						components.AsChild,

						A(
							Href("/getting-started"),

							Text("Get Started"),
							lucide.ArrowUpRight(),
						),
					),
				),
			),
		),
	)
}

func HeroBackground(items ...Item) Node {
	size := 64
	vertical := 18
	horizontal := 18
	list := make([]struct{}, horizontal*vertical)

	return El("svg",
		components.ItemsOf(
			items,
			Width(fmt.Sprintf("%d", size*vertical)),
			Height(fmt.Sprintf("%d", size*horizontal)),
			Class("absolute size-fit"),

			MapIdx(list, func(item struct{}, i int) Node {
				x := i % horizontal * size
				y := i / horizontal * size

				return El("rect",
					Attr("x", fmt.Sprintf("%d", x)),
					Attr("y", fmt.Sprintf("%d", y)),
					Width(fmt.Sprintf("%d", size)),
					Height(fmt.Sprintf("%d", size)),
					Class("stroke-border transition-all duration-100 ease-in-out [&:not(:hover)]:duration-1000 fill-background hover:fill-border"),
				)
			}),
		)...,
	)
}
