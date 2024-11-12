# Testing & Benchmarking Your Project

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