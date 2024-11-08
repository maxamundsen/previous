# Build System

## Building
WDE can be built by directly invoking the Go compiler from inside the `./src` directory:
```sh
go build
```

This will generate an executable in the same directory called `webdawgengine`.
Modifying the first line in the `go.mod` file (located in the `./src` directory) will change the module name, and thus the executable name.

### Build Constants
The Go programming language does not have in-depth compile-time execution, so

## Deploying