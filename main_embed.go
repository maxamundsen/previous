//go:build embed

package main

import "embed"

//go:embed wwwroot
var staticAssets embed.FS
