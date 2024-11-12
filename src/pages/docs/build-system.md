# Build System

## Building
WDE can be built by directly invoking the Go compiler from inside the `./src` directory:
```sh
go build
```

This will generate a statically-linked executable in the same directory called `webdawgengine`.
Modifying the first line in the `go.mod` file (located in the `./src` directory) will change the module name, and thus the executable name.

### Build Script
By default WDE ships with build scripts, `build.bat` and `build.sh`, located in the `/src` directory.
These scripts are useful for running multiple build commands (if necessary).

To build with the build script, run:

MacOS / Linux:
```sh
sh build.sh
```

Windows:
```
build.bat
```

### Build Constants
The Go programming language can conditionally include code at _compile time_ by providing a condition in a comment at the top of a file.

```sh
go build -tags=debug
```