package components

import (
	"previous/basic"
	"previous/security"
	"strings"

	. "maragu.dev/gomponents"
)

// @Macro - At runtime this inserts an attribute component with the hash of the input
// At compile time, the input is collected, and used to generate a global css file.
// WARNING: THE PREPROCESSOR EXPANSION ONLY WORKS WITH *STRING LITERAL* INPUTS
// ATTEMPTING TO USE THIS WITH DYNAMIC INPUTS WILL GENERATE GARBAGE
func InlineStyle(input string) Node {
	input = strings.ReplaceAll(input, "\n", " ")
	input = strings.ReplaceAll(input, "\t", "")

	s, _ := security.HighwayHash58(input)
	s = basic.GetFirstNChars(s, 8)

	return Attr("__inlinecss_" + s)
}
