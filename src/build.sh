#!/bin/sh

set -e

./tailwindcss-macos-arm64 -i styles/global.css -o wwwroot/css/style.css --minify

/usr/local/go/bin/go build -tags=debug