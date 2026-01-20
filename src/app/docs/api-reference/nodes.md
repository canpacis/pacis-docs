# Nodes

## Helper Nodes

### Map & MapIdx

#### Map `Map[E any](s []E, fn func(E) Node) Node`

Applies a function to each element in a slice, returning a fragment containing all resulting child nodes.

**Parameters**

* `s []E` — Input slice of elements.
* `fn func(E) Node` — Function applied to each element that produces a `Node`.

**Returns**

* `Node` — A `Fragment` containing all nodes produced by mapping `fn` over `s`.

**Usage Example**

{noaccessory=true}
```go
html.Map(users, func(u User) Node {
	return html.Li(html.Text(u.Name))
})
```

#### MapIdx `MapIdx[E any](s []E, fn func(E, int) Node) Node`

Same as `Map`, but also passes the index of each element to the mapping function.

**Parameters**

* `s []E` — Input slice.
* `fn func(E, int) Node` — Function that receives the element and its index.

**Returns**

* `Node` — A `Fragment` containing all nodes produced by mapping `fn` over `s` with indices.

**Usage Example**

{noaccessory=true}
```go
html.MapIdx(items, func(item string, i int) Node {
	return html.Li(html.Text(fmt.Sprintf("%d. %s", i+1, item)))
})
```

### Iter & Iter2

#### Iter `Iter[T any](s iter.Seq[T], fn func(T) Node) Node`

Iterates over a generic sequence (`iter.Seq`) and applies a function to each element, returning a fragment of nodes.

**Parameters**

* `s iter.Seq[T]` — Sequence to iterate over.
* `fn func(T) Node` — Function applied to each yielded element.

**Returns**

* `Node` — A `Fragment` containing all nodes produced by `fn`.

**Usage Example**

{noaccessory=true}
```go
html.Iter(stream, func(value string) Node {
	return html.P(html.Text(value))
})
```

#### Iter2 `Iter2[K any, V any](s iter.Seq2[K, V], fn func(K, V) Node) Node`

Iterates over a key/value sequence, applying a function to each pair.

**Parameters**

* `s iter.Seq2[K, V]` — Sequence yielding key–value pairs.
* `fn func(K, V) Node` — Function applied to each key/value pair.

**Returns**

* `Node` — A `Fragment` containing all nodes produced by `fn`.

**Usage Example**

{noaccessory=true}
```go
html.Iter2(settings, func(k string, v any) Node {
	return html.Div(
		html.Strong(html.Text(k+":")),
		html.Span(html.Text(fmt.Sprint(v))),
	)
})
```

### Conditional Rendering

#### If `If(cond bool, item Item) Item`

Returns `item` if `cond` is true, otherwise returns an empty fragment.
Useful for conditional rendering in declarative markup.

**Parameters**

* `cond bool` — The condition to evaluate.
* `item Item` — The node to include if true.

**Returns**

* `Item` — Either `item` or an empty `Fragment`.

**Usage Example**

{noaccessory=true}
```go
html.If(user.LoggedIn, html.P(html.Text("Welcome back!")))
```

#### IfFn `IfFn(cond bool, fn func() Item) Item`

Conditionally calls and returns the result of a function if `cond` is true.

**Parameters**

* `cond bool` — The condition to check.
* `fn func() Item` — Function returning a node to render if true.

**Returns**

* `Item` — Result of `fn()` or an empty `Fragment`.

**Usage Example**

{noaccessory=true}
```go
html.IfFn(len(posts) > 0, func() Item {
	return html.Ul(html.Map(posts, func(p Post) Node {
		return html.Li(html.Text(p.Title))
	}))
})
```

### Switch-like Rendering

#### SwitchCase `type SwitchCase[T comparable] struct`

Represents a case for a switch-like conditional structure.

**Fields**

* `Value []T` — Values that trigger this case.
* `Fn func() Item` — Function that returns the node when matched.
* `Default bool` — Marks this case as the default.

#### Switch `Switch[T comparable](expr T, cases ...*SwitchCase[T]) Item`

