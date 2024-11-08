# Getting Started / Installation
First to get started, you'll need to get started.
Clone the WebDawgEngine repository to your local machine with the following command:
```shell
git clone https://github.com/maxamundsen/WebDawgEngine.git
```

## Building the codebase
Building this repo requires invoking the Go compiler, which can be found [here](https://go.dev/).

First, move to the `/src` directory.
This is where our Go source code is located.
```
cd ./src
```

To build the demo site UNIX-like machines:
```
sh build.sh
```

For Windows machines:
```
build.bat
```

This will output a statically-linked binary called `webdawgengine`, located in the `./src` directory.

## Running the program
To execute the program, run the `webdawgengine` executable.

UXIX-like:
```
./webdawgengine
```

Windows:
```
webdawgengine.exe
```

By default, WDE will run on port `8080`.
This can be modified in the configuration file, which is discussed in a later chapter.
