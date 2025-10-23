# Quick Start

Your entry point to your web app is the `main.go` file in the root of your working directory.

{file="main.go"}
```go
package main

import (
	"github.com/canpacis/pacis-app/src/app"
	"github.com/canpacis/pacis/server"
	"github.com/canpacis/pacis/server/middleware"
)

// These options will be populated based 
// on your environment (dev, prod)
var options *server.Options

func main() {
  // Create a new server
  server := server.New(options)

  // Ready function manages your static assets.
  // See prod.go & dev.go files.
  Ready(server)

  // Register your routes
  server.HandlePage("GET /", app.HomePage, app.RootLayout, middleware.DefaultGzip)

  server.Serve()
}
```

Right now, it creates a server, registers a home page and serves your application. The `*server.Server` struct is basically an http multiplexer, you can use it like a `*http.ServeMux` but it is enchanced with a few more capabilities.

Go ahead and open up the `src/app/home.go` file where the home page is rendered.

{file="src/app/home.go"}
```go
package app

import (
  // Dot import to not use the html. prefix all the time.
  . "github.com/canpacis/pacis/html"
  "github.com/canpacis/pacis/server"
)

func HomePage(*server.Server) Node {
  return Div(
    Class("min-h-screen flex flex-col items-center justify-center"),

    Img(Src("/icon.svg"), Class("h-16")),
    H1(
      Class("text-4xl font-bold my-4"),

      Text("Build With Pacis"),
    ),
    P(
      Class("flex items-center gap-2"),

      Text("Update the "),
      Span(
        Class("bg-neutral-200 text-neutral-700 pointer-events-none inline-flex h-5 w-fit min-w-5 items-center justify-center gap-1 rounded-sm px-1 font-sans text-xs font-medium select-none [&_svg:not([class*='size-'])]:size-3 [[data-slot=tooltip-content]_&]:bg-background/20 [[data-slot=tooltip-content]_&]:text-background dark:[[data-slot=tooltip-content]_&]:bg-background/10"),

        Text("src/app/home.go"),
      ),
      Text(" file to see changes."),
    ),
    A(
      Target("_blank"),
      Href("https://pacis.canpacis.com"),
      Class("underline text-sky-400 text-sm mt-4"),

      Text("Learn More"),
    ),
  )
}
```

This file uses Pacis' [templating language](/getting-started/templating) to build up a page. We will talk about it in the next section.

Try changing some of the content and see the changes instantly!

***

Last, but not least, you should have a look at the `src/app/layout.go` file where the layout of this page comes from.

{file="src/app/layout.go"}
```go
package app

import (
	. "github.com/canpacis/pacis/html"
	"github.com/canpacis/pacis/server"
	"github.com/canpacis/pacis/server/font"
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
			Script(Defer, Type("module"), Src(server.Asset("/src/web/main.ts"))),
			server.HMR(),
			font.Head(
				font.New("Inter", font.WeightList{font.W100, font.W900}, font.Auto, font.Latin, font.LatinExt),
			),
		),
		Body(
			children,
			Script(Src(server.Asset("/src/web/stream.ts"))),
		),
	)
}
```

The layout function uses the server to dynamically get assets paths seamlessly. The call to `server.Asset()` will return the bundled file path in production and the vite proxy path in development.

It also calls the `server.HMR()` function for invluding vite's hmr functionality. In production, this function will no-op.

It uses Pacis' font package `github.com/canpacis/pacis/server/font` to load and configure a [Google Font](https://fonts.google.com/).

And lastly, inside the body, it renders its children along with an helper TypeScript file for async streaming.