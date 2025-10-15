# Rendering

Up until now, we have been talking about building a page, layouts and components with Pacis' [templating language](/getting-started/templating). But how are they rendered? When are they rendered?

Everything you write, Pacis attempts to prerender it. That's it! The context in which you are building your UIs is one without an HTTP request. Everytime you call a `Div()` or a `Class()` function, Pacis creates a data structure that can be rendered in an arbitrary amount of chunks. If these chunks are not context bound, meaning they do not need any data from outside (like an HTTP request), they generate their content immediately. Anything that needs contextual data is rendered upon a request.

Consider the following example:

{noaccessory=true}
```go
P(Text("Hello"))
```

This piece of UI is purely static and can be rendered immediately. This means as soon as you start your server process, Pacis will generate this.


## Components

A  `Component` is a runtime primitive that the `html` package provides to you to create your rendering boundaries. This is the part Pacis _cannot_ immediately render. **Components need a request to render**.

Take the previous example. Let's say we want to greet everyone by their name instead of a generic "Hello" message. If we want to get a `name` value from, let's say, url search parameters, we are going to need a context. Components provide you that context! The definition of `Component` is literally like so in the `html` package:

{file="html/node.go"}
```go
type Component func(context.Context) Node
```

So with a component, we can render a dynamic UI:

{noaccessory=true}
```go
Component(func (ctx context.Context) Node {
  name := getNameFromURLParams(ctx)
  return P(Textf("Hello %s", name))
})
```

This might seem clunky at first but by following a few simple conventions, you will understand what this can do for you.

Remember, most of your context bound UI live with each other. So you create components out of them. Let's make a component of our greeter UI.

{noaccessory=true}
```go
func Greeter(ctx context.Context) {
  name := getNameFromURLParams(ctx)
  return P(Textf("Hello %s", name))
}

// Use somewhere without a context

Component(Greeter)
```

When you see a function that accepts a context, you'll quickly understand its execution environment. _Functions with contexts render on demand with a request!_

## Pages & Layouts

You define your pages & layouts similar to composites:

{noaccessory=true}
```go
func Layout(head, children Node) Node {
  ...
}

func Page() Node {
  ...
}
```

Notice they don't have a context parameter, meaning Pacis will render them ahead of time. When you register a route with a page and a layout 
Pacis will build the static parts of that route.