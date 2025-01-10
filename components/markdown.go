package components

import (
	. "maragu.dev/gomponents"

	"github.com/gomarkdown/markdown"
)

func Markdown(input string) Node {
	html := markdown.ToHTML([]byte(input), nil, nil)
	return Raw(string(html))
}
