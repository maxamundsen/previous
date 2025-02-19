package examples

import (
	"net/http"
	. "previous/components"

	. "maragu.dev/gomponents"
)

func LoremIpsumHxHandler(w http.ResponseWriter, r *http.Request) {
	func() Node {
		return Text(LOREM_IPSUM)
	}().Render(w)
}
