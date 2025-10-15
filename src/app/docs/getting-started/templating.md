# Templating

This guide explains how to construct HTML pages in Go using Pacis’s functional, type-safe templating API.

## Items and Nodes

Everything in Pacis is an Item, the most general type in the system.
An Item can be either:

- a Node, which renders HTML (e.g. elements, text, fragments, components), or
- a Property, which applies behavior or attributes to elements (e.g. Class, Href, ID).

When you pass arguments into an element function like `Div(...)`, Pacis automatically figures out which arguments are nodes (content) and which are properties (attributes).

{noaccessory=true}
```go
Div(
    Class("p-4 bg-gray-50"), // Property
    H1(Text("Title")),       // Node
    P(Text("Content")),      // Node
)
```

Each element function accepts any combination of Items in any order.

## HTML Elements

Every HTML tag is represented by a Go function that returns a Node:

{noaccessory=true}
```go
P(...)      // <p>
Div(...)    // <div>
A(...)      // <a>
H1(...)     // <h1>
Span(...)   // <span>
```

You compose elements by nesting them, and everything is statically typed and compiled, no string parsing or templates.

## Attributes

Attributes are defined as **properties**, using helper functions:

{noaccessory=true}
```go
A(
    Href("https://example.com"),
    Target("_blank"),
    Class("text-blue-500 underline"),
    Text("Link"),
)
```

Common property helpers include:

| Function         | Description                           |
|--------|------:|
| `Href(url)`      | Sets the `href` attribute             |
| `Class(classes)` | Applies CSS classes (space-separated) |
| `ID(id)`         | Sets the element’s `id`               |
| `Src(path)`      | Sets image or script `src`            |
| `Target(value)`  | Sets link `target` (`_blank`, etc.)   |


## Conditional Rendering

Pacis includes utilities for conditionally returning nodes or attributes:

{noaccessory=true}
```go
If(condition, P(Text("Visible if true")))

IfFn(user != nil, func() Item {
    return Span(Text(user.Name))
})

// Conditional attribute
Div(
    If(isRightAligned, Class("ml-auto")),
)
```

* `If(cond, item)` returns the item if `cond` is true, otherwise an empty fragment.
* `IfFn(cond, fn)` defers evaluation, useful when the item requires computation or depends on a possibly `nil` value.

Both functions return an `Item`, so they can be used as children **or** as attributes.

## Looping and Mapping

Use `Map` or `MapIdx` to render slices:

{noaccessory=true}
```go
Ul(
    Class("flex flex-col gap-2"),
    Map(items, func(item Todo) Node {
        return Li(Text(item.Title))
    }),
)
```

- `Map(slice, fn)` iterates over items and returns a `Fragment` of all rendered nodes.
- `MapIdx(slice, fn)` also passes the index of each element.
- `Iter` and `Iter2` support `iter.Seq` and `iter.Seq2` for generator-style iteration.

## Text Content

Insert plain text using `Text()` or formatted text with `Textf()`:

{noaccessory=true}
```go
P(Text("Safe text"))
Span(Textf("Hello, %s!", name))
```

Text is always HTML-escaped for safety.

## Fragments

Use `Fragment()` to group multiple nodes without a wrapper element:

{noaccessory=true}
```go
Fragment(
    P(Text("First")),
    P(Text("Second")),
)
```

Fragments are useful when returning multiple sibling nodes from a function or loop.

## Composites

Composites are plain Go functions returning `Node`s:

{noaccessory=true}
```go
func Card(title, desc, href string) Node {
    return A(
        Href(href),
        Class("block p-4 border rounded-lg hover:bg-slate-50"),
        H3(Class("font-semibold mb-2"), Text(title)),
        P(Class("text-sm text-slate-600"), Text(desc)),
    )
}

// usage:
Card("Title", "Description", "/details")
```

Since they are just functions, composites can take any data, conditionally render children, or compose other composites.
