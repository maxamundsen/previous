package components

import (
	. "maragu.dev/gomponents"
	. "maragu.dev/gomponents/html"

	"strings"
)

// Automatically uses inline styles from
// the javascript extension
func InlineStyle(input string) Node {
	// minify
	input = strings.ReplaceAll(input, "\n", "")
	input = strings.ReplaceAll(input, "\t", "")

	return StyleEl(Raw("me {" + input + "}"))
}

func StyleReset() Node {
	return StyleEl(Raw("me \\* {all: unset;}"))
}


func InlineStylePseudo(pseudo string, input string) Node {
	// minify
	input = strings.ReplaceAll(input, "\n", "")
	input = strings.ReplaceAll(input, "\t", "")

	return StyleEl(Raw("me" + pseudo + " { " + input + " }"))
}
