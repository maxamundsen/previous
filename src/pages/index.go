package pages

import (
	"net/http"
)

func IndexController(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "/docs", http.StatusFound)
}
