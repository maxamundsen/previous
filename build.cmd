@echo off
setlocal enabledelayedexpansion

if "%1" == "" (
    go build
) else if "%1" == "debug" (
    gdlv -tags "devel" debug
) else if "%1" == "debug_cli" (
	dlv debug --build-flags "-tags=devel"
) else if "%1" == "fmt" (
	go fmt ./...
) else (
    echo Available options: debug, debug_cli, fmt
)

endlocal