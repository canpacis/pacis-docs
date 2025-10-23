# Server

The server package provides the core HTTP server implementation for the Pacis framework, including environment configuration, middleware management, static asset handling, development server proxying, and graceful shutdown capabilities.

---

## Types

### Environment

{noaccessory=true}
```go
type Environment string
```

Environment describes the app environment and lets the other parts of the app choose environment specific behaviour.

**Constants:**

- `Dev = Environment("DEVELOPMENT")` - Represents a development environment of the app
- `Prod = Environment("PRODUCTION")` - Represents a production environment of the app


### Options

{noaccessory=true}
```go
type Options struct {
    Env       Environment
    Port      string
    DevServer *url.URL
    Logger    *slog.Logger
    Mux       *http.ServeMux
}
```

Options holds the configuration settings for the server, including environment, port, development server URL, logger, and HTTP request multiplexer.

---

### Server

{noaccessory=true}
```go
type Server struct {
    *http.ServeMux
    // contains filtered or unexported fields
}
```

Server embeds `http.ServeMux` and extends it with middleware support, asset manifest management, and environment-specific behaviors.

#### Methods

##### Use

{noaccessory=true}
```go
func (s *Server) Use(middlewares ...middleware.Middleware)
```

Adds middleware(s) to the application's middleware stack.

##### HandlePage

{noaccessory=true}
```go
func (s *Server) HandlePage(pattern string, page Page, layout Layout, middlewares ...middleware.Middleware)
```

Registers a page handler at the specified pattern with optional layout and middlewares. Automatically registers POST actions if the page implements the Actions interface.

##### SetBuildDir

{noaccessory=true}
```go
func (s *Server) SetBuildDir(name string, dir fs.FS, vite fs.FS) error
```

Configures the build directory for serving static assets and loads the Vite manifest for production builds.

##### RegisterDevHandlers

{noaccessory=true}
```go
func (s *Server) RegisterDevHandlers()
```

Sets up reverse proxy handlers for development mode, forwarding asset requests to the development server.

##### SetNotFoundPage

{noaccessory=true}
```go
func (s *Server) SetNotFoundPage(page Page, layout Layout)
```

Configures a custom 404 not found page with the specified layout.

##### Asset

{noaccessory=true}
```go
func (s *Server) Asset(name string) string
```

Returns the URL or path for a given asset name. In development mode, constructs the asset URL using the development server. In production mode, retrieves the asset entry from the manifest.

##### Serve

{noaccessory=true}
```go
func (s *Server) Serve()
```

Starts the HTTP server and handles graceful shutdown on receiving SIGINT or SIGTERM signals.

---

### Page

{noaccessory=true}
```go
type Page interface {
    Page() html.Node
}
```

Page defines the interface for renderable pages. Types implementing this interface must also provide a `Metadata()` method.

---

### PageFunc

{noaccessory=true}
```go
type PageFunc func() html.Node
```

PageFunc is a function type that implements the Page interface.

#### Methods

##### Page

{noaccessory=true}
```go
func (p PageFunc) Page() html.Node
```

Returns the HTML node for the page.

##### Metadata

{noaccessory=true}
```go
func (PageFunc) Metadata() *metadata.Metadata
```

Returns empty metadata for the page.

---

### Layout

{noaccessory=true}
```go
type Layout func(*Server, html.Node, html.Node) html.Node
```

Layout defines a function type that takes a server, head node, and children node as input, and returns a modified html.Node. It is typically used to apply layout transformations or wrappers to HTML nodes.

---

### ActionFunc

{noaccessory=true}
```go
type ActionFunc func(http.ResponseWriter, *http.Request) error
```

ActionFunc defines a function type for handling form actions.

---

### RequestDetail

{noaccessory=true}
```go
type RequestDetail struct {
    Header     http.Header
    URL        *url.URL
    Host       string
    Method     string
    Pattern    string
    RemoteAddr string
    RequestURI string
    Cookies    []*http.Cookie
}
```

RequestDetail provides access to HTTP request information within server rendering contexts.

#### Methods

##### Cookie

{noaccessory=true}
```go
func (d *RequestDetail) Cookie(name string) (*http.Cookie, error)
```

