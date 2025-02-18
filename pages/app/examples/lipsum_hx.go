package examples

import (
	"net/http"
	. "previous/components"

	. "maragu.dev/gomponents"
)

func LoremIpsumHxPage(w http.ResponseWriter, r *http.Request) {
	func() Node {
		return Text(LOREM_IPSUM)
	}().Render(w)
}
