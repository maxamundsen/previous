# Applications

This codebase contains multiple applications internally.

Each application contains a `main.go` file, containing a `main()` function, which serves as the entrypoint for the running process.

## Build System / Metaprogram
The codebase metaprogram is responsible for generating code, and preprocessing data.
Without
## Web Server
The primary application in the codebase is the web server.
This application is responsible for handling HTTP requests, and generating HTML, or JSON to serve to the requester.

## SQL Migrator
The migration utility allows sql scripts to be pushed to a remote server, while tracking version information.

## Password Generator
Simple utility that generates passwords.