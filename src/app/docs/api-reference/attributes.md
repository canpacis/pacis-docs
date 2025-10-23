# Attributes

## Accept

{noaccessory=true}
```go
func Accept(value string) *Attribute
```

The accept attribute takes as its value a comma-separated list of one or more file types, or unique file type specifiers, describing which file types to allow.

[MDN Reference](https://developer.mozilla.org/en-US/docs/Web/HTML/Reference/Attributes/accept)

---

## AcceptCharset

{noaccessory=true}
```go
func AcceptCharset(value string) *Attribute
```

Specifies the character encodings that are to be used for the form submission.

---

## AccessKey

{noaccessory=true}
```go
func AccessKey(value string) *Attribute
```

Provides a hint for generating a keyboard shortcut for the current element.

---

## Action

{noaccessory=true}
```go
func Action(value string) *Attribute
```

Specifies where to send the form-data when a form is submitted.

---

## Align

{noaccessory=true}
```go
func Align(value string) *Attribute
```

Specifies the alignment of an element.

---

## Alt

{noaccessory=true}
```go
func Alt(value string) *Attribute
```

Defines text that can replace the image in the page.

Browsers do not always display images. There are a number of situations in which a browser might not display images, such as:

- Non-visual browsers (such as those used by people with visual impairments)
- The user chooses not to display images (saving bandwidth, privacy reasons)
- The image is invalid or an unsupported type

In these cases, the browser may replace the image with the text in the element's alt attribute. For these reasons and others, provide a useful value for alt whenever possible.

Setting this attribute to an empty string (alt="") indicates that this image is not a key part of the content (it's decoration or a tracking pixel), and that non-visual browsers may omit it from rendering. Visual browsers will also hide the broken image icon if the alt attribute is empty and the image failed to display.

This attribute is also used when copying and pasting the image to text, or saving a linked image to a bookmark.

[MDN Reference](https://developer.mozilla.org/en-US/docs/Web/HTML/Reference/Elements/img#alt)

---

## Aria

{noaccessory=true}
```go
func Aria(name, value string) *Attribute
```

Creates an ARIA attribute by prefixing the name with "aria-".

---

## As

{noaccessory=true}
```go
func As(value string) *Attribute
```

Specifies the type of content being loaded by a `<link>` element.

---

## Async

{noaccessory=true}
```go
func Async(value string) *Attribute
```

Indicates that the script should be executed asynchronously.

---

## Autocomplete

{noaccessory=true}
```go
func Autocomplete(value string) *Attribute
```

It is available on `<input>` elements that take a text or numeric value as input, `<textarea>` elements, `<select>` elements, and `<form>` elements.

[MDN Reference](https://developer.mozilla.org/en-US/docs/Web/HTML/Reference/Attributes/autocomplete)

---

## Autofocus

{noaccessory=true}
```go
func Autofocus(value string) *Attribute
```

Indicates that an element should automatically get focus when the page loads.

---

## Autoplay

{noaccessory=true}
```go
func Autoplay(value string) *Attribute
```

Indicates that the audio or video should start playing as soon as it is ready.

---

## Charset

{noaccessory=true}
```go
func Charset(value string) *Attribute
```

This attribute declares the document's character encoding. If the attribute is present, its value must be an ASCII case-insensitive match for the string "utf-8", because UTF-8 is the only valid encoding for HTML5 documents. `<meta>` elements which declare a character encoding must be located entirely within the first 1024 bytes of the document.

[MDN Reference](https://developer.mozilla.org/en-US/docs/Web/HTML/Reference/Elements/meta#charset)

---

## Checked

{noaccessory=true}
```go
var Checked = Attr("checked", "")
```

A boolean attribute indicating that a checkbox or radio button is selected.

---

## CiteAttr

{noaccessory=true}
```go
func CiteAttr(value string) *Attribute
```

A URL that designates a source document or message for the information quoted. This attribute is intended to point to information explaining the context or the reference for the quote.

[MDN Reference](https://developer.mozilla.org/en-US/docs/Web/HTML/Reference/Elements/blockquote#cite)

---

## Class

{noaccessory=true}
```go
func Class(class string) *Attribute
```

The class global attribute is a list of the classes of the element, separated by ASCII whitespace.

[MDN Reference](https://developer.mozilla.org/en-US/docs/Web/HTML/Reference/Global_attributes/class)

---

## ColorAttr

{noaccessory=true}
```go
func ColorAttr(value string) *Attribute
```

Specifies a color value.

---

## Cols

{noaccessory=true}
```go
func Cols(value string) *Attribute
```

Specifies the visible width of a text area.

---

## Colspan

{noaccessory=true}
```go
func Colspan(value string) *Attribute
```

Specifies the number of columns a table cell should span.

---

## Content

{noaccessory=true}
```go
func Content(value string) *Attribute
```

The content attribute specifies the value of a metadata name defined by the `<meta>` name attribute. It takes a string as its value, and the expected syntax varies depending on the name value used.

[MDN Reference](https://developer.mozilla.org/en-US/docs/Web/HTML/Reference/Attributes/content)

---

## ContentEditable

{noaccessory=true}
```go
var ContentEditable = Attr("contenteditable", "")
```

The contenteditable global attribute is an enumerated attribute indicating if the element should be editable by the user. If so, the browser modifies its widget to allow editing.

[MDN Reference](https://developer.mozilla.org/en-US/docs/Web/HTML/Reference/Global_attributes/contenteditable)

---

## Controls

{noaccessory=true}
```go
func Controls(value string) *Attribute
```

Indicates whether the browser should show playback controls for media elements.

---

## Coords

{noaccessory=true}
```go
func Coords(value string) *Attribute
```

Specifies the coordinates of an area in an image map.

---

## Crossorigin

{noaccessory=true}
```go
func Crossorigin(value string) *Attribute
```

The crossorigin attribute, valid on the `<audio>`, `<img>`, `<link>`, `<script>`, and `<video>` elements, provides support for CORS, defining how the element handles cross-origin requests, thereby enabling the configuration of the CORS requests for the element's fetched data. Depending on the element, the attribute can be a CORS settings attribute.

The crossorigin content attribute on media elements is a CORS settings attribute.

These attributes are enumerated, and have the following possible values:

Request uses CORS headers and credentials flag is set to 'same-origin'. There is no exchange of user credentials via cookies, client-side TLS certificates or HTTP authentication, unless destination is the same origin.

Request uses CORS headers, credentials flag is set to 'include' and user credentials are always included.

Setting the attribute name to an empty value, like crossorigin or crossorigin="", is the same as anonymous.

An invalid keyword and an empty string will be handled as the anonymous keyword.

By default (that is, when the attribute is not specified), CORS is not used at all. The user agent will not ask for permission for full access to the resource and in the case of a cross-origin request, certain limitations will be applied based on the type of element concerned.

> The crossorigin attribute is not supported for rel="icon" in Chromium-based browsers. See the open Chromium issue.

[MDN Reference](https://developer.mozilla.org/en-US/docs/Web/HTML/Reference/Attributes/crossorigin)

---

## Data

{noaccessory=true}
```go
func Data(name, value string) *Attribute
```

The data-* global attributes form a class of attributes called custom data attributes, that allow proprietary information to be exchanged between the HTML and its DOM representation by scripts.

[MDN Reference](https://developer.mozilla.org/en-US/docs/Web/HTML/Reference/Global_attributes/data-*)

---

## Datetime

{noaccessory=true}
```go
func Datetime(value string) *Attribute
```

Specifies the date and time of an element.

---

## DefaultAttr

{noaccessory=true}
```go
func DefaultAttr(value string) *Attribute
```

Specifies that a track should be enabled by default.

---

## Defer

{noaccessory=true}
```go
var Defer = Attr("defer", "")
```

Indicates that the script should be executed after the document has been parsed.

---

## Dirname

{noaccessory=true}
```go
func Dirname(value string) *Attribute
```

The dirname attribute can be used on the `<textarea>` element and several `<input>` types and describes the directionality of the element's text content during form submission. The browser uses this attribute's value to determine whether text the user has entered is left-to-right or right-to-left oriented. When used, the element's text directionality value is included in form submission data along with the dirname attribute's value as the name of the field.

[MDN Reference](https://developer.mozilla.org/en-US/docs/Web/HTML/Reference/Attributes/dirname)

---

## Disabled

{noaccessory=true}
```go
var Disabled = Attr("disabled", "")
```

The Boolean disabled attribute, when present, makes the element not mutable, focusable, or even submitted with the form. The user can neither edit nor focus on the control, nor its form control descendants.

[MDN Reference](https://developer.mozilla.org/en-US/docs/Web/HTML/Reference/Attributes/disabled)

---

## Download

{noaccessory=true}
```go
func Download(value string) *Attribute
```

Indicates that the hyperlink is to be used for downloading a resource.

---

## Draggable

{noaccessory=true}
```go
func Draggable(value string) *Attribute
```

The draggable global attribute is an enumerated attribute that indicates whether the element can be dragged, either with native browser behavior or the HTML Drag and Drop API.

The draggable attribute may be applied to elements that strictly fall under the HTML namespace, which means that it cannot be applied to SVGs. For more information about what namespace declarations look like, and what they do, see Namespace crash course.

draggable can have the following values:

> This attribute is enumerated and not Boolean. A value of true or false is mandatory, and shorthand like `<img draggable>` is forbidden. The correct usage is `<img draggable="true">`.

If this attribute is not set, its default value is auto, which means drag behavior is the default browser behavior: only text selections, images, and links can be dragged. For other elements, the event ondragstart must be set for drag and drop to work, as shown in this comprehensive example.

[MDN Reference](https://developer.mozilla.org/en-US/docs/Web/HTML/Reference/Global_attributes/draggable)

---

## Enctype

{noaccessory=true}
```go
func Enctype(value string) *Attribute
```

Specifies how form data should be encoded when submitting to the server.

---

## EnterKeyHint

{noaccessory=true}
```go
func EnterKeyHint(value string) *Attribute
```

Hints at the type of action label for the Enter key on virtual keyboards.

---

## For

{noaccessory=true}
```go
func For(value string) *Attribute
```

The for attribute is an allowed attribute for `<label>` and `<output>`. When used on a `<label>` element it indicates the form element that this label describes. When used on an `<output>` element it allows for an explicit relationship between the elements that represent values which are used in the output.

[MDN Reference](https://developer.mozilla.org/en-US/docs/Web/HTML/Reference/Attributes/for)

---

## FormAttr

{noaccessory=true}
```go
func FormAttr(value string) *Attribute
```

The form HTML attribute associates a form-associated element with a `<form>` element within the same document. This attribute applies to the `<button>`, `<fieldset>`, `<input>`, `<object>`, `<output>`, `<select>`, and `<textarea>` elements.

[MDN Reference](https://developer.mozilla.org/en-US/docs/Web/HTML/Reference/Attributes/form)

---

## FormAction

{noaccessory=true}
```go
func FormAction(value string) *Attribute
```

The URL that processes the form submission. This value can be overridden by a formaction attribute on a `<button>`, `<input type="submit">`, or `<input type="image">` element. This attribute is ignored when method="dialog" is set.

[MDN Reference](https://developer.mozilla.org/en-US/docs/Web/HTML/Reference/Elements/form#action)

---

## Headers

{noaccessory=true}
```go
func Headers(value string) *Attribute
```

Specifies one or more header cells a cell is related to.

---

## Height

{noaccessory=true}
```go
func Height(value string) *Attribute
```

The height attribute defines the vertical length of an element in the user coordinate system.

[MDN Reference](https://developer.mozilla.org/en-US/docs/Web/SVG/Reference/Attribute/height)

---

## Hidden

{noaccessory=true}
```go
func Hidden(value string) *Attribute
```

The hidden global attribute is an enumerated attribute indicating that the browser should not render the contents of the element. For example, it can be used to hide elements of the page that can't be used until the login process has been completed.

[MDN Reference](https://developer.mozilla.org/en-US/docs/Web/HTML/Reference/Global_attributes/hidden)

---

## High

{noaccessory=true}
```go
func High(value string) *Attribute
```

Specifies the range that is considered to be a high value for a meter element.

---

## Href

{noaccessory=true}
```go
func Href(value string) *Attribute
```

The URL that the hyperlink points to. Links are not restricted to HTTP-based URLs — they can use any URL scheme supported by browsers:

- Telephone numbers with tel: URLs
- Email addresses with mailto: URLs
- SMS text messages with sms: URLs
- Executable code with javascript: URLs
- While web browsers may not support other URL schemes, websites can with registerProtocolHandler()

Moreover other URL features can locate specific parts of the resource, including:

- Sections of a page with document fragments
- Specific text portions with text fragments
- Pieces of media files with media fragments

[MDN Reference](https://developer.mozilla.org/en-US/docs/Web/HTML/Reference/Elements/a#href)

---

## HrefLang

{noaccessory=true}
```go
func HrefLang(value string) *Attribute
```

Hints at the human language of the linked URL. No built-in functionality. Allowed values are the same as the global lang attribute.

[MDN Reference](https://developer.mozilla.org/en-US/docs/Web/HTML/Reference/Elements/a#hreflang)

---

## HttpEquiv

{noaccessory=true}
```go
func HttpEquiv(value string) *Attribute
```

The http-equiv attribute of the `<meta>` element allows you to provide processing instructions for the browser as if the response that returned the document included certain HTTP headers. The metadata is document-level metadata that applies to the whole page.

When a `<meta>` element has an http-equiv attribute, a content attribute defines the corresponding http-equiv value. For example, the following `<meta>` tag tells the browser to refresh the page after 5 minutes.

[MDN Reference](https://developer.mozilla.org/en-US/docs/Web/HTML/Reference/Elements/meta/http-equiv)

---

## ID

{noaccessory=true}
```go
func ID(value string) *Attribute
```

The id global attribute defines an identifier (ID) that must be unique within the entire document.

[MDN Reference](https://developer.mozilla.org/en-US/docs/Web/HTML/Reference/Global_attributes/id)

---

## Inert

{noaccessory=true}
```go
func Inert(value string) *Attribute
```

Indicates that the element and its descendants are not interactive.

---

## InputMode

{noaccessory=true}
```go
func InputMode(value string) *Attribute
```

Provides a hint to browsers about the type of virtual keyboard to display.

---

## IsMap

{noaccessory=true}
```go
func IsMap(value string) *Attribute
```

Indicates that the image is part of a server-side image map.

---

## Kind

{noaccessory=true}
```go
func Kind(value string) *Attribute
```

Specifies the kind of text track for media elements.

---

## LabelAttr

{noaccessory=true}
```go
func LabelAttr(value string) *Attribute
```

Specifies a user-readable title for an element.

---

## Lang

{noaccessory=true}
```go
func Lang(value string) *Attribute
```

The lang global attribute helps define the language of an element: the language that non-editable elements are written in, or the language that the editable elements should be written in by the user. The attribute contains a single BCP 47 language tag.

> The default value of lang is the empty string, which means that the language is unknown. Therefore, it is recommended to always specify an appropriate value for this attribute.

[MDN Reference](https://developer.mozilla.org/en-US/docs/Web/HTML/Reference/Global_attributes/lang)

---

## List

{noaccessory=true}
```go
func List(value string) *Attribute
```

Identifies a list of pre-defined options to suggest to the user.

---

## Loop

{noaccessory=true}
```go
func Loop(value string) *Attribute
```

Indicates that the media should start playing again when it reaches the end.

---

## Low

{noaccessory=true}
```go
func Low(value string) *Attribute
```

Specifies the range that is considered to be a low value for a meter element.

---

## Max

{noaccessory=true}
```go
func Max(value string) *Attribute
```

The max attribute defines the maximum value that is acceptable and valid for the input containing the attribute. If the value of the element is greater than this, the element fails validation. This value must be greater than or equal to the value of the min attribute. If the max attribute is present but is not specified or is invalid, no max value is applied. If the max attribute is valid and a non-empty value is greater than the maximum allowed by the max attribute, constraint validation will prevent form submission.

The max attribute is valid for the numeric input types, including the date, month, week, time, datetime-local, number and range types, and both the `<progress>` and `<meter>` elements. It is a number that specifies the most positive value a form control to be considered valid.

If the value exceeds the max value allowed, the validityState.rangeOverflow will be true, and the control will be matched by the :out-of-range and :invalid pseudo-classes.

[MDN Reference](https://developer.mozilla.org/en-US/docs/Web/HTML/Reference/Attributes/max)

---

## MaxLength

{noaccessory=true}
```go
func MaxLength(value string) *Attribute
```

The maxlength attribute defines the maximum string length that the user can enter into an `<input>` or `<textarea>`. The attribute must have an integer value of 0 or higher.

The length is measured in UTF-16 code units, which is often but not always equal to the number of characters. If no maxlength is specified, or an invalid value is specified, the input has no maximum length.

Any maxlength value must be greater than or equal to the value of minlength, if present and valid. The input will fail constraint validation if the length of the text value of the field is greater than maxlength UTF-16 code units long. Constraint validation is only applied when the value is changed by the user.

[MDN Reference](https://developer.mozilla.org/en-US/docs/Web/HTML/Reference/Attributes/maxlength)

---

## Media

{noaccessory=true}
```go
func Media(value string) *Attribute
```

The media attribute defines which media the theme color defined in the content attribute should be applied to. Its value is a media query, which defaults to all if the attribute is missing. This attribute is only relevant when the element's name attribute is set to theme-color. Otherwise, it has no effect, and should not be included.

[MDN Reference](https://developer.mozilla.org/en-US/docs/Web/HTML/Reference/Elements/meta#media)

---

## Method

{noaccessory=true}
```go
func Method(value string) *Attribute
```

The HTTP method to submit the form with. The only allowed methods/values are (case insensitive):

- post: The POST method; form data sent as the request body.
- get (default): The GET; form data appended to the action URL with a ? separator. Use this method when the form has no side effects.
- dialog: When the form is inside a `<dialog>`, closes the dialog and causes a submit event to be fired on submission, without submitting data or clearing the form.

This value is overridden by formmethod attributes on `<button>`, `<input type="submit">`, or `<input type="image">` elements.

[MDN Reference](https://developer.mozilla.org/en-US/docs/Web/HTML/Reference/Elements/form#method)

---

## Min

{noaccessory=true}
```go
func Min(value string) *Attribute
```

The min attribute defines the minimum value that is acceptable and valid for the input containing the attribute. If the value of the element is less than this, the element fails validation. This value must be less than or equal to the value of the max attribute.

Some input types have a default minimum. If the input has no default minimum and a value is specified for min that can't be converted to a valid number (or no minimum value is set), the input has no minimum value.

It is valid for the input types including: date, month, week, time, datetime-local, number and range types, and the `<meter>` element.

[MDN Reference](https://developer.mozilla.org/en-US/docs/Web/HTML/Reference/Attributes/min)

---

## MinLength

{noaccessory=true}
```go
func MinLength(value string) *Attribute
```

The minlength attribute defines the minimum string length that the user can enter into an `<input>` or `<textarea>`. The attribute must have an integer value of 0 or higher.

The length is measured in UTF-16 code units, which is often but not always equal to the number of characters. If no minlength is specified, or an invalid value is specified, the input has no minimum length. This value must be less than or equal to the value of maxlength, otherwise the value will never be valid, as it is impossible to meet both criteria.

The input will fail constraint validation if the length of the text value of the field is less than minlength UTF-16 code units long, with validityState.tooShort returning true. Constraint validation is only applied when the value is changed by the user. Once submission fails, some browsers will display an error message indicating the minimum length required and the current length.

minlength does not imply required; an input only violates a minlength constraint if the user has input a value. If an input is not required, an empty string can be submitted even if minlength is set.

[MDN Reference](https://developer.mozilla.org/en-US/docs/Web/HTML/Reference/Attributes/minlength)

---

## Multiple

{noaccessory=true}
```go
func Multiple(value string) *Attribute
```

The Boolean multiple attribute, if set, means the form control accepts one or more values. The attribute is valid for the email and file input types and the `<select>`. The manner by which the user opts for multiple values depends on the form control.

[MDN Reference](https://developer.mozilla.org/en-US/docs/Web/HTML/Reference/Attributes/multiple)

---

## Muted

{noaccessory=true}
```go
func Muted(value string) *Attribute
```

Indicates whether the audio output should be muted by default.

---

## Name

{noaccessory=true}
```go
func Name(value string) *Attribute
```

The name and content attributes can be used together to provide document metadata in terms of name-value pairs, with the name attribute giving the metadata name, and the content attribute giving the value.

[MDN Reference](https://developer.mozilla.org/en-US/docs/Web/HTML/Reference/Elements/meta#name)

---

## NoValidate

{noaccessory=true}
```go
var NoValidate = Attr("novalidate", "")
```

This Boolean attribute indicates that the form shouldn't be validated when submitted. If this attribute is not set (and therefore the form is validated), it can be overridden by a formnovalidate attribute on a `<button>`, `<input type="submit">`, or `<input type="image">` element belonging to the form.

[MDN Reference](https://developer.mozilla.org/en-US/docs/Web/HTML/Reference/Elements/form#novalidate)

---

## Placeholder

{noaccessory=true}
```go
func Placeholder(value string) *Attribute
```

Effective placeholder text includes a word or short phrase that hints at the expected data type, not an explanation or prompt. The placeholder must not be used instead of a `<label>`. As the placeholder is not visible if the value of the form control is not null, using placeholder instead of a `<label>` for a prompt harms usability and accessibility.

The placeholder attribute is supported by the following input types: text, search, url, tel, email, and password. It is also supported by the `<textarea>` element. The example below shows the placeholder attribute in use to explain the expected format of an input field.

> Except in `<textarea>` elements, the placeholder attribute can't include any line feeds (LF) or carriage returns (CR). If either is included in the value, the placeholder text will be clipped.

[MDN Reference](https://developer.mozilla.org/en-US/docs/Web/HTML/Reference/Attributes/placeholder)

---

## PropertyAttr

{noaccessory=true}
```go
func PropertyAttr(value string) *Attribute
```

Defines a property for metadata purposes.

---

## Rel

{noaccessory=true}
```go
func Rel(value string) *Attribute
```

The rel attribute defines the relationship between a linked resource and the current document. Valid on `<link>`, `<a>`, `<area>`, and `<form>`, the supported values depend on the element on which the attribute is found.

The type of relationships is given by the value of the rel attribute, which, if present, must have a value that is an unordered set of unique space-separated keywords. Differently from a class name, which does not express semantics, the rel attribute must express tokens that are semantically valid for both machines and humans. The current registries for the possible values of the rel attribute are the IANA link relation registry, the HTML Living Standard, and the freely-editable existing-rel-values page in the microformats wiki, as suggested by the Living Standard. If a rel attribute not present in one of the three sources above is used some HTML validators (such as the W3C Markup Validation Service) will generate a warning.

The following table lists some of the most important existing keywords. Every keyword within a space-separated value should be unique within that value.

The rel attribute is relevant to the `<link>`, `<a>`, `<area>`, and `<form>` elements, but some values only relevant to a subset of those elements. Like all HTML keyword attribute values, these values are case-insensitive.

The rel attribute has no default value. If the attribute is omitted or if none of the values in the attribute are supported, then the document has no particular relationship with the destination resource other than there being a hyperlink between the two. In this case, on `<link>` and `<form>`, if the rel attribute is absent, has no keywords, or if not one or more of the space-separated keywords above, then the element does not create any links. `<a>` and `<area>` will still created links, but without a defined relationship.

[MDN Reference](https://developer.mozilla.org/en-US/docs/Web/HTML/Reference/Attributes/rel)

---

## Required

{noaccessory=true}
```go
var Required = Attr("required", "")
```

The Boolean required attribute, if present, indicates that the user must specify a value for the input before the owning form can be submitted.

The required attribute is supported by text, search, url, tel, email, password, date, month, week, time, datetime-local, number, checkbox, radio, file, `<input>` types along with the `<select>` and `<textarea>` form control elements. If present on any of these input types and elements, the :required pseudo class will match. If the attribute is not included, the :optional pseudo class will match.

The attribute is not supported on, or relevant to, range and color input types, as both have default values. Type color defaults to #000000. Type range defaults to the midpoint between min and max — with min and max defaulting to 0 and 100 respectively in most browsers if not declared. required is also not supported on the hidden input type — users cannot be expected to fill out a hidden form field. Finally, required is not supported on any button input types, including image.

In the case of a same named group of radio buttons, if a single radio button in the group has the required attribute, a radio button in that group must be checked, although it doesn't have to be the one on which the attribute is applied. To improve code maintenance, it is recommended to either include the required attribute in every same-named radio button in the group, or else in none.

In the case of a same named group of checkbox input types, only the checkboxes with the required attribute are required.

> Setting aria-required="true" tells a screen reader that an element (any element) is required, but has no bearing on the optionality of the element.

[MDN Reference](https://developer.mozilla.org/en-US/docs/Web/HTML/Reference/Attributes/required)

---

## Role

{noaccessory=true}
```go
func Role(value string) *Attribute
```

Defines the role of an element in the accessibility tree.

---

## ShadowRootMode

{noaccessory=true}
```go
func ShadowRootMode(value string) *Attribute
```

Creates a shadow root for the parent element. It is a declarative version of the Element.attachShadow() method and accepts the same enumerated values.

[MDN Reference](https://developer.mozilla.org/en-US/docs/Web/HTML/Reference/Elements/template#shadowrootmode)

---

## SlotAttr

{noaccessory=true}
```go
func SlotAttr(value string) *Attribute
```

The slot global attribute assigns a slot in a shadow DOM shadow tree to an element: An element with a slot attribute is assigned to the slot created by the `<slot>` element whose name attribute's value matches that slot attribute's value. You can have multiple elements assigned to the same slot by using the same slot name. Elements without a slot attribute are assigned to the unnamed slot, if one exists.

For examples, see our Using templates and slots guide.

[MDN Reference](https://developer.mozilla.org/en-US/docs/Web/HTML/Reference/Global_attributes/slot)

---

## Src

{noaccessory=true}
```go
func Src(value string) *Attribute
```

Specifies the URL of an external resource.

---

## StyleAttr

{noaccessory=true}
```go
func StyleAttr(value string) *Attribute
```

The style global attribute contains CSS styling declarations to be applied to the element. Note that it is recommended for styles to be defined in a separate file or files. This attribute and the `<style>` element have mainly the purpose of allowing for quick styling, for example for testing purposes.

[MDN Reference](https://developer.mozilla.org/en-US/docs/Web/HTML/Reference/Global_attributes/style)

---

## Tabindex

{noaccessory=true}
```go
func Tabindex(value string) *Attribute
```

The tabindex global attribute allows developers to make HTML elements focusable, allow or prevent them from being sequentially focusable (usually with the Tab key, hence the name) and determine their relative ordering for sequential focus navigation.

[MDN Reference](https://developer.mozilla.org/en-US/docs/Web/HTML/Reference/Global_attributes/tabindex)

---

## Target

{noaccessory=true}
```go
func Target(value string) *Attribute
```

Where to display the linked URL, as the name for a browsing context (a tab, window, or `<iframe>`). The following keywords have special meanings for where to load the URL:

- _self: The current browsing context. (Default)
- _blank: Usually a new tab, but users can configure browsers to open a new window instead.
- _parent: The parent browsing context of the current one. If no parent, behaves as _self.
- _top: The topmost browsing context. To be specific, this means the "highest" context that's an ancestor of the current one. If no ancestors, behaves as _self.
- _unfencedTop: Allows embedded fenced frames to navigate the top-level frame (i.e., traversing beyond the root of the fenced frame, unlike other reserved destinations). Note that the navigation will still succeed if this is used outside of a fenced frame context, but it will not act like a reserved keyword.

[MDN Reference](https://developer.mozilla.org/en-US/docs/Web/HTML/Reference/Elements/a#target)

---

## TitleAttr

{noaccessory=true}
```go
func TitleAttr(value string) *Attribute
```

The title global attribute contains text representing advisory information related to the element it belongs to.

[MDN Reference](https://developer.mozilla.org/en-US/docs/Web/HTML/Reference/Global_attributes/title)

---

## Type

{noaccessory=true}
```go
func Type(value string) *Attribute
```

A string specifying the type of control to render. For example, to create a checkbox, a value of checkbox is used. If omitted (or an unknown value is specified), the input type text is used, creating a plaintext input field.

[MDN Reference](https://developer.mozilla.org/en-US/docs/Web/HTML/Reference/Elements/input#type)

---

## Value

{noaccessory=true}
```go
func Value(value string) *Attribute
```

The input control's value. When specified in the HTML, this is the initial value, and from then on it can be altered or retrieved at any time using JavaScript to access the respective HTMLInputElement object's value property. The value attribute is always optional, though should be considered mandatory for checkbox, radio, and hidden.

[MDN Reference](https://developer.mozilla.org/en-US/docs/Web/HTML/Reference/Elements/input#value)

---

## Width

{noaccessory=true}
```go
func Width(value string) *Attribute
```

The width attribute defines the horizontal length of an element in the user coordinate system.

[MDN Reference](https://developer.mozilla.org/en-US/docs/Web/SVG/Reference/Attribute/width)