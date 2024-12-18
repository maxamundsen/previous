package main

import (
	// "saral/middleware"
	"log"
	"net/http"
	"net/http/httptest"
	"strings"

	"saral/middleware"

	"saral/pages"
	"saral/pages/app"
	"saral/pages/app/examples"
	"saral/pages/auth"
	"saral/pages/components"
	"saral/pages/docs"

	"saral/api"
)

func mapPageRoutes(mux *http.ServeMux) {
	// create aliases for middleware so binding routes is easier
	id := middleware.LoadIdentity
	sess := middleware.LoadSessionFromCookie

	// root (index) handler
	mapIndex(mux)

	// auth handlers
	mux.HandleFunc("/auth/login", id(auth.LoginController, false))
	mux.HandleFunc("/auth/logout", id(auth.LogoutController, true))

	// app handlers
	mux.HandleFunc("/app/dashboard", id(sess(app.DashboardController), true))

	mux.HandleFunc("/app/examples/forms", id(sess(examples.FormController), true))
	mux.HandleFunc("/app/examples/api-fetch", id(sess(examples.ApiFetchController), true))
	mux.HandleFunc("/app/examples/htmx", id(sess(examples.HtmxController), true))
	mux.HandleFunc("/app/examples/htmx/counter/{count}", id(sess(examples.HtmxCounterController), true))
	mux.HandleFunc("/app/examples/alpine", id(sess(examples.AlpineController), true))
	mux.HandleFunc("/app/examples/upload", id(sess(examples.UploadController), true))
	mux.HandleFunc("/app/examples/smtp", id(sess(examples.SmtpController), true))

	mux.HandleFunc("/app/api-demo", id(sess(app.ApiDemoController), true))

	mux.HandleFunc("/app/account", id(sess(app.AccountController), true))

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

	log.Println("Mapped page routes")
}

func mapApiRoutes(mux *http.ServeMux) {
	id := middleware.LoadIdentity

	mux.HandleFunc("POST /api/auth/login", api.LoginController)

	mux.HandleFunc("/api/test", api.TestController)
	mux.HandleFunc("/api/account", id(api.AccountController, true))

	log.Println("Mapped API routes")
}

func mapIndex(mux *http.ServeMux) {
	fs := http.FileServer(http.Dir("wwwroot"))
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// serve docs page if route is literally '/'
		if r.URL.Path == "/" {
			pages.IndexController(w, r)
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
