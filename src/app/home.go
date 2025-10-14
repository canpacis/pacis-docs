package app

import (
	"fmt"

	"github.com/canpacis/pacis/components"
	. "github.com/canpacis/pacis/html"
	"github.com/canpacis/pacis/lucide"
	"github.com/canpacis/pacis/server"
)

func Home(*server.Server) Node {
	return Main(
		Div(
			Class("relative flex h-[500px] w-full min-h-screen flex-col items-center justify-center overflow-hidden rounded-lg"),

			HeroBackground(
				Class("top-0 bottom-0 m-auto skew-y-12 z-0 h-screen w-screen lg:h-auto lg:w-auto [mask-image:radial-gradient(400px_circle_at_center,white,transparent)] md:[mask-image:radial-gradient(800px_circle_at_center,white,transparent)]"),
			),

			Div(
				Class("relative z-10 flex flex-col items-center gap-8 max-w-xl px-4 mx-auto text-center"),

				Span(
					Class("flex gap-2 items-center"),

					Img(Src("/logo.webp"), Class("h-7")),
					H2(Class("uppercase text-2xl font-light select-none"), Text("Pacis")),
				),
				H1(
					Class("text-3xl md:text-5xl font-black"),

					Text("The SSR Library for Go Developers."),
				),
				P(
					Class("font-medium"),

					Text("Write type-safe templates with IntelliSense support, compose reusable components that make sense. When you're ready to ship, compile everything down to one binary and deploy anywhere. Performance is built-in, caching is opt-in, and streaming is automatic."),
				),

				Div(
					Class("flex justify-center"),

					A(
						Class("inline-flex items-center justify-center gap-2 whitespace-nowrap rounded-md text-sm font-medium transition-colors focus-visible:outline-none focus-visible:ring-1 focus-visible:ring-ring disabled:pointer-events-none disabled:opacity-50 [&_svg]:pointer-events-none [&_svg]:size-4 [&_svg]:shrink-0 bg-primary text-primary-foreground shadow hover:bg-primary/90 h-10 px-4 py-2"),
						Href("/getting-started"),

						Text("Get Started"),
						lucide.ArrowUpRight(),
					),
				),
			),
		),
	)
}

func HeroBackground(items ...Item) Node {
	size := 64
	vertical := 24
	horizontal := 24
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
					Class("stroke-border transition-all duration-100 ease-in-out [&:not(:hover)]:duration-1000 fill-transparent hover:fill-border"),
				)
			}),
		)...,
	)
}
