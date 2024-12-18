# Routing

The codebase keeps routing simple by placing all routes in a single file `routes.go`.
This allows routes (and subsequently the code behind the route) to be looked up with one search in your editor.

A basic routes file looks something like this:
```go
package main

import (
    "saral/pages/public"
)

func mapPageRoutes(mux *http.ServeMux) {
    mux.HandleFunc("/public/about", public.AboutController)
}
```

The function `mapPageRoutes()` is called from the `main()` function inside `main.go`.

## Index / Static Content

The route for the site root `/` is special, since it also handles `Status 404` errors, and serving static files.
There is a builtin `mapIndex()` function for this special case that is called from `mapPageRoutes()`.

This route will first check if the URL is literally `/`.
If that is the case, it will call the controller, or redirect specified.
If the route is not `/`, it will search the `/src/wwwroot` directory for the specified file, with the path matching the route, and serve that file.
If there is no such file (or there is an error reading the file), the error page is shown.

## API Routes

There isn't anything that differenciates an "API route" from an HTML page route, besides your own personal organization.
The codebase contains a `mapApiRoutes()` function for organizational purposes only.

```go
func mapApiRoutes(mux *http.ServeMux) {
    mux.HandleFunc("POST /api/auth/login", api.Login)
}
```