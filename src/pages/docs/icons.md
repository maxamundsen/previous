# Views: Icons

Icons are handled quite differently than other approaches.
Instead of using font files, images, or SVG files, icons are stored as key-value pair of strings to strings, exposed as a component function that can be called with a desired size.

```go
var icons = map[string]string{
	"menu": `<svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round" class="lucide lucide-menu"><line x1="4" x2="20" y1="12" y2="12"/><line x1="4" x2="20" y1="6" y2="6"/><line x1="4" x2="20" y1="18" y2="18"/></svg>`,
	...
}
```

Using icons:

```go
func MyComponent() Node {
	return Div(
		Icon("menu", 24),
	)
}
```

## How to add new icons

Navigate to the `/src/pages/components/icon.go` file, and paste the SVG as a string entry in the map.
You may need to set the `viewBox` property to `0 0 24 24` in order for things to scale properly.

## Colored Icons
Icons can be colored if `fill="currentColor"`, or `stroke="currentColor"` is set (this is dependent on the icon you are using).

```go
func MyComponent() Node {
	return Div(Class("text-red-500"),
		Icon("users", 24), // this icon will draw in red
	)
}
```

## Why?
Although this approach is unconventional, it solves the following problems:

- Icons are not bound to a large font bundle.
- Icons are embedded directly in the HTML output after the rendering process.
- No need for separate HTTP requests to load many icons on a page.
- External dependencies -- Since icons are represented as a string in your code, you do not need to fetch them from a third party at any point.
