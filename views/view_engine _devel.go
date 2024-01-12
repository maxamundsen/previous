//go:build devel

package views

import "embed"

//go:embed *.html
var embeddedTemplates embed.FS
