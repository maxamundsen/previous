package main

import (
	// "saral/middleware"
	"net/http"
	"net/http/httptest"
	"strings"

	"saral/components"
	"saral/docs"
)

// manually mapped routes->controllers go here.
func mapDocumentationRoutes(mux *http.ServeMux) {
	// docs handlers
	mux.HandleFunc("/docs", docs.IndexController)

	// dynamically map documentation pages
	for _, v := range docs.DocList {
		if len(v.SubList) == 0 {
			if v.Slug != "" {
				mux.HandleFunc("/docs/"+v.Slug, docs.DocController)
			}
		} else {
			for _, k := range v.SubList {
				if k.Slug != "" {
					mux.HandleFunc("/docs/"+k.Slug, docs.DocController)
				}
			}
		}
	}
}

// index controller is handled specially
func mapIndexRoute(mux *http.ServeMux) {
	fs := http.FileServer(http.Dir("wwwroot"))
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// serve docs page if route is literally '/'
		if r.URL.Path == "/" {
			http.Redirect(w, r, "/docs", http.StatusFound)
			return
		}

		// By default, any unmapped route will route to '/', so make sure
		// the URL is actually '/' or else 404
		if strings.HasSuffix(r.URL.Path, "/") {
			w.WriteHeader(http.StatusNotFound)

			components.ErrorPage(http.StatusNotFound).Render(w)
			return
		}

		rr := &httptest.ResponseRecorder{Code: http.StatusOK}

		// otherwise, serve a static file (assuming it exists)
		fs.ServeHTTP(rr, r)

		if rr.Code != http.StatusOK {
			w.WriteHeader(rr.Code)
			components.ErrorPage(rr.Code).Render(w)
		} else {
			fs.ServeHTTP(w, r)
		}
	})
}
