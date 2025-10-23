# Middlewares

Middlewares connect one process to another.
In the context of **Pacis**, they let you write **transformers** for your endpoints.

These transformations operate at the **HTTP level**, not the **page level**, meaning you don’t change how a specific page behaves or renders, but rather how HTTP requests to that route are handled.

## Shape

Pacis’ `server` package builds on Go’s standard `net/http` library and stays close to its conventions.
For example, `server.New` returns an `http.Handler`, allowing you to use it anywhere the standard handler interface is expected.

Middlewares follow the same principle.
The `server/middleware` package defines a simple `Middleware` interface:

{file="server/middleware.go"}
```go
type Middleware interface {
  Apply(http.Handler) http.Handler
}
```

The `Middleware` interface allows you to attach data to a struct and provide an `Apply` method to turn it into a reusable middleware.
Because `Apply` both accepts and returns an `http.Handler`, it’s compatible with Go’s conventional middleware pattern, meaning you can also use most third-party middlewares that follow this shape.

## Example: Cache Middleware

Let’s create a custom cache middleware that controls how long responses can be cached.

{noaccessory=true}
```go
type CacheMiddleware struct {
  Duration time.Duration
}
```

Now implement the `Apply` method:

{noaccessory=true}
```go
func (m *CacheMiddleware) Apply(next http.Handler) http.Handler {
  return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Cache-Control", fmt.Sprintf("public, max-age=%d", int64(m.Duration.Seconds())))
    next.ServeHTTP(w, r)
  })
}
```

> Note: the middleware wraps the incoming handler, sets the cache header, and then calls the next handler in the chain.

Caching is such a common use case that Pacis provides a built-in cache middleware in the `server/middleware` package:

{noaccessory=true}
```go
// Cache responses for one day
cache := middleware.NewCacheMiddleware(time.Hour * 24)
server.HandlePage("/", Page, Layout, cache)
```

From logging and locale management to automatic color scheme handling and authentication,
Pacis provides a variety of built-in middlewares.

See the [API Reference](/api-reference/middleware) for a complete list.