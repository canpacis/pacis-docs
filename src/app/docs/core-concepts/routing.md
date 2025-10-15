# Routing

Unlike the zeitgeist that is today's meta-frameworks, Pacis does not use a file based router. Routes are registered manually. You need to create a `Server` to start routing.

{file="main.go"}
```go
server := server.New(*server.Options{})
```

The server struct is simply an extended `*http.ServeMux`, so you can use it like a regular multiplexer, register handlers, pass it to places where an `http.Handler` interface is expected.

You can register your pages with their layouts and middlewares as well.

{file="main.go"}
```go
server.HandlePage("/", HomePage, RootLayout, middleware.DefaultGzip)
```

But what is `HomePage`? It could be a function or struct that satisfies the `Page` interface.

{noaccessory=true}
```go
// It must return a Node
func HomePage() Node

// Or...

type Home struct {}

// A Page function that returns a Node
func (*Home) Page() Node 

var HomePage = &Home{}
```

> This is to be able to provide additional [metadata](/api-reference/metadata) about a page, like `<meta>` elements.

The `RootLayout` is a bit more strict about this stuff. It must strictly be of type `server.LayoutFn`:

{file="server/handler.go"}
```go
type LayoutFn func(head html.Node, children html.Node) html.Node
```

Layout, like a middleware, takes `head` and `children` values and returns a final `Node`.

Speaking of middlewares, lastly, the `HandlePage` method takes any amount of middlewares that will be applied to that route. Let's talk about them in the next section.