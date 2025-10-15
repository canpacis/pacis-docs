# Request Data

## Generic Request Data

You can some basic request detail via the `Detail()` function.

{file="server/items.go"}
```go
func Detail(context.Context) *RequestDetail

type RequestDetail struct {
	URL *url.URL
	Host,
	Method,
	Pattern,
	RemoteAddr,
	RequestURI string
	Cookies []*http.Cookie
}
```

Use it like so:

{noaccessory=true}
```go
func Component(ctx context.Context) Node {
  detail := server.Detail(ctx)
  ...
}
```

## Your Request Data

More often than not, you will need a set of data which can get tricky to put together. Server provides a `Data()` function to pass your `struct`s to populate it with the information you need.

You need to use struct tags, similar to json, to scan the values onto your struct.

{noaccessory=true}
```go
type PageData struct {
  Language string `header:"Accept-Language"` // Read Accept-Language header
  Token string `cookie:"access_token"` // Read "access_token" cookie
  Search string `query:"search"` // Read "search" url search parameter
  Slug string `path:"slug"` // Read dynamic path value "slug"
}

func Component(ctx context.Context) Node {
  data, err := server.Data[PageData](ctx)
  ...
}
```

> The `Data()` function is generic, you need to pass your `struct` type as the type parameter.

> The returned data will be a pointer to your `struct` type e.g: `*PageData`