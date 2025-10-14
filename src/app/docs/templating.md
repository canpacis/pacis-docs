# Templating

This guide explains how to construct pages from scratch with a functional approach to building type-safe HTML templates.

## Core Concepts

### 1. HTML Nodes

Everything in Pacis is a `Node`. The library provides functions for every HTML element that return nodes:

```go
P(...)        // <p>
Div(...)      // <div>
A(...)        // <a>
H1(...)       // <h1>
Span(...)     // <span>
```

Nodes are composed by passing arguments to element functions. These arguments can be attributes, classes, text content, or child nodes.

### 2. Building HTML Elements

Create elements by calling their corresponding function and passing content as arguments:

```go
// Simple element with text
P(Text("Hello world"))

// Element with multiple children
Div(
    H1(Text("Title")),
    P(Text("Content")),
)
```

Children are processed in order and appended to the element. The element function returns a Node that can be composed into parent elements.

### 3. Attributes

Attributes are applied using helper functions:

```go
A(
    Href("https://example.com"),
    Target("_blank"),
    Class("text-blue-500"),
    
    Text("Link"),
)
```

Common attribute functions include:

- **`Href(url)`** - Sets the `href` attribute on links and forms
- **`Class(classes)`** - Applies CSS classes (space-separated)
- **`ID(id)`** - Sets the element's `id`
- **`Src(path)`** - Sets image/script source
- **`Target(target)`** - Sets link target (`_blank`, etc.)

### 4. Conditional Rendering

Use `If` for conditional nodes and `IfFn` for deferring the execution 
of the code, this is useful for checking `nil` values and rendering
if values exist.

```go
If(true, func() Item {
  return P(...)
})

IfFn(value != nil, func() Item {
  return P(Text(value.Title))
})

// Conditional attribute
If(forwards, Class("ml-auto"))
```

`If` and `IfFn` function return an `html.Item`, meaning you can use them as child nodes & attributes.

### 5. Looping and Mapping

Use `Map` to render lists of items:

```go
Ul(
    Class("flex flex-col gap-2"),

    Map(items, func(item ItemType) Node {
        return Li(
            P(Text(item.Name))
        )
    })
)
```

`Map` takes a slice and a function that transforms each item into a Node. The resulting nodes are appended to the parent.

> Use `MapIdx` to get the index for each element as well.
>
> `MapIdx(items, func(item ItemType, i int) Node {})`

### 6. Text Content

Insert text using the `Text()` function:

```go
P(Text("This is content"))
Span(Text("More text"))
```

Text nodes are automatically escaped for safety.

### 7. Fragments

Group multiple nodes without a wrapper element using `Fragment()`:

```go
Fragment(
    P(Text("First")),
    P(Text("Second")),
    P(Text("Third"))
)
```

This is useful in `Map` callbacks or when you need multiple top-level elements.

### 8. Custom Elements and Components

Create reusable components as functions returning Nodes:

```go
func Card(title, description, href string) Node {
	return A(
		Href(href),
		Class("block p-4 rounded-lg border hover:bg-slate-50 transition-colors"),

		H3(Class("font-semibold mb-2"), Text(title)),
		P(Class("text-sm text-slate-600"), Text(description)),
	)
}

// ...

Card("Title", "Description", "https://pacis.canpacis.com")
```

Pass data as function parameters and return a Node.

## Execution Flow

1. **Element Creation**: Call element functions like `Div()`, `P()`, etc.
2. **Attribute Application**: Pass attributes and classes in any order
3. **Child Composition**: Pass child nodes and text as arguments
4. **Context Resolution**: Deferred attributes are computed during rendering when the request context is available
5. **HTML Generation**: The complete node tree is rendered to HTML
6. **Response**: HTML is sent to the browser as part of the HTTP response

## Type Safety

Pacis catches most errors at compile time. For example, passing an invalid argument to an element function will fail to compile. This prevents many common template bugs found in string-based templating languages.
