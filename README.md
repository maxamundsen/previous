# Saral: A simple web codebase.

## What is Saral?
Saral is a project template / codebase that provides simple, and powerful facilities for building websites and web applications.
Saral is _not_ a framework, or library.
This project is a showcase of what is possible with simple procedural code, while avoiding needlessly overcomplicated abstractions.

At the end of the day, your job as a web developer is to dynamically generate HTML, and serve it back to a requester.
Saral was built to remove the friction that is common with popular, overengineered solutions, so you can focus on getting the job done.

The name "Saral" translates to "simple" in Hindi.

### Programming Environment
The codebase is implemented in the [Go](https://go.dev/) programming language, as it is easy to use, procedural, compiles to a static executable, and provides a featureful standard library.

If you don't like Go, or the implementations of specific codebase features, feel free to reimplement them in your language of choice, using your own personal style!
This project is a simply a template that can be used as the base for your next million dollar app, or as a reference for your own system.

## Why not some other popular framework?
Other frameworks sell the promise of a "good developer experience" without any efforts towards robustness, longevity, or under-the-hood simplicity.
Many of these frameworks update frequently, deprecate features, and change APIs, forcing you to do lots of housekeeping to keep your website running.
On top of this, frameworks often make heavy use of third-party dependencies, which dramatically increases the fragility of your project.

Again, Saral is _not_ a framework!
By using
### Advantages

- Compile to a statically-linked executable - Extremely easy to deploy to your platform of choice; docker not required.
- Statically type-checked - No need for "fake" type-checking systems (looking at you TypeScript). Just compile the code!
- Pure procedural code - No esoteric HTML template syntax, or configuration languages here! Since everything is handled server-side, you can use any function to transform your data to build an HTML output, without needing to leave Go.
- Vendored dependencies - By vendoring third party code as source, you can feel confident that you application will be long lasting, and robust.
- Simple and extendable - This is a project template, not a framework, it is very easy to rip out features you don't need, or implement your own if you wish. The code you write should be tailored to the problems _you_ are trying to solve, not fighting against strict framework guidelines / APIs.

## Is Saral a good fit for my project?
Probably!

I believe that all web applications can be this simple.
If your application relies on complicated features that extend beyond the scope of the original intention of the web, you should probably consider a native solution instead.
If your website is only serving static content with a single, or few pages, use a static site generator, or perhaps plain ol' HTML + CSS.

The template is currently geared towards authenticated CRUD applications, however it can be easily extended (or minimized) to facilitate e-commerce, blog sites, documentation, or personal sites.

## Does Saral "scale"?
Absolutely.

The simplicity of this project allows instances to easily be deployed across multiple servers.
Saral is considered "monolithic architecture", and doesn't "autoscale" with layers of complex microservices.
If this is your goal, this project is probably not a good fit.

By default, Saral does not store application state in memory, so you are safe to use a reverse-proxy to load balance incoming traffic, without worrying about loss of state.