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
The Go programming language can conditionally include code at _compile time_ by including a special comment at the top of a file.
This feature is similar to the static `#if` in Jai, or `#ifdef` in C/C++, although less ergonomic.

WDE provides a `build` package, which only exists to set the `DEBUG` constant at compile time.

When compiling the program with the following command, the `DEBUG` constant is true. Else, it is false.

```sh
go build -tags=debug
```

```go
func main() {
	if build.DEBUG {
		fmt.Println("DEBUG BUILD")
	} else {
		fmt.Println("RELEASE BUILD")
	}
}
```

Program Output:
```
> DEBUG BUILD
```