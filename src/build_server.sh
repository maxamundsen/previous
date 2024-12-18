#!/bin/sh

set -e

case "$1" in
  	run_debug)
		dlv exec server
    ;;
  	debug)
		./tailwindcss-macos-arm64 -i styles/global.css -o wwwroot/css/style.css --minify && go build -tags=debug -gcflags=all="-N -l" ./cmd/server
    ;;
    *)
  		./tailwindcss-macos-arm64 -i styles/global.css -o wwwroot/css/style.css --minify && go build ./cmd/server
    ;;
esac