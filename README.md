# Saral: A simple web codebase.

## What is Saral?
Saral is a reference codebase that provides simple, and powerful facilities for building websites and web applications.
Saral is _not_ a framework, or library.
This project is a showcase of what a web application can look like with simple procedural code, while avoiding needlessly overcomplicated abstractions.

At the end of the day, the job of a "web developer" is to dynamically generate HTML, and serve it back to a requester.
Saral was built to remove the friction that is common with popular, overengineered solutions, so you can focus on getting the job done.

The name "Saral" translates to "simple" in Hindi.

## Programming Environment
The codebase is implemented in the [Go](https://go.dev/) programming language, as it is easy to use, procedural, compiles to a static executable, and provides a featureful standard library.

### Pre-requisites

- A copy of the [Go compiler](https://go.dev/).
- Bash shell for running included build scripts. This is typically included in Linux distributions, and MacOS, but needs t

## Getting Started / Installation
Clone the Saral repository to your local machine with the following command:
```sh
git clone https://github.com/maxamundsen/Saral.git
```

Set your working directory to the `/src` directory, and run the following command:
```sh
sh build_server.sh
```

This outputs a `server` or `server.exe` file inside the `/src` directory.
When executed, the web server will listen on TCP port `:9090`.
Navigate to the `/docs` page to read the full codebase documentation.

Any third party dependencies are included in the project.
This includes fonts, icons, Go libraries, and JS libraries.
Go libraries are vendored, and their source code is compiled alongside user-level code.
JavaScript dependencies are bundled and minified.

## Why not another solution?
Other tools and frameworks sell the promise of a "good developer experience" without any efforts towards robustness, longevity, or "real" simplicity.
Many of these frameworks update frequently, deprecate features, change APIs regularly, and have _many_ dependencies, forcing you to do _lots_ of housekeeping to keep your website running.

The code that drives your website / application should be specialized to fit the problems you are trying to solve.
By subscribing to a framework, you lock yourself into a set of rules that may not align with those problems.
A reference codebase such as Saral provides all the tools you need to get up and running quickly, while proving the flexibily to change things on your own accord.

### Primary Benefits

- Compile to a statically-linked executable - Extremely easy to deploy to your platform of choice; docker not required.
- Statically type-checked by the compiler - No need for half-baked pseudo-type-checking systems.
- Procedural code - No esoteric template / configuration syntax required. Pages are written using pure, procedural Go functions.
- Vendored dependencies - By vendoring third party code as source, you can feel confident that you application will be long lasting, and robust.
- Simple and extendable - This project serves as a reference for your individual needs. You are expected to extend or reduce this codebase as needed. For example, if you are building a simple personal site, you may not need a database, or authentication system! Just delete those parts!
