package handlers

import (
	"net/http"
	"webdawgengine/views"
)

func indexHandler(w http.ResponseWriter, r *http.Request) {
	viewData := make(map[string]interface{})
	viewData["title"] = "Index"

	model := views.NewViewModel(nil, viewData)

	// By default, any unmapped route will route to '/', so make sure
	// the URL is actually '/' or else 404
	if r.URL.Path != "/" {
		w.WriteHeader(http.StatusNotFound)
		model.ViewData["title"] = "An error has occurred."
		model.ViewData["error_code"] = "Error 404"
		model.ViewData["error_msg"] = "Page not found."
		views.RenderWebpage(w, "error", model)
		return
	}

	views.RenderWebpage(w, "index", model)
}
