# Previous: A simple web codebase.

## What is Previous?
Previous is a reference codebase that provides simple, and powerful facilities for building websites and web applications.
Previous is _not_ a framework, or library.
This project is a showcase of what a web application can look like with simple procedural code, while avoiding needlessly overcomplicated abstractions.

At the end of the day, the job of a "web developer" is to dynamically generate HTML, and serve it back to a requester.
Previous was built to remove the friction that is common with popular, overengineered solutions, so you can focus on getting the job done.

## Programming Environment
This project is implemented in the [Go](https://go.dev/) programming language, as it is easy to use, procedural, compiles to a static executable, and provides a featureful standard library.

### Pre-requisites

- The [Go compiler](https://go.dev/).

## Getting Started / Installation
Clone the Previous repository:
```sh
git clone https://github.com/maxamundsen/Previous.git
```

Run the metaprogram:
```sh
go run ./cmd/metagen --env=dev build
```

Compile and run the server application:
```sh
go build ./cmd/server && ./server
```

When executed, the web server will listen on TCP port `:9090`.
This can be changed in the [config](/docs/configuration).


Any third party dependencies are included in the project.
This includes fonts, icons, Go libraries, and JS libraries.
Go libraries are vendored, and their source code is compiled alongside user-level code.
JavaScript dependencies are bundled and minified.

## Why not another solution?
Other tools and frameworks sell the promise of a "good developer experience" without any efforts towards robustness, longevity, or "true" simplicity.
Many of these frameworks deprecate features, change APIs regularly, and have _many_ dependencies, forcing you to do _lots_ of housekeeping to keep your website running.

The code that drives your website / application should be specialized to fit the problems you are trying to solve.
By subscribing to a framework, you lock yourself into a set of rules that may not align with those problems.
In contrast to a framework, a reference codebase provides all the tools you need to get up and running quickly, while proving the means to extend the code to fit your project.

### Primary Benefits

- Compile to a statically-linked executable - Extremely easy to deploy to your platform of choice; docker not required.
- Everything is statically type-checked by the compiler - This includes HTML components, and database queries.
- Procedural code - No esoteric template / configuration syntax required. Pages are written using pure, procedural Go functions.
- Vendored dependencies - By vendoring third party code as source, you can feel confident that you application will be long lasting, and robust.
- Simple and extendable - This project serves as a reference for your individual needs. You are expected to extend or reduce this codebase as needed. For example, if you are building a simple personal site, you may not need a database, or authentication system! Just delete those parts!
