# Controllers

In Saral, a "controller" is simply a function that handles an HTTP request.
The controller is responsible for parsing form submissions, URL query parameters, and calling out to the appropriate functions to generate a response.
This response is typically in HTML (page) or JSON (API) form.

In Go, an HTTP handler is defined as a function that takes two arguments: a response writer, and a request pointer:

```go
func ExampleController(w http.ResponseWriter, r *http.Request) {
    // get query params
    data := r.URL.Query().Get("username")

    // parse form data
    r.ParseForm()

    info := r.FormValue("info")
    ...
}
```

## Page Controllers

A "page" is a controller that writes generated HTML to the `http.ResponseWriter`.

The concept of `views` will be covered in the next chapter, although this demonstration should be sufficient in understanding how a "page" is written to the response writer:

```go
func ExampleController(w http.ResponseWriter, r *http.Request) {
	type model struct {
		Field1 int
		Field2 string
		Field3 bool
	}

	data := model{
		Field1: 1,
		Field2: "Hello, world!",
		Field3: true,
	}

    // pass the data to the view function, which gets written to the http.ResponseWriter `w`
	ExampleView(data).Render(w)
}
```

## API Controllers

An API controller is identical to a page controller, except outputs a JSON response.

There are a two builtin functions in Saral that assist with generating JSON from structs:

```go
ApiWriteJSON(w http.ResponseWriter, data interface{})
ApiReadJSON(w http.ResponseWriter, r *http.Request, dst interface{})
```

Example of _writing_ serialized struct data to a response:
```go
func ExampleApiController(w http.ResponseWriter, r *http.Request) {
	type model struct {
		Field1 int
		Field2 string
		Field3 bool
	}

	data := model{
		Field1: 1,
		Field2: "Hello, world!",
		Field3: true,
	}

	ApiWriteJSON(w, data)
}
```

Example of _reading_ serialized data into a struct:
```go
type LoginInfo struct {
	Username string
	Password string
}

func LoginController(w http.ResponseWriter, r *http.Request) {
	var loginInfo LoginInfo

	err := ApiReadJSON(w, r, &loginInfo)
    if err != nil {
        // handle error
    }
    ...
}
```
## File System & Naming Convention

There is no enforced binding between controllers, routes, and the file system.
You are free to call any controller, from any route, where the controller lives in any package, anywhere in the file system.

However, to keep things tidy, the codebase follows a standard convention:

- All routes are mapped in the `routes.go` file.
- Page controllers live in the `/src/pages/{PAGE_GROUP}/{SPECIFIC_PAGE}.go` file.
	- Ex: A login controller would be located in `/src/pages/auth/login.go`
- HTTP handler functions have the word "Controller" suffixed.
	- Ex: `IndexController`, `LogoutController`, `DashboardController`

