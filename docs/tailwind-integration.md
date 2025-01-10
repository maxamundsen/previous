# Views: Tailwind Integration

## Usage & Compiling
The Previous codebase includes a copy of the Tailwind compiler alongside the build script, where it is called.

```sh
sh build_server.sh
```

This automatically builds the Tailwind stylesheet from `/src/pages/**/*.go` source files, and outputs a stylesheet to `/src/wwwroot/css/style.css`.

## Why Tailwind?
TailwindCSS reduces the boilerplate, and ceremony required when authoring CSS, when using a component-based design.

Traditionally, HTML elements can be arbitrarily grouped and targeted for styling via CSS classes.
Consider the following example for a simple dialog box:

```css
.dialog-box {
	font-weight: bold;
	color: red;
	...
}

.dialog-header {
	font-size: 18px;
	color: green;
}
```

```go
func DialogBox(header string, children ...Node) Node {
	return Div(Class("dialog-box"),
		Div(Class("dialog-header"), Text(header)),
		Group(children),
	)
}
```

Using CSS classes this way _duplicates_ the amount of work required, since you are _already_ grouping elements together with the component system!
TailwindCSS essentially allows styles to be coupled with the component, while also providing a uniform design system to make authoring CSS easier, and maintainable down the road.

By using TailwindCSS to inline styles, the above example turns into this:

```go
func DialogBox(header string, children ...Node) Node {
	return Div(Class("font-bold text-red-500"),
		Div(Class("text-lg text-green-500"), Text(header)),
		Group(children),
	)
}
```

While this may seem awkward and perhaps verbose, it avoids arbitrary CSS classes inside an established component system.

[This article](https://adamwathan.me/css-utility-classes-and-separation-of-concerns/) provides great insight from the author of TailwindCSS.