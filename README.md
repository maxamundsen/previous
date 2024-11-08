# WebDawgEngine

WebDawgEngine is an extremely simple and lightweight "framework" for creating dynamic websites using the Go programming language. This framework is designed for creating _useful and functional_ websites.

The name "WebDawgEngine" pays homage to [PyDawgEngine](https://github.com/RyDawg-Studios/PyDawgEngine), a high-level framework for quickly building games in Python.

## Included Functionality
WebDawgEngine includes the following functionality out of the box:

- **Dynamic routing** : Define routes to be handled by simple handler procedures
- **Static Routing** : Serve static assets from `/wwwroot/assets` directory
- **View Engine**
  - Wrapper around Go standard library HTML templates
    - Templates are embedded in the binary when compiled for a release build.
    - Templates can be edited and previewed without recompiling in debug builds.
  - ViewModel & ViewData structures for easy view data management
  - Integration with the framework's identity system for easily displaying user information
  - Integration with HTMX for partial HTML content swapping
  - Serve view directly through HTTP pipeline
  - Output compiled view as byte buffer
- **JSON**
  - Automatic serialization/de-serialization provided by the standard library
  - Respond with, and accept JSON with handler functions
- **Session based authentication / authorization**
  - Session store interface definition
    - Create, fetch, and delete sessions
    - Session store key is stored and transmitted via a secure cookie
  - In-memory session store implementation
  - MySQL session store implementation
- **Identity System** : Data structures for simple user & permissions management
- **Database Access** (Preconfigured for MySQL)
- **File Uploads** : Handle uploaded files from a route handler
-  **SMTP Client**
   - Built on top of standard library SMTP client
   - View Engine integration for mail templates
- **Configuration File** (`config.json`)
  - Program configuration parsed at runtime during program startup
  - Optional `config.devel.json` options for debug builds.
  - Console warnings for missing configuration values

Each of these modules are extremely simple, and can be modified, or removed depending on the needs of the application.

## Compiling and Debugging
### Release Build
To compile the application, install the go compiler, and execute the following from the project directory:

```sh
go build
```

You should now have a binary called `webdawgengine.exe`. Compiling this way will yield a "Release mode" binary, which is ready for deployment.

Note that a `config.json` file is required for the application to start. When initializing static routes, `/wwwroot` must also be present.

### Debug Build
If you want to build a debug binary, install the dlv debugger, and execute the following from the project directory:

```sh
dlv debug --build-flags "-tags=devel"
```

## Basic Usage & Project Structure
Once you have the build environment setup, you can begin customizing the project for your specific needs.

Inspecting `main.go`, you will find calls to initialize structures such as the http multiplexer, routes, configuration, and database objects.

The `handlers` package contains `dynamic_routes.go` and `static_routes.go`. These contain functions that map URLs to HTTP handler functions. Any asset that lives in `/wwwroot/assets` will get mapped by the static route handler. All http handler functions should live inside this package as well.

HTTP handler functions are extremely simple, and can manipulate an HTTP request, and response writer directly. The response writer can be passed to the View Engine™ to easily integrate HTML templates into your HTTP response.

The `managers` package houses core business logic. Data fetched from the database should not be transformed within a handler, but rather in a manager.

Features such as `snailmail` (the SMTP server), or `database` live in their own package, and directory in the project. There is no enforced structure here, so any package can call out to any other package. Typically you will want to call a package such as `auth` or `database` from an HTTP handler in the handlers package. Example structure provided here:

```
1.
HTTP request -> handler function -> database -> manager -> view engine -> write to http response and return from handler function -> request completed

2.
Fetch user request -> handlers.userHandler() -> database.GetAllUsers() ->  managers.UppercaseEmails() -> views.RenderWebpage() -> return response
```

The example handlers demonstrate simple tasks that are typically managed by handler functions.

The source code is extremely easy to read, and includes comments to provide additional context when necessary.
## Dependencies

_NOTE: All dependencies are already included in this repository._

This framework is mostly built using the Go standard library, and only uses external libraries for MySQL connectivity, and encryption. These packages are currently included in the distribution in the `/vendors` directory.

The HTMX JavaScript library is included in the `/wwwroot/assets/lib` directory to facilitate unobtrusive, partial HTML loading using `XMLHttpRequest()` without the need for full a full document load. Without the need to write any JavaScript, elements on a page can be made "interactive" by including HTML attributes. HTMX only carries out the bare minimum client-side scripting to enable this, and the application state is still held by the server.

Bootstrap CSS is included in the `/wwwroot/assets/lib` directory. It is a helpful (but heavyweight) solution for styling websites with minimal effort.

## Programming Style
This framework is written in a procedural style, with the addition of lightweight interface structures with attached methods. The framework provides minimal abstractions on top of the standard library, to allow you to write simple and maintainable code, without the mental overhead of over-engineered solutions.

Unfortunately this means `IdentityService.UserDTOAuthenticationIdentityManagerInterfaceInjectableSingletonFactory()` is not included in the framework.

## Philosophy
It is an undeniable fact that data-heavy business / enterprise software favors the web platform. There is no need for client installations, and updates to the software are continually delivered without the need for client input.  For applications that contact a centralized data source, such as a relational database, this model is quite effective. The HTTP protocol provides a _simple_ means of handling input and output in a text format that is easily handled by a server.

HTML is a markup language designed to be used alongside the HTTP protocol. It is a text format that is rendered by a web browser to display formatted text, and provides basic controls for sending data back to a server. HTML is served to the client over HTTP, and the data displayed on that page should reflect the state present on the server. This implies that HTML (and thus the server) _is_ the source of truth regarding application state. _This is how the web was designed_.

**If this model does not fit your application, please consider other solutions such as a native desktop program**.

JavaScript is a _scripting_ language designed to modify web pages after they have been loaded by the client. Ignoring the quirks of the language, it can, in theory, be used to build client-side programs of arbitrary complexity. This has been the case for the past ~10 years, where applications have been written exclusively in JavaScript, and communicate back to the server via JSON APIs. These applications must manage state on both the client, and the server.

Because JavaScript was not designed to be used this way, various developer tools and language subsets have been created to "enhance" the developer experience. These additional tools kill the robustness of an application, as they have fragile, unstable APIs, and ludicrous dependency trees. At the end of the day, you are still just making a website with dynamic content, which _eventually_ gets rendered to HTML.

The goal of this project, is to facilitate the generating of dynamic HTML, and serving it to a client over HTTP. The project is supposed to be extremely flexible and hackable, so feel free to add or remove any part of it to fit your specific needs.