Returns the named cookie from the request, or an error if not found.

---

## Functions

### New

{noaccessory=true}
```go
func New(options *Options) *Server
```

Creates and returns a new Server instance using the provided Options. Sets default values for any missing options, including the HTTP mux, development server URL, environment, and logger. Attaches default middleware for logging and recovery, and sets the default "not found" page handler.

---

### DefaultLayout

{noaccessory=true}
```go
func DefaultLayout(server *Server, head html.Node, children html.Node) html.Node
```

DefaultLayout provides a standard HTML document structure with doctype, head, and body elements.

---

### PageHandler

{noaccessory=true}
```go
func PageHandler(server *Server, page Page, layout Layout, middlewares ...middleware.Middleware) http.Handler
```

Creates an HTTP handler that processes requests using the provided page, applies the specified layout function, and supports an optional list of middleware functions. The handler renders the page, applies the layout, and supports chunked rendering if needed. All provided and application-level middlewares are applied in order.

---

### ActionsHandler

{noaccessory=true}
```go
func ActionsHandler(server *Server, actions map[string]ActionFunc, middlewares ...middleware.Middleware) http.Handler
```

Creates an HTTP handler for processing form actions. Actions are identified by the `__action` query parameter and dispatched to the corresponding ActionFunc.

---

### Async

{noaccessory=true}
```go
func Async(comp html.Component, fallback html.Node) html.Component
```

Wraps a component for asynchronous rendering with an optional fallback node displayed during loading.

---

### Data

{noaccessory=true}
```go
func Data[T any](ctx context.Context) (*T, error)
```

Extracts and populates a data structure from the request context, scanning path parameters, query parameters, cookies, and headers.

---

### FormData

{noaccessory=true}
```go
func FormData[T any](r *http.Request) (*T, error)
```

Parses and populates a data structure from form POST data.

---

### Detail

{noaccessory=true}
```go
func Detail(ctx context.Context) (*RequestDetail, error)
```

Retrieves HTTP request details from the server rendering context.

---

### Redirect

{noaccessory=true}
```go
func Redirect(ctx context.Context, to string) html.Node
```

Returns a node that triggers an HTTP redirect (302 Found) to the specified URL during rendering.

---

### RedirectComponent

{noaccessory=true}
```go
func RedirectComponent(to string) html.Component
```

Returns a component that triggers an HTTP redirect (302 Found) to the specified URL.

---

### RedirectWith

{noaccessory=true}
```go
func RedirectWith(ctx context.Context, to string, status int) html.Node
```

Returns a node that triggers an HTTP redirect with a custom status code to the specified URL during rendering.

---

### RedirectWithComponent

{noaccessory=true}
```go
func RedirectWithComponent(to string, status int) html.Component
```

Returns a component that triggers an HTTP redirect with a custom status code to the specified URL.

---

### NotFound

{noaccessory=true}
```go
func NotFound(ctx context.Context) html.Node
```

Returns a node that marks the current request as not found, triggering the 404 handler.

---

### SetCookie

{noaccessory=true}
```go
func SetCookie(ctx context.Context, cookie *http.Cookie) html.Node
```

Returns a node that sets an HTTP cookie during rendering.

---

### SetCookieComponent

{noaccessory=true}
```go
func SetCookieComponent(cookie *http.Cookie) html.Component
```

Returns a component that sets an HTTP cookie during rendering.

---

### Form

{noaccessory=true}
```go
func Form(name string, items ...html.Item) html.Node
```

Creates an HTML form element configured to POST to the specified action handler. The action name is automatically appended as the `__action` query parameter.

---

## Usage Example

{noaccessory=true}
```go
package main

import (
    "github.com/canpacis/pacis/server"
    "github.com/canpacis/pacis/html"
)

func main() {
    options := &server.Options{
        Env:  server.Dev,
        Port: ":8080",
    }
    
    srv := server.New(options)
    
    srv.HandlePage("/", server.PageFunc(func() html.Node {
        return html.H1(html.Text("Hello, World!"))
    }), server.DefaultLayout)
    
    srv.Serve()
}
```