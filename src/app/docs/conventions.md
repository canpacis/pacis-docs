# Conventions

## Importing The `html` Package

I prefer to _dot_ import the `github.com/canpacis/pacis/html` package, which is unusual for Go. I must say that using this import method should probably be avoided in any program. Having said that, the `html` package is pretty much designed to be used in this way and I believe there are situations in which using this feature is fine. 

Importing the package this way lets you use its functions without prefixing them with the `html.` namespace and makes it easier to develop.

Go's `staticcheck` linter will probably yell at you for doing this[^staticcheck]. Updating your `staticcheck.conf` file to include the package name will resolve the issue. The template repository comes with a configured file for you.

```conf
dot_import_whitelist = ["github.com/canpacis/pacis/html"]
```

## Project Structure

The template comes with the following structure.

```
└── root/
    ├── main.go
    ├── prod.go
    ├── dev.go
    ├── public/
    │   └── favicon.ico
    └── src/
        ├── app/
        │   └── page.go
        └── web/
            ├── style.css/
            └── main.ts/
```

A typical Pacis app would follow this by adding any ui or business logic related code inside the `src/app` directory. Your public files should go in the `public` directory and your frontend assets like scripts and styles should be in `src/app/web` directory.

In the end, your app might look like this:

```
└── root/
    ├── main.go
    ├── prod.go
    ├── dev.go
    ├── public/
    │   ├── favicon.ico
    │   ├── logo.png
    │   ├── sitemap.xml
    │   └── robots.txt
    └── src/
        ├── app/
        │   ├── page.go
        │   ├── layout.go
        │   ├── about.go
        │   └── blog.go
        ├── components/
        │   ├── button/
        │   │   └── button.go
        │   └── card/
        │       └── card.go
        ├── icons/
        │   ├── icon.go
        │   └── icons.go
        └── web/
            ├── main.ts
            ├── stream.ts
            ├── alpine.ts
            └── global.css
```

## Components

Notice that above, each component has its own folder. It is better to namespace the components by their package. Because we import the `html` package via a dot import, common components like **buttons** can get name collisions. I avoid them by giving components their own packages.

Provide a `New() *Component` function in your package for creating a new instance of your component.

{file="src/components/button/button.go"}
```go
package button

import . "github.com/canpacis/pacis/html"

func New() Node {
  ...
}
```

## Element Items

HTML elements defined in the `html` package accept a variadic item of `html.Iten` interfaces. This includes HTML attributes like `class`, `id` and `src` while also acommodating other HTML elements and nodes like, `Text` and `Fragment`.

Having every item in the same argument list, especially when building a complex, nested UI, can get really confusing really quick. To minimize, I advise keeping your HTML attributes from your child nodes.

No ❌

{noaccessory=true}
```go
Div(
  ID("app"),
  Class("h-screen"),
  P(...),
  Section(...)
)
```

Do **not** mix attributes and child nodes.

No ❌

{noaccessory=true}
```go
Div(
  ID("app"),
  
  P(...),
  Class("h-screen"),
  Section(...)
)
```

Yes ✅

{noaccessory=true}
```go
Div(
  ID("app"),
  Class("h-screen"),
  
  P(...),
  Section(...)
)
```

Ocasionally inlining your element with a 1 or 2 attributes/child nodes is fine.

OK ✅

{noaccessory=true}
```go
P(Class(...), Text(...))
```


[^staticcheck]:[https://staticcheck.dev/docs/checks/#ST1001](https://staticcheck.dev/docs/checks/#ST1001)