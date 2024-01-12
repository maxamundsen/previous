package handlers

import (
	"gohttp/views"
	"net/http"
)

type indexModel struct {
	Base views.ViewBase
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	viewData := make(map[string]interface{})
	viewData["Title"] = "Index"
	base := views.NewViewBase(nil, viewData)

	model := indexModel{
		base,
	}

	// By default, any unmapped route will route to '/', so make sure
	// the URL is actually '/' or else 404
	if r.URL.Path != "/" {
		w.WriteHeader(http.StatusNotFound)
		model.Base.ViewData["ErrorCode"] = "Error 404"
		model.Base.ViewData["ErrorMsg"] = "Page not found."
		views.RenderTemplate(w, "error", model)
		return
	}

	views.RenderTemplate(w, "index", model)
}
