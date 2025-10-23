# Middleware

Package middleware provides HTTP middleware utilities for handling color scheme preferences, locale selection, caching, logging, and gzip compression.

## Interfaces

### Middleware

{noaccessory=true}
```go
type Middleware interface {
    Apply(http.Handler) http.Handler
}
```

The Middleware interface defines the standard interface for all middleware components in the package.

### Authenticator

{noaccessory=true}
```go
type Authenticator interface {
    Authenticate(*http.Request) (any, error)
    OnError(http.ResponseWriter, *http.Request, error)
}
```

The Authenticator interface defines methods for authenticating requests and handling authentication errors.

## Types

### ColorScheme

{noaccessory=true}
```go
type ColorScheme struct {
    UseHeaders bool
    Key        string
}
```

ColorScheme is a middleware that manages the user's color scheme preference via a cookie. It checks for the "pacis_color_scheme" cookie in the incoming request. If the cookie exists and its value is "light" or "dark", it uses that value as the theme. Otherwise, it defaults to "light" and sets the cookie accordingly. The selected theme is stored in the request's context under the key "theme" for downstream handlers to access.

#### Methods

**Name**
{noaccessory=true}
```go
func (*ColorScheme) Name() string
```

**Apply**
{noaccessory=true}
```go
func (m *ColorScheme) Apply(h http.Handler) http.Handler
```

### Locale

{noaccessory=true}
```go
type Locale struct {
    // contains filtered or unexported fields
}
```

Locale is a middleware that determines the user's preferred language from a cookie ("pacis_locale"), a "lang" form value, or the "Accept-Language" HTTP header, falling back to the provided default language if none are set or valid. It parses the locale, creates an i18n.Localizer, and injects both the localizer and the language tag into the request context for downstream handlers to use.

#### Methods

**Name**
{noaccessory=true}
```go
func (*Locale) Name() string
```

**Apply**
{noaccessory=true}
```go
func (l *Locale) Apply(h http.Handler) http.Handler
```

### Cache

{noaccessory=true}
```go
type Cache struct {
    // contains filtered or unexported fields
}
```

Cache returns a middleware that sets the "Cache-Control" header on HTTP responses, specifying that the response can be cached by any cache and defining the maximum age (in seconds) that the response is considered fresh. The duration parameter determines the max-age value. This middleware should be used to control client and proxy caching behavior for HTTP handlers.

#### Methods

**Name**
{noaccessory=true}
```go
func (*Cache) Name() string
```

**Apply**
{noaccessory=true}
```go
func (c *Cache) Apply(h http.Handler) http.Handler
```

### Logger

{noaccessory=true}
```go
type Logger struct {
    // contains filtered or unexported fields
}
```

Logger returns a middleware that logs HTTP requests using the provided slog.Logger. It records the request method, status code, path, remote address, user agent, and duration. The middleware wraps the next http.Handler and logs the request details after it is served.

**Example usage:**

{noaccessory=true}
```go
mux.Use(Logger(logger))
```

**Parameters:**
- logger - the slog.Logger instance used for logging request details.

**Returns:**
- A middleware function compatible with http.Handler.

#### Methods

**Name**
{noaccessory=true}
```go
func (*Logger) Name() string
```

**Apply**
{noaccessory=true}
```go
func (l *Logger) Apply(h http.Handler) http.Handler
```

### Gzip

{noaccessory=true}
```go
type Gzip struct {
    Dev bool
}
```

GzipHandler wraps an HTTP handler, to transparently gzip the response body if the client supports it (via the Accept-Encoding header). This will compress at the default compression level.

#### Methods

**Name**
{noaccessory=true}
```go
func (*Gzip) Name() string
```

**Apply**
{noaccessory=true}
```go
func (g *Gzip) Apply(h http.Handler) http.Handler
```

### Recover

{noaccessory=true}
```go
type Recover struct {
    // contains filtered or unexported fields
}
```

Recover middleware handles panics in HTTP handlers.

#### Methods

**Name**
{noaccessory=true}
```go
func (*Recover) Name() string
```

**Apply**
{noaccessory=true}
```go
func (m *Recover) Apply(h http.Handler) http.Handler
```

### Authentication

{noaccessory=true}
```go
type Authentication struct {
    Authenticator Authenticator
}
```

Authentication middleware handles request authentication using an Authenticator implementation.

#### Methods

**Name**
{noaccessory=true}
```go
func (*Authentication) Name() string
```

**Apply**
{noaccessory=true}
```go
func (m *Authentication) Apply(h http.Handler) http.Handler
```

## Functions

### NewColorScheme

{noaccessory=true}
```go
func NewColorScheme(key string) *ColorScheme
```

NewColorScheme creates a new ColorScheme middleware with the specified cookie key.

### DefaultColorScheme

{noaccessory=true}
```go
var DefaultColorScheme = &ColorScheme{Key: "pacis_color_scheme", UseHeaders: false}
```

DefaultColorScheme provides a default ColorScheme middleware instance.

### GetColorScheme

{noaccessory=true}
```go
func GetColorScheme(ctx context.Context) string
```

GetColorScheme retrieves the color scheme (theme) from the provided context. It expects the context to have a value associated with the key "theme" of type string.

### NewLocale

{noaccessory=true}
```go
func NewLocale(key string, bundle *i18n.Bundle, defaultlang language.Tag) *Locale
```

NewLocale creates a new Locale middleware with the specified cookie key, i18n bundle, and default language.

### GetLocalizer

{noaccessory=true}
```go
func GetLocalizer(ctx context.Context) *i18n.Localizer
```

GetLocalizer retrieves the localizer struct from the provided context. It expects the context to have a value associated with the key "localizer" of type *i18n.Localizer.

### GetLocale

{noaccessory=true}
```go
func GetLocale(ctx context.Context) *language.Tag
```

GetLocale retrieves the locale value from the provided context. It expects the context to have a value associated with the key "locale" of type *language.Tag.

### NewCache

{noaccessory=true}
```go
func NewCache(dur time.Duration) *Cache
```

NewCache creates a new Cache middleware with the specified duration for the max-age directive.

### NewLogger

{noaccessory=true}
```go
func NewLogger(logger *slog.Logger) *Logger
```

NewLogger creates a new Logger middleware with the specified slog.Logger instance.

### DefaultGzip

{noaccessory=true}
```go
func DefaultGzip(dev bool) *Gzip
```

DefaultGzip creates a new Gzip middleware. When dev is true, compression is disabled.

### NewRecover

{noaccessory=true}
```go
func NewRecover(logger *slog.Logger, callback func(any)) *Recover
```

NewRecover creates a new Recover middleware with the specified logger and optional callback function.

### NewAuthentication

{noaccessory=true}
```go
func NewAuthentication(authr Authenticator) *Authentication
```

NewAuthentication creates a new Authentication middleware with the specified Authenticator implementation.

### GetUser

{noaccessory=true}
```go
func GetUser[T any](ctx context.Context) T
```

GetUser retrieves the authenticated user from the provided context. The type parameter T specifies the expected user type.