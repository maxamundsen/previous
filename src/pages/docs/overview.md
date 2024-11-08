# WebDawgEngine Documentation
Welcome to the WebDawgEngine (WDE) documentation!
This is indended to be a full guide for using the WDE system.

The documentation is written
## What is WebDawgEngine?
WebDawgEngine (WDE) is a project template / design methodology that provides simple, yet powerful tools for building websites and web applications.
WebDawgEngine is _not_ a framework, or library.
WDE is a showcase of what is possible with procedural, data-oriented code, while avoiding needlessly overcomplicated abstractions.

The name "WebDawgEngine" pays homage to [PyDawgEngine](https://github.com/RyDawg-Studios/PyDawgEngine), a high-level framework for quickly building games in Python.

At the end of the day, your job as a web developer is to dynamically generate HTML, and serve it back to the requester.
WDE was built to remove the friction that is common with popular, overengineered solutions, so you can focus on getting work done.

WDE is written in the [Go](https://go.dev/) programming language, as it , while providing a featureful standard library.
In fact, most of the code in WDE is just a wrapper around the Go standard library!

If you don't like Go, or the implementations of WDE features, feel free to reimplement them in your language of choice, using your own personal style!
This project is a simply a template that can be used as the base for your next million dollar app, or as a reference for the next killer web framework (we seriously need more of these I promise).

## Why not some other popular framework?
Other frameworks sell the promise of "a nice developer experience" without any efforts towards long term stability, or under-the-hood simplicity.
Many of these frameworks update frequently, deprecate features, and change APIs, forcing you to do a ton of housekeeping to keep your app running.
On top of this, frameworks often make heavy use of third-party dependencies, which severely increases the fragility of your project.

### Main Advantages

- Compile to a static binary - Extremely easy to deploy to your platform of choice; docker not required.
- Statically type-checked - No need for additional type-checking systems, just use the compiler!
- Pure procedural code - No esoteric template syntax, or configuration languages here!
- Vendored dependencies - By vendoring third party code, you can feel confident that you application will be long lasting, and robust.

### Disadvantages

- Minimal integration with existing component libraries or auth libraries (although you probably don't need these things).
-
- Robustness and longevity may hinder job security as you don't have to update your project constantly.

## Is the WDE system a good fit for my project?
Maybe!

We believe that all web applications can be this simple.
Seriously.
If your application relies on complicated features that extend beyond the scope of the original web vision, the web might not be a suitable platform, and you should consider a traditional native application.

If your website is only serving static content with a single, or few pages, just use plain ol' HTML + CSS!