Evaluates multiple cases against an expression, returning the node from the first match or a default case if present.

**Parameters**

* `expr T` — The value to test.
* `cases ...*SwitchCase[T]` — The cases to evaluate.

**Returns**

* `Item` — The matching node or an empty fragment.

**Usage Example**

{noaccessory=true}
```go
html.Switch(status,
	html.Case(html.P(html.Text("OK")), 200),
	html.Case(html.P(html.Text("Not Found")), 404),
	html.Default(html.P(html.Text("Unknown"))),
)
```

#### Case `Case[T comparable](item Item, v ...T) *SwitchCase[T]`

Creates a case that returns a static node when any of the values match.

**Parameters**

* `item Item` — The node to render.
* `v ...T` — Values to match.

**Returns**

* `*SwitchCase[T]` — The constructed case.


#### CaseFn `CaseFn[T comparable](fn func() Item, v ...T) *SwitchCase[T]`

Creates a case that runs a function if the case matches.

**Parameters**

* `fn func() Item` — Function producing the node.
* `v ...T` — Values to match.

**Returns**

* `*SwitchCase[T]` — The constructed case.


#### Default `Default[T comparable](item Item) *SwitchCase[T]`

Creates a default case that always returns the provided static node.

**Parameters**

* `item Item` — Node to render if no cases match.

**Returns**

* `*SwitchCase[T]` — Default case.


#### DefaultFn `DefaultFn[T comparable](fn func() Item) *SwitchCase[T]`

Creates a default case that executes the provided function when matched.

**Parameters**

* `fn func() Item` — Function producing the node.

**Returns**

* `*SwitchCase[T]` — Default case.


### JSON Rendering

#### JSONNode `type JSONNode struct`

Represents a node that serializes arbitrary data as JSON, with optional indentation.

**Fields**

* `Data any` — The data to encode as JSON.
* `Indent string` — Indentation for formatted output.


#### WithIndent `func (n *JSONNode) WithIndent(indent string) *JSONNode`

Sets indentation for the JSON output and returns the node itself (builder-style).

**Parameters**

* `indent string` — The indentation prefix (e.g. `"  "`).

**Returns**

* `*JSONNode` — The same node with updated indentation.

#### Render `func (n *JSONNode) Render(w ChunkWriter) error`

Implements rendering logic for JSON output. Serializes `Data` to JSON and writes it to the writer.

#### JSON `JSON(data any) *JSONNode`

Creates a new `JSONNode` that will serialize the given data as JSON.

**Usage Example**

{noaccessory=true}
```go
html.Pre(html.JSON(map[string]any{
	"name": "Alice",
	"age":  28,
}).WithIndent("\t"))
```

### Unsafe Raw HTML

#### `type RawUnsafe string`

Represents raw HTML content inserted directly into the output without escaping.
:triangle-alert: Use cautiously — this bypasses HTML escaping and can introduce XSS vulnerabilities.

**Implements**

* `Item`
* `Node`

#### Render `func (t RawUnsafe) Render(w ChunkWriter) error`

Writes raw HTML directly to the output stream.

**Usage Example**

{noaccessory=true}
```go
html.RawUnsafe("<b>Unescaped bold text</b>")
```


## HTML Elements

### A 

`A(items ...Item) *Element`

Creates an HTML `<a>` (anchor) element for hyperlinks.

[MDN Reference](https://developer.mozilla.org/en-US/docs/Web/HTML/Element/a)

---

### Abbr 

`Abbr(items ...Item) *Element`

Creates an HTML `<abbr>` element representing an abbreviation or acronym.

[MDN Reference](https://developer.mozilla.org/en-US/docs/Web/HTML/Element/abbr)

---

### Address 

`Address(items ...Item) *Element`

Creates an HTML `<address>` element representing contact information.

[MDN Reference](https://developer.mozilla.org/en-US/docs/Web/HTML/Element/address)

---

### Area 

`Area(items ...Item) *Element`

Creates an HTML `<area>` element for defining clickable regions in an image map.

[MDN Reference](https://developer.mozilla.org/en-US/docs/Web/HTML/Element/area)

---

### Article 

`Article(items ...Item) *Element`

Creates an HTML `<article>` element representing self-contained content.

[MDN Reference](https://developer.mozilla.org/en-US/docs/Web/HTML/Element/article)

---

### Aside 

`Aside(items ...Item) *Element`

Creates an HTML `<aside>` element for tangentially related content or sidebars.

[MDN Reference](https://developer.mozilla.org/en-US/docs/Web/HTML/Element/aside)

---

### Audio 

`Audio(items ...Item) *Element`

Creates an HTML `<audio>` element to embed sound content.

[MDN Reference](https://developer.mozilla.org/en-US/docs/Web/HTML/Element/audio)

---

### B 

`B(items ...Item) *Element`

Creates an HTML `<b>` element for stylistically offset text (without semantic importance).

[MDN Reference](https://developer.mozilla.org/en-US/docs/Web/HTML/Element/b)

---

### Base 

`Base(items ...Item) *Element`

Creates an HTML `<base>` element specifying the base URL for relative URLs.

[MDN Reference](https://developer.mozilla.org/en-US/docs/Web/HTML/Element/base)

---

### Bdi 

`Bdi(items ...Item) *Element`

Creates an HTML `<bdi>` element isolating text for bidirectional formatting.

[MDN Reference](https://developer.mozilla.org/en-US/docs/Web/HTML/Element/bdi)

---

### Bdo 

`Bdo(items ...Item) *Element`

Creates an HTML `<bdo>` element to override text direction.

[MDN Reference](https://developer.mozilla.org/en-US/docs/Web/HTML/Element/bdo)

---

### Blockquote 

`Blockquote(items ...Item) *Element`

Creates an HTML `<blockquote>` element for extended quotations.

[MDN Reference](https://developer.mozilla.org/en-US/docs/Web/HTML/Element/
---blockquote)


### Body 

`Body(items ...Item) *Element`

Creates an HTML `<body>` element representing the main document content.

[MDN Reference](https://developer.mozilla.org/en-US/docs/Web/HTML/Element/body)

---

### Br 

`Br(items ...Item) *Element`

Creates an HTML `<br>` element for line breaks.

[MDN Reference](https://developer.mozilla.org/en-US/docs/Web/HTML/Element/br)

---

### Button 

`Button(items ...Item) *Element`

Creates an HTML `<button>` element representing an interactive button.

[MDN Reference](https://developer.mozilla.org/en-US/docs/Web/HTML/Element/button)

---

### Canvas 

`Canvas(items ...Item) *Element`

Creates an HTML `<canvas>` element for drawing graphics.

[MDN Reference](https://developer.mozilla.org/en-US/docs/Web/HTML/Element/canvas)

---

### Caption 

`Caption(items ...Item) *Element`

Creates an HTML `<caption>` element defining a table’s title or caption.

[MDN Reference](https://developer.mozilla.org/en-US/docs/Web/HTML/Element/caption)

---

### Cite 

`Cite(items ...Item) *Element`

Creates an HTML `<cite>` element for citing a creative work.

[MDN Reference](https://developer.mozilla.org/en-US/docs/Web/HTML/Element/cite)

---

### Code 

`Code(items ...Item) *Element`

Creates an HTML `<code>` element for displaying code fragments.

[MDN Reference](https://developer.mozilla.org/en-US/docs/Web/HTML/Element/code)

---

### Col 

`Col(items ...Item) *Element`

Creates an HTML `<col>` element defining table columns.

[MDN Reference](https://developer.mozilla.org/en-US/docs/Web/HTML/Element/col)

---

### Colgroup 

`Colgroup(items ...Item) *Element`

Creates an HTML `<colgroup>` element grouping table columns.

[MDN Reference](https://developer.mozilla.org/en-US/docs/Web/HTML/Element/colgroup)

---

### DataEl 

`DataEl(items ...Item) *Element`

Creates an HTML `<data>` element linking human-readable and machine-readable content.

[MDN Reference](https://developer.mozilla.org/en-US/docs/Web/HTML/Element/data)

---

### Datalist 

`Datalist(items ...Item) *Element`

Creates an HTML `<datalist>` element defining predefined options for inputs.

[MDN Reference](https://developer.mozilla.org/en-US/docs/Web/HTML/Element/datalist)

---

### Dd 

`Dd(items ...Item) *Element`

Creates an HTML `<dd>` element defining description list details.

[MDN Reference](https://developer.mozilla.org/en-US/docs/Web/HTML/Element/dd)

---

### Del 

`Del(items ...Item) *Element`

Creates an HTML `<del>` element for deleted text.

[MDN Reference](https://developer.mozilla.org/en-US/docs/Web/HTML/Element/del)

---

### Details 

`Details(items ...Item) *Element`

Creates an HTML `<details>` element for collapsible disclosure widgets.

[MDN Reference](https://developer.mozilla.org/en-US/docs/Web/HTML/Element/details)

---

### Dfn 

`Dfn(items ...Item) *Element`

Creates an HTML `<dfn>` element marking a term being defined.

[MDN Reference](https://developer.mozilla.org/en-US/docs/Web/HTML/Element/dfn)

---

### Dialog 

`Dialog(items ...Item) *Element`

Creates an HTML `<dialog>` element for modal or non-modal dialogs.

[MDN Reference](https://developer.mozilla.org/en-US/docs/Web/HTML/Element/dialog)

---

### Div 

`Div(items ...Item) *Element`

Creates an HTML `<div>` element as a generic container.

[MDN Reference](https://developer.mozilla.org/en-US/docs/Web/HTML/Element/div)

---

### Dl 

`Dl(items ...Item) *Element`

Creates an HTML `<dl>` element representing a description list.

[MDN Reference](https://developer.mozilla.org/en-US/docs/Web/HTML/Element/dl)

---

### Dt 

`Dt(items ...Item) *Element`

Creates an HTML `<dt>` element for a term in a description list.

[MDN Reference](https://developer.mozilla.org/en-US/docs/Web/HTML/Element/dt)

---

### Em 

`Em(items ...Item) *Element`

Creates an HTML `<em>` element for emphasizing text.

[MDN Reference](https://developer.mozilla.org/en-US/docs/Web/HTML/Element/em)

---

### Fieldset 

`Fieldset(items ...Item) *Element`

Creates an HTML `<fieldset>` element used to group several form controls and labels.

[MDN Reference](https://developer.mozilla.org/en-US/docs/Web/HTML/Element/fieldset)

---

### Figcaption 

`Figcaption(items ...Item) *Element`

Creates an HTML `<figcaption>` element that provides a caption or legend for its parent `<figure>`.

[MDN Reference](https://developer.mozilla.org/en-US/docs/Web/HTML/Element/
---figcaption)


### Figure 

`Figure(items ...Item) *Element`

Creates an HTML `<figure>` element representing self-contained content, optionally with a caption.

[MDN Reference](https://developer.mozilla.org/en-US/docs/Web/HTML/Element/figure)

---

### Footer 

`Footer(items ...Item) *Element`

Creates an HTML `<footer>` element representing the footer of a section or page.

[MDN Reference](https://developer.mozilla.org/en-US/docs/Web/HTML/Element/footer)

---

### Form 

`Form(items ...Item) *Element`

Creates an HTML `<form>` element representing a section containing interactive controls for data submission.


[MDN Reference](https://developer.mozilla.org/en-US/docs/Web/HTML/Element/form)


### H1 

`H1(items ...Item) *Element`

Creates an HTML `<h1>` element for a top-level heading.

[MDN Reference](https://developer.mozilla.org/en-US/docs/Web/HTML/Reference/Elements/Heading_Elements)


### H2 

`H2(items ...Item) *Element`

Creates an HTML `<h2>` element for a second-level heading.

[MDN Reference](https://developer.mozilla.org/en-US/docs/Web/HTML/Reference/Elements/Heading_Elements)


### H3 

`H3(items ...Item) *Element`

Creates an HTML `<h3>` element for a third-level heading.

[MDN Reference](https://developer.mozilla.org/en-US/docs/Web/HTML/Reference/Elements/Heading_Elements)


### H4 

`H4(items ...Item) *Element`

Creates an HTML `<h4>` element for a fourth-level heading.

[MDN Reference](https://developer.mozilla.org/en-US/docs/Web/HTML/Reference/Elements/Heading_Elements)


### H5 

`H5(items ...Item) *Element`

Creates an HTML `<h5>` element for a fifth-level heading.

[MDN Reference](https://developer.mozilla.org/en-US/docs/Web/HTML/Reference/Elements/Heading_Elements)


### H6 

`H6(items ...Item) *Element`

Creates an HTML `<h6>` element for a sixth-level heading.

[MDN Reference](https://developer.mozilla.org/en-US/docs/Web/HTML/Reference/Elements/Heading_Elements)


### Head 

`Head(items ...Item) *Element`

Creates an HTML `<head>` element containing metadata and links for the document.

[MDN Reference](https://developer.mozilla.org/en-US/docs/Web/HTML/Element/head)

---

### Header 

`Header(items ...Item) *Element`

Creates an HTML `<header>` element for introductory or navigational content.

[MDN Reference](https://developer.mozilla.org/en-US/docs/Web/HTML/Element/header)

---

### Hgroup 

`Hgroup(items ...Item) *Element`

Creates an HTML `<hgroup>` element grouping heading elements and related content.

[MDN Reference](https://developer.mozilla.org/en-US/docs/Web/HTML/Element/hgroup)

---

### Hr 

`Hr(items ...Item) *Element`

Creates an HTML `<hr>` element representing a thematic break between sections.

[MDN Reference](https://developer.mozilla.org/en-US/docs/Web/HTML/Element/hr)

---

### Html 

`Html(items ...Item) *Element`

Creates an HTML `<html>` element representing the root element of a document.

[MDN Reference](https://developer.mozilla.org/en-US/docs/Web/HTML/Element/html)

---

### I 

`I(items ...Item) *Element`

Creates an HTML `<i>` element representing text set off for idiomatic or technical reasons.

[MDN Reference](https://developer.mozilla.org/en-US/docs/Web/HTML/Element/i)

---

### Iframe 

`Iframe(items ...Item) *Element`

Creates an HTML `<iframe>` element embedding another HTML page within the current one.

[MDN Reference](https://developer.mozilla.org/en-US/docs/Web/HTML/Element/iframe)

---

### Img 

`Img(items ...Item) *Element`

Creates an HTML `<img>` element embedding an image.

[MDN Reference](https://developer.mozilla.org/en-US/docs/Web/HTML/Element/img)

---

### Input 

`Input(items ...Item) *Element`

Creates an HTML `<input>` element for data entry controls.

[MDN Reference](https://developer.mozilla.org/en-US/docs/Web/HTML/Element/input)

---

### Ins 

`Ins(items ...Item) *Element`

Creates an HTML `<ins>` element representing inserted text.

[MDN Reference](https://developer.mozilla.org/en-US/docs/Web/HTML/Element/ins)

---

### Kbd 

`Kbd(items ...Item) *Element`

Creates an HTML `<kbd>` element denoting user input text.

[MDN Reference](https://developer.mozilla.org/en-US/docs/Web/HTML/Element/kbd)

---

### Label 

`Label(items ...Item) *Element`

Creates an HTML `<label>` element labeling a form control.

[MDN Reference](https://developer.mozilla.org/en-US/docs/Web/HTML/Element/label)

---

### Legend 

`Legend(items ...Item) *Element`

Creates an HTML `<legend>` element that provides a caption for a `<fieldset>`.

[MDN Reference](https://developer.mozilla.org/en-US/docs/Web/HTML/Element/legend)

---

### Li 

`Li(items ...Item) *Element`

Creates an HTML `<li>` element representing a list item.

[MDN Reference](https://developer.mozilla.org/en-US/docs/Web/HTML/Element/li)

---

### Link 

`Link(items ...Item) *Element`

Creates an HTML `<link>` element linking an external resource (like stylesheets or icons).

[MDN Reference](https://developer.mozilla.org/en-US/docs/Web/HTML/Element/link)

---

### Main 

`Main(items ...Item) *Element`

Creates an HTML `<main>` element representing the dominant content of the `<body>`.

[MDN Reference](https://developer.mozilla.org/en-US/docs/Web/HTML/Element/main)

---

### MapEl 

`MapEl(items ...Item) *Element`

Creates an HTML `<map>` element used with `<area>` to define image maps.

[MDN Reference](https://developer.mozilla.org/en-US/docs/Web/HTML/Element/map)

---

### Mark 

`Mark(items ...Item) *Element`

Creates an HTML `<mark>` element representing highlighted text for reference or emphasis.

[MDN Reference](https://developer.mozilla.org/en-US/docs/Web/HTML/Element/mark)

---

### Menu 

`Menu(items ...Item) *Element`

Creates an HTML `<menu>` element representing an unordered list of items.

[MDN Reference](https://developer.mozilla.org/en-US/docs/Web/HTML/Element/menu)

---

### Meta 

`Meta(items ...Item) *Element`

Creates an HTML `<meta>` element containing metadata that cannot be represented otherwise.

[MDN Reference](https://developer.mozilla.org/en-US/docs/Web/HTML/Element/meta)

---

### Meter 

`Meter(items ...Item) *Element`

Creates an HTML `<meter>` element representing a scalar or fractional measurement.

[MDN Reference](https://developer.mozilla.org/en-US/docs/Web/HTML/Element/meter)

---

### Nav 

`Nav(items ...Item) *Element`

Creates an HTML `<nav>` element representing navigation links.

[MDN Reference](https://developer.mozilla.org/en-US/docs/Web/HTML/Element/nav)

---

### Noscript 

`Noscript(items ...Item) *Element`

Creates an HTML `<noscript>` element defining content shown when scripts are disabled.

[MDN Reference](https://developer.mozilla.org/en-US/docs/Web/HTML/Element/noscript)

---

### Object 

`Object(items ...Item) *Element`

Creates an HTML `<object>` element embedding external resources such as images or plugins.

[MDN Reference](https://developer.mozilla.org/en-US/docs/Web/HTML/Element/object)

---

### Ol 

`Ol(items ...Item) *Element`

Creates an HTML `<ol>` element representing an ordered (numbered) list.

[MDN Reference](https://developer.mozilla.org/en-US/docs/Web/HTML/Element/ol)

---

### Optgroup 

`Optgroup(items ...Item) *Element`

Creates an HTML `<optgroup>` element grouping `<option>` elements within a `<select>`.

[MDN Reference](https://developer.mozilla.org/en-US/docs/Web/HTML/Element/optgroup)

---

### Option 

`Option(items ...Item) *Element`

Creates an HTML `<option>` element representing an item in a selection list.

[MDN Reference](https://developer.mozilla.org/en-US/docs/Web/HTML/Element/option)

---

### Output 

`Output(items ...Item) *Element`

Creates an HTML `<output>` element representing the result of a calculation or user action.

[MDN Reference](https://developer.mozilla.org/en-US/docs/Web/HTML/Element/output)

---

### P 

`P(items ...Item) *Element`

Creates an HTML `<p>` element representing a paragraph of text.

[MDN Reference](https://developer.mozilla.org/en-US/docs/Web/HTML/Element/p)

---

### Picture 

`Picture(items ...Item) *Element`

Creates an HTML `<picture>` element for responsive images using multiple sources.

[MDN Reference](https://developer.mozilla.org/en-US/docs/Web/HTML/Element/picture)

---

### Pre 

`Pre(items ...Item) *Element`

Creates an HTML `<pre>` element for preformatted text preserving whitespace.

[MDN Reference](https://developer.mozilla.org/en-US/docs/Web/HTML/Element/pre)

---

### Progress 

`Progress(items ...Item) *Element`

Creates an HTML `<progress>` element showing task completion progress.

[MDN Reference](https://developer.mozilla.org/en-US/docs/Web/HTML/Element/progress)

---

### Q 

`Q(items ...Item) *Element`

Creates an HTML `<q>` element representing a short inline quotation.

[MDN Reference](https://developer.mozilla.org/en-US/docs/Web/HTML/Element/q)

---

### Rp 

`Rp(items ...Item) *Element`

Creates an HTML `<rp>` element defining fallback parentheses for ruby annotations.

[MDN Reference](https://developer.mozilla.org/en-US/docs/Web/HTML/Element/rp)

---

### Rt 

`Rt(items ...Item) *Element`

Creates an HTML `<rt>` element specifying the ruby text of a ruby annotation.

[MDN Reference](https://developer.mozilla.org/en-US/docs/Web/HTML/Element/rt)

---

### Ruby 

`Ruby(items ...Item) *Element`

Creates an HTML `<ruby>` element representing ruby annotations, commonly for East Asian typography.

[MDN Reference](https://developer.mozilla.org/en-US/docs/Web/HTML/Element/ruby)

---

### S 

`S(items ...Item) *Element`

Creates an HTML `<s>` element rendering text with a strikethrough to indicate inaccuracy or irrelevance.

[MDN Reference](https://developer.mozilla.org/en-US/docs/Web/HTML/Element/s)

---

### Samp 

`Samp(items ...Item) *Element`

Creates an HTML `<samp>` element representing sample output from a computer program.

[MDN Reference](https://developer.mozilla.org/en-US/docs/Web/HTML/Element/samp)

---

### Script 

`Script(items ...Item) *Element`

Creates an HTML `<script>` element embedding or referencing executable code such as JavaScript.

[MDN Reference](https://developer.mozilla.org/en-US/docs/Web/HTML/Element/script)

---

### Search 

`Search(items ...Item) *Element`

Creates an HTML `<search>` element representing a section dedicated to search functionality.

[MDN Reference](https://developer.mozilla.org/en-US/docs/Web/HTML/Element/search)

---

### Section 

`Section(items ...Item) *Element`

Creates an HTML `<section>` element representing a standalone section of a document.

[MDN Reference](https://developer.mozilla.org/en-US/docs/Web/HTML/Element/section)

---

### Select 

`Select(items ...Item) *Element`

Creates an HTML `<select>` element representing a drop-down list of options.

[MDN Reference](https://developer.mozilla.org/en-US/docs/Web/HTML/Element/select)

---

### Slot 

`Slot(items ...Item) *Element`

Creates an HTML `<slot>` element, part of the Web Components API, acting as a placeholder for child content.


[MDN Reference](https://developer.mozilla.org/en-US/docs/Web/HTML/Element/slot)


### Small 

`Small(items ...Item) *Element`

Creates an HTML `<small>` element representing side-comments or fine print.

[MDN Reference](https://developer.mozilla.org/en-US/docs/Web/HTML/Element/small)

---

### Source 

`Source(items ...Item) *Element`

Creates an HTML `<source>` element specifying media sources for `<picture>`, `<audio>`, or `<video>`.

[MDN Reference](https://developer.mozilla.org/en-US/docs/Web/HTML/Element/source)

---

### Span 

`Span(items ...Item) *Element`

Creates an HTML `<span>` element as a generic inline container for phrasing content.

[MDN Reference](https://developer.mozilla.org/en-US/docs/Web/HTML/Element/span)

---

### Strong 

`Strong(items ...Item) *Element`

Creates an HTML `<strong>` element indicating strong importance or urgency.

[MDN Reference](https://developer.mozilla.org/en-US/docs/Web/HTML/Element/strong)

---

### Style 

`Style(items ...Item) *Element`

Creates an HTML `<style>` element containing CSS styling declarations.

[MDN Reference](https://developer.mozilla.org/en-US/docs/Web/HTML/Element/style)

---

### Sub 

`Sub(items ...Item) *Element`

Creates an HTML `<sub>` element displaying subscript text.

[MDN Reference](https://developer.mozilla.org/en-US/docs/Web/HTML/Element/sub)

---

### Summary 

`Summary(items ...Item) *Element`

Creates an HTML `<summary>` element defining a summary or caption for a `<details>` disclosure box.

[MDN Reference](https://developer.mozilla.org/en-US/docs/Web/HTML/Element/summary)

---

### Sup 

`Sup(items ...Item) *Element`

Creates an HTML `<sup>` element displaying superscript text.

[MDN Reference](https://developer.mozilla.org/en-US/docs/Web/HTML/Element/sup)

---

### Table 

`Table(items ...Item) *Element`

Creates an HTML `<table>` element representing tabular data.

[MDN Reference](https://developer.mozilla.org/en-US/docs/Web/HTML/Element/table)

---

### Tbody 

`Tbody(items ...Item) *Element`

Creates an HTML `<tbody>` element grouping the main content rows of a table.

[MDN Reference](https://developer.mozilla.org/en-US/docs/Web/HTML/Element/tbody)

---

### Td 

`Td(items ...Item) *Element`

Creates an HTML `<td>` element representing a data cell in a table.

[MDN Reference](https://developer.mozilla.org/en-US/docs/Web/HTML/Element/td)

---

### Template 

`Template(items ...Item) *Element`

Creates an HTML `<template>` element holding inert DOM fragments for later use.

[MDN Reference](https://developer.mozilla.org/en-US/docs/Web/HTML/Element/template)

---

### Textarea 

`Textarea(items ...Item) *Element`

Creates an HTML `<textarea>` element representing a multi-line text input control.

[MDN Reference](https://developer.mozilla.org/en-US/docs/Web/HTML/Element/textarea)

---

### Tfoot 

`Tfoot(items ...Item) *Element`

Creates an HTML `<tfoot>` element grouping the footer rows of a table.

[MDN Reference](https://developer.mozilla.org/en-US/docs/Web/HTML/Element/tfoot)

---

### Th 

`Th(items ...Item) *Element`

Creates an HTML `<th>` element representing a header cell in a table.

[MDN Reference](https://developer.mozilla.org/en-US/docs/Web/HTML/Element/th)

---

### Thead 

`Thead(items ...Item) *Element`

Creates an HTML `<thead>` element grouping header rows in a table.

[MDN Reference](https://developer.mozilla.org/en-US/docs/Web/HTML/Element/thead)

---

### Time 

`Time(items ...Item) *Element`

Creates an HTML `<time>` element representing a specific time or date.

[MDN Reference](https://developer.mozilla.org/en-US/docs/Web/HTML/Element/time)

---

### Title 

`Title(items ...Item) *Element`

Creates an HTML `<title>` element defining the document’s title.

[MDN Reference](https://developer.mozilla.org/en-US/docs/Web/HTML/Element/title)

---

### Tr 

`Tr(items ...Item) *Element`

Creates an HTML `<tr>` element defining a row in a table.

[MDN Reference](https://developer.mozilla.org/en-US/docs/Web/HTML/Element/tr)

---

### U 

`U(items ...Item) *Element`

Creates an HTML `<u>` element representing underlined or annotated text.

[MDN Reference](https://developer.mozilla.org/en-US/docs/Web/HTML/Element/u)

---

### Ul 

`Ul(items ...Item) *Element`

Creates an HTML `<ul>` element representing an unordered (bulleted) list.

[MDN Reference](https://developer.mozilla.org/en-US/docs/Web/HTML/Element/ul)

---

### Var 

`Var(items ...Item) *Element`

Creates an HTML `<var>` element representing a variable name in a mathematical or programming context.

[MDN Reference](https://developer.mozilla.org/en-US/docs/Web/HTML/Element/var)

---

### Video 

`Video(items ...Item) *Element`

Creates an HTML `<video>` element embedding a media player for video content.

[MDN Reference](https://developer.mozilla.org/en-US/docs/Web/HTML/Element/video)

---

### Wbr 

`Wbr(items ...Item) *Element`

Creates an HTML `<wbr>` element defining a word break opportunity.

[MDN Reference](https://developer.mozilla.org/en-US/docs/Web/HTML/Element/wbr)
