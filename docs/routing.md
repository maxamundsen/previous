# Routing
A `route` is a mapping between a specified URL, and a `controller`, wher a controller is a function that handles an HTTP request.
You can think of a `route` as the answer to the question: "What code is executed when I navigate to this URL?".

Routes can be either be specified manually, or automatically, via the filesystem (generated from `metagen`).

A basic `routes.go` file looks something like this:
```go
package main

import (
    "previous/pages/public"
)

func mapPageRoutes(mux *http.ServeMux) {
    mux.HandleFunc("/public/about", public.AboutController)
}
```

The function `mapPageRoutes()` is called from the `main()` function inside `main.go`, which maps user-defined routes.

## Index / Static Content

The route for the site root `/` is unique, because it handles `404 not found` errors, and static assets.
There is a builtin `mapIndex()` function for this special case that is called from `mapPageRoutes()`.

Routes matching `/` literally, will call the index controller.
If the route does not match `/`, or any of the previously mapped routes, it will try to serve a file from `/src/wwwroot` directory, matching the route specified.
If the requested static asset is not found, this route will call the "404 page not found" controller.

Example:

Suppose `/app/dashboard` is mapped to a controller.
The route `/app/test` is unmapped.
Requests to `/app/test` are captured by the index route.
Since `/app/test` does not match the literal `/` route, the server attempts to serve the file located in `/src/wwwroot/app/test`.
Because there is no file found here, it will respond with a`404` page instead.

## Middleware
Controllers can be _composed_ to allow additional "preflight" logic before sending the final response to the requester.
This is known as a "middleware chain".
You can add as many controllers as you want to the chain, to "forward" the request to the next controller in the chain.

The most common use of middleware is authentication logic.
If you want to "protect" many routes from non-authenticated users, you would need to include the authentication logic in each controller.
With middleware, individual routes do not need to include this boilerplate, since you can call them _after_ the authentication middleware route.

## API Routes

There isn't anything that differenciates an "API route" from an HTML page route, besides your own personal organization.
The codebase contains a `mapApiRoutes()` function for organizational purposes only.

```go
func mapApiRoutes(mux *http.ServeMux) {
    mux.HandleFunc("POST /api/auth/login", api.Login)
}
```