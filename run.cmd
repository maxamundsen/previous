@echo off
setlocal enabledelayedexpansion

if "%1" == "build" (
    go build
) else if "%1" == "build_embed" (
    go build -tags embed
) else if "%1" == "debug" (
	dlv debug --build-flags "-tags=devel"
) else if "%1" == "fmt" (
	go fmt ./...
) else (
    echo Available options: build, build_embed, debug, fmt
)

endlocal