package components

import (
	"previous/basic"
	"previous/security"
	"strings"

	. "maragu.dev/gomponents"
)

// @Macro - At runtime this inserts an attribute component with the hash of the input
// At compile time, the input is collected, and used to generate a global css file.
func InlineStyle(input string) Node {
	input = strings.ReplaceAll(input, "\n", " ")
	input = strings.ReplaceAll(input, "\t", "")

	s, _ := security.HighwayHash58(input)
	s = basic.GetFirstNChars(s, 8)

	return Attr("__inlinecss_" + s)
}