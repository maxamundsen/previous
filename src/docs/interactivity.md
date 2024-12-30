# Views: Interactivity

Contemporary websites and applications often feature "interactivity", which is a fancy term for "something changed on the screen without refreshing the whole page".
Due to the design and implementation of the web browser, interactivity requires _some_ JavaScript.
It is possible to achieve interactivity while keeping all of the benefits of server-side component generation.

## HTMX

[HTMX](https://htmx.org/) is a lightweight JavaScript shim that gets included in the `Root` "layout" component.
It allows HTTP requests to be sent by interacting with HTML elements, and swapping the response into the current DOM.

Functionally, this is no different than binding an eventlistener to an element, and triggering an XHR request.
However, HTMX provides easy-to-use semantics that describe this transaction without the verbosity of hand-written JavaScript.

### HTMX Components

Unlike regular components, which are reusable snippets of code and HTML, HTMX "components" require a route, and a controller, since you are still responding to an HTTP request.
However, unlike most other controllers, which return full page views, HTMX controllers return partial HTML.

By convention, HTMX routes are suffixed with `-hx`, and controllers are suffixed with `Hx`.

```go
// /src/routes.go
mux.HandleFunc("/app/example/users-hx", example.UsersControllerHx)
```

```go
// /src/pages/app/example.go
func UsersControllerHx(w http.ResponseWriter, r *http.Request) {
	UserList().Render(w)
}

func UserList() Node {
	return Ul(
		Li(Text("Matthew")),
		Li(Text("Mark")),
		Li(Text("Luke")),
		Li(Text("John")),
	)
}
```

Note that there is no additional layout specified at the component level.
You are purely returning an HTML fragment that will be swapped into another page.

The following is an example use of the previous HTMX component being used in a page view:

```go
func ExampleView() Node {
	return AppLayout(
		
	)
}
```