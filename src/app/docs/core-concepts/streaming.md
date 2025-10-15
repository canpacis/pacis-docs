# Streaming

Streaming with Pacis really easy. Becuase the rendering boundary is defined very well and we have primitives like `html.Component`, it is trivial to add asynchronous rendering to your apps.

The `server` package provides an `Async()` function that returns an `html.Component`

{file="server/items.go"}
```go
func Async(html.Component, html.Node) html.Component 
```

> It takes a `Component` to stream and a `Node` to replace it while it is loading.

A simple example that streams a "slow" component. The surrounding content will be rendered immediately and our slow component will be late to the game.

{noaccessory=true}
```go
func Page() Node {
  Div(
    P(Text("This will render immediately")),
    server.Async(Slow, P(Text("Loading..."))),
    P(Text("This will also render immediately")),
  )
}

// This is a component now
func Slow(context.Context) Node {
  // Simulate slow action
  time.Sleep(time.Seconds * 5)

  return P(Text("This will render late"))
}
```

Try replacing the `Async()` call with just a component and see how slow the page is loading.

{noaccessory=true}
```go
func Page() Node {
  Div(
    P(Text("This will render immediately")),
    Component(Slow),
    P(Text("This will also render immediately")),
  )
}
```