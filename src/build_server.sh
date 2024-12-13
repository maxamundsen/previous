#!/bin/sh

set -e

./tailwindcss-macos-arm64 -i styles/global.css -o wwwroot/css/style.css --minify

case "$1" in
  	run_debug)
		dlv exec server
    ;;
  	debug)
		go build -tags=debug -gcflags=all="-N -l" ./cmd/server
    ;;
    *)
  		go build ./cmd/server
    ;;
esac