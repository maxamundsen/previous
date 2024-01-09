//go:build embed

package views

import "embed"

//go:embed *.html
var embeddedTemplates embed.FS