# Views: Component System
In Previous, HTML is dynamically generated within the web server process upon an HTTP request.
Rendering HTML on the server results in simpler code, and a smaller codebase.
Despite what some may claim, server-generated HTML does _not_ compromise on interactivity, or "user experience".

Advantages to server-rendered HTML over client-rendered:

- [Locality of behavior](https://htmx.org/essays/locality-of-behaviour/).
- Simpler system design.
- No duplicate validation, authentication, and data modeling on both the client and server.

## Components

To avoid repeating template logic and view content, Previous utilizes a component system to construct HTML.
There is no "template language/syntax" like _Mustache_, or _Razor_ to write HTML.
In fact, by recognizing that HTML is a tree data structure, you don't even need to write HTML markup at all!

A "component" is simply a function, that returns a `gomponents.Node`.
A `gomponents.Node` is an interface, that can render itself to a `writer` as HTML.
This means you can include any valid Go code inside a component.
This _also_ means that components are debuggable!
When a component is rendered, all subsequent components down the tree are called, and written to the `writer`.

Previous uses [Gomponents](https://www.gomponents.com/) for component building.
Please refer to their documentation for detailed documentation.

## Component Construction
In order to use components, you must import the `gomponents` package.
It is house-style to use a dot import (module contents injected directly into scope) for components.

```go
import (
	. "maragu.dev/gomponents"      // gives us primative components such as `Text`
	. "maragu.dev/gomponents/html" // gives us HTML components like `H1` and `P`
)
```

Components at their simplest form, take this form:

```go
func ExampleComponent() Node {
	return P(Class("fw-bold"),
		Text("Hello, World!"),
	)
}
```

Components can be called inside other components:

```go
func BiggerComponent() Node {
	return Div(Class("some-class"),
		ExampleComponent(), // this is defined by us earlier (see above)
		Span(Text("More content...")),
	)
}

```

Components can also be composed by passing (spreading) an array of `Node` as an argument to another component:
```go
func MyLayoutComponent(children ...Node) Node {
	return Main(
		Nav(Text("Hello this is the application layout")),
		Group(children),
	)
}

// usage of layout
func MyPageComponent() Node {
	return MyLayoutComponent(
		P(Text("Page content!")),
	)
}
```

This is useful for creating "layouts"/"templates"/"master pages", and also reusable component building blocks.



## Building A Complete Page

In this codebase, all pages stem from the `Root` component, which includes

```go
func BasicPageView(children ...Node) Node {
	return Doctype(
		HTML(
			Lang("en"),
			Head(
				Meta(Charset("utf-8")),
				Meta(Name("viewport"), Content("width=device-width, initial-scale=1")),
				Meta(Name("description"), Content("Previous")),
			),
			Body(
				H1(Text("Hello, World!")),
				Group(children),
			),
		),
	)
}
```
## Special Builtins

### Text Component
Text is inserted into a component by using the `Text` component.
It takes a string as an argument.
During the rendering process, it will place the input text string inside HTML tags.

You can also use the `ToText` component to render non-string arguments as HTML text.
Under the hood, this is a wrapper around `printf`.

```go
func ParagraphComponent(input string, numberInput int) Node {
	return P(Text(input), ToText(numberInput))
}

// ParagraphComponent("hello", 5)
// => <p>hello5</p>
```
Note that the `Text` component will _automatically escape HTML_!
This component is safe to use for user generated content without worrying about XSS!

### Attribute Component
Attributes are called like child components, although they are not rendered as children, but rather as HTML attributes.
This is done using the `Attr` component:

```go
func AttributeExampleComponent() Node {
	return Div(Attr("some-data", "some-value"))
}

// => <div some-data="some-value"></div>
```

By passing one argument to `Attr` instead of two, it will render to HTML without a value attached.

```go
func SingleAttributeExampleComponent() Node {
	return Div(Attr("single-attribute"))
}

// => <div single-attribute></div>
```
### Class / Classes Component
Classes can be added to components by calling the `Class` component with a string containing the list of CSS classes.

```go
func HeadingComponent() Node {
	return H1(Class("text-center text-primary"), Text("Hello, World!"))
}

// => <h1 class="text-center text-primary">Hello, World!</h1>
```

You can alternatively use the `Classes` struct to conditionally include CSS classes onto a component:

```go
func HeadingComponent(cond bool) Node {
	return H1(Classes{"text-center fw-bold": true, "text-success": cond}, Text("Hello, World!"))
}

// HeadingComponent(false)
// => <h1 class="text-center fw-bold">Hello, World!</h1>
//
// HeadingComponent(true)
// => <h1 class="text-center fw-bold text-success">Hello, World!</h1>
```

### Map Component (looping)
There is a builtin component designed for handling data iteration.
Large components will have many subcomponents, and it is simply more ergonomic to use the `Map` helper, rather than a normal `for` loop.
The map component takes in an array of any type, and a "callback" component.
For each element in the array, the component function is called with the element of the array being passed to it as an argument.


```go
func MapListComponent(people []string) Node {
	return Ul(
		Map(people, func(person string) Node {
			return Li(Text(person)),
		})
	),
}
```

### Conditional Components
There are builtin components to handle conditional logic: `If`, `IfElse`, `Iff`, and `IffElse`.
Just like the `Map` component, it is purely to improve the ergonomics of the component system.

```go
// Signatures:
func If(condition bool, t Node) Node
func IfElse(condition bool, t Node, f Node) Node
func Iff(condition bool, t func() Node) Node
func IffElse(condition bool, t func() Node, f func() Node) Node

// Usage:
func MyComponent(cond bool) Node {
	return P(
		Text("This text will always print. "),
		If(cond, Text("This text will only show if input is true!")),
	)
}
```

There is a special version of the `If` function (named `Iff`) that will execute a callback function if the condition is true.

We also include `IfElse` and `IffElse` components for handling false cases.

### The "Raw" Component
If you want to insert plain HTML into a view, you can use the `Raw` component.
Note that this component _bypasses ALL HTML validation_!
Used improperly, this is a one way ticket to XSS!

```go
func SomeComponent() Node {
	return Raw(`
		<p>Hello world</p>
		<script>
			console.log("Pwned!");
		</script>
	`)
}
```

## Rendering
As discussed in previous chapters on [routing](/docs/routing), and [controllers](/docs/controllers), views, components must be "rendered" from a page controller.
The entire structure of your component tree is constructed in memory first, then transformed into HTML.
This happens when you call `YourComponent().Render(w)`.

```go
type PageData struct {
	someData int
}

func MyWebpageView(data PageData) Node {
	return MyWebpageLayout(
		P(Text("Welcome to my webpage!")),
		P(ToText(data.someData)),
	)
}
func MyWebpageController(w http.ResponseWriter, r *http.Response) {
	// handle page logic here
	data := getPageData() // returns a PageData
	...

	// render the component tree to the response writer
	MyWebpageView(data).Render(w)
}
```

It is important to remember that calling the `Render` function is what begins all subsequent calls down the component tree.
Components are just regular functions that contain any arbitrary code!
When you call the `Render` function on a `Node` tree, _all_ functions down the tree are called.
As expected, if you call a component function recursively without bound, the stack memory will overflow.

## Hygiene
Because components are arbitrary functions with the capability of performing complex logic, you must take great care when structuring your page views.
Although it is _possible_, it is not recommended to, for example, call out to the database directly from a component.
Instead, you should pass the result from a database call _to_ a component.

With large `Node` trees containing many subcomponents, potentially spanning across different files and packages, it can become difficult to keep track of who is calling what.
To help keep things organized, put HTTP logic inside the `Controller`, "business logic" inside a delegated package (Ex: database related stuff), and only leave "display logic" up to the component.