//go:build embed

package main

import "embed"

//go:embed wwwroot/favicon.ico
var content embed.FS

//go:embed wwwroot/assets
var staticAssets embed.FS
