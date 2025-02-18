package main

import (
	// "previous/middleware"
	"net/http"
	"previous/middleware"
	"previous/pages"
	"previous/pages/api"
	"previous/pages/api/auth"
	"previous/pages/app"
	"previous/pages/app/examples"
	auth1 "previous/pages/auth"
)

func mapRoutes(mux *http.ServeMux) {
	id := middleware.LoadIdentity
	sess := middleware.LoadSession
	cors := middleware.EnableCors

	mux.HandleFunc("/api/account", cors(id(api.AccountApiPage, true)))
	mux.HandleFunc("/api/auth/login", id(auth.LoginPage, false))
	mux.HandleFunc("/api/test", api.TestPage)
	mux.HandleFunc("/app/account", id(sess(app.AccountPage), true))
	mux.HandleFunc("/app/dashboard", id(sess(app.DashboardPage), true))
	mux.HandleFunc("/app/examples/api-fetch", id(sess(examples.ApiFetchPage), true))
	mux.HandleFunc("/app/examples/api-fetch-hx", id(examples.ApiFetchHxPage, true))
	mux.HandleFunc("/app/examples/autotable", id(sess(examples.AutoTablePage), true))
	mux.HandleFunc("/app/examples/autotable-hx", id(sess(examples.AutoTableHxPage), true))
	mux.HandleFunc("/app/examples/forms", id(sess(examples.FormPage), true))
	mux.HandleFunc("/app/examples/html-sanitization", id(sess(examples.HtmlSanitizationPage), true))
	mux.HandleFunc("/app/examples/htmx", id(sess(examples.HtmxPage), true))
	mux.HandleFunc("/app/examples/htmx-counter/{count}", examples.HtmxCounterPage)
	mux.HandleFunc("/app/examples", id(sess(examples.IndexPage), true))
	mux.HandleFunc("/app/examples/inline-styles", id(sess(examples.InlineStylesPage), true))
	mux.HandleFunc("/app/examples/lipsum-hx", id(sess(examples.LoremIpsumHxPage), true))
	mux.HandleFunc("/app/examples/markdown", id(sess(examples.MarkdownPage), true))
	mux.HandleFunc("/app/examples/smtp", id(sess(examples.SmtpPage), true))
	mux.HandleFunc("/app/examples/surreal", id(sess(examples.AlpinePage), true))
	mux.HandleFunc("/app/examples/ui-playground", id(sess(examples.UIPlaygroundPage), true))
	mux.HandleFunc("/app/examples/upload", id(sess(examples.UploadPage), true))
	mux.HandleFunc("/auth/login", id(sess(auth1.LoginPage), true))
	mux.HandleFunc("/auth/logout", id(sess(auth1.LogoutPage), true))
	mux.HandleFunc("/", pages.IndexPage)
}

