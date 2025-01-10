package pages

import (
	. "previous/components"

	"net/http"
	"net/http/httptest"
	"strings"
)


func IndexController(w http.ResponseWriter, r *http.Request) {
	// serve home page if route is literally '/'
	if r.URL.Path == "/" {
		http.Redirect(w, r, "/docs", http.StatusFound)
		return
	}

	// By default, any unmapped route will route to '/', so make sure
	// the URL is actually '/' or else 404
	if strings.HasSuffix(r.URL.Path, "/") {
		w.WriteHeader(http.StatusNotFound)

		ErrorPage(http.StatusNotFound).Render(w)
		return
	}

	rr := &httptest.ResponseRecorder{Code: http.StatusOK}

	// otherwise, serve a static file (assuming it exists)
	fs.ServeHTTP(rr, r)

	if rr.Code != http.StatusOK {
		w.WriteHeader(rr.Code)
		ErrorPage(rr.Code).Render(w)
	} else {
		fs.ServeHTTP(w, r)
	}
}