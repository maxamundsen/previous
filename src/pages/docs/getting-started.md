# Getting Started / Installation
First to get started, you'll need to get started.
Clone the Saral repository to your local machine with the following command:
```sh
git clone https://github.com/maxamundsen/Saral.git
```

Compiling the codebase requires a copy of the [Go compiler](https://go.dev).
Check the `go.mod` file located in the `/src` directory for the minimum compiler version required.

Any third party dependencies are included in the project.
This includes fonts, icons, Go libraries, and JS libraries.
Go libraries are vendored, and their source code is compiled alongside user-level code.
Javascript dependencies are bundled and minified.
