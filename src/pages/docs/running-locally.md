# Running a WDE Server Locally

After building the WDE project, you can run the server locally by executing the following inside a terminal:

UXIX-like:
```sh
./webdawgengine
```

Windows:
```
webdawgengine.exe
```

By default, WDE will run on port `8080`.
This can be modified in the configuration file, which is discussed in a later chapter.

## Debugging
You can use the [delve](https://github.com/go-delve/delve) debugger to debug the project.

```sh
dlv debug --build-flags "-tags=debug"
```