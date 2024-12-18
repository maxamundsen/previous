# Build System

## Building
To build the project, simply run the included build script.
By default, it will build a release build, with all debug symbols removed.
```sh
sh build_server.sh
```

This will generate a statically-linked executable in the same directory called `saral`.
Modifying the first line in the `go.mod` file (located in the `./src` directory) will change the module name, and thus the executable name.

### Build Constants
The Go programming language can conditionally include code at _compile time_ by including a special comment at the top of a file.
To use code in these specially commented files, you must pass the specified flag to the compiler.
This feature is similar to the static `#if` in Jai, or `#ifdef` in C/C++.

In order to make using this feature ergonomic, the codebase includes a `build` package, which exists to set the value of the `DEBUG` constant at _compile time_.

To build the server application in debug mode, use the following command:

```sh
sh build_server.sh debug
```

This code example demonstrates how to use the compile time `DEBUG` constant:

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

## Testing & Benchmarking Your Project

Go offers out-of-the-box tooling for testing and benchmarking your code.
Any source file suffixed with `_test` will be considered when running the test suite.

To run all tests:
```sh
go test -v ./...
```

To run all benchmarks:
```sh
go test -bench=. ./...
```