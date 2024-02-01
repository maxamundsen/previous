package handlers

import (
	"gohttp/views"
	"net/http"
)

func indexHandler(w http.ResponseWriter, r *http.Request) {
	viewData := make(map[string]interface{})
	viewData["Title"] = "Index"

	model := views.NewViewModel(nil, viewData)

	// By default, any unmapped route will route to '/', so make sure
	// the URL is actually '/' or else 404
	if r.URL.Path != "/" {
		w.WriteHeader(http.StatusNotFound)
		model.ViewData["Title"] = "An error has occurred."
		model.ViewData["ErrorCode"] = "Error 404"
		model.ViewData["ErrorMsg"] = "Page not found."
		views.RenderWebpage(w, "error", model)
		return
	}

	views.RenderWebpage(w, "index", model)
}
