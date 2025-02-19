package main

import (
	// "previous/middleware"
	"net/http"
	"previous/middleware"
	"previous/handlers"
	"previous/handlers/api"
	api_auth "previous/handlers/api/auth"
	"previous/handlers/app"
	"previous/handlers/app/examples"
	"previous/handlers/auth"
)

func mapRoutes(mux *http.ServeMux) {
	id := middleware.LoadIdentity
	sess := middleware.LoadSession
	cors := middleware.EnableCors

	mux.HandleFunc("/api/account", cors(id(api.AccountApiPage, true)))
	mux.HandleFunc("/api/auth/login", id(api_auth.LoginPage, false))
	mux.HandleFunc("/api/test", api.TestPage)
	mux.HandleFunc("/app/account", id(sess(app.AccountPage), true))
	mux.HandleFunc("/app/dashboard", id(sess(app.DashboardPage), true))
	mux.HandleFunc("/app/examples/api-fetch", id(sess(examples.ApiFetchPage), true))
	mux.HandleFunc("/app/examples/api-fetch-hx", id(examples.ApiFetchHxPage, true))
	mux.HandleFunc("/app/examples/autotable", id(sess(examples.AutoTablePage), true))
	mux.HandleFunc("/app/examples/autotable-hx", id(sess(examples.AutoTableHxHandler), true))
	mux.HandleFunc("/app/examples/forms", id(sess(examples.FormPage), true))
	mux.HandleFunc("/app/examples/html-sanitization", id(sess(examples.HtmlSanitizationPage), true))
	mux.HandleFunc("/app/examples/htmx", id(sess(examples.HtmxPage), true))
	mux.HandleFunc("/app/examples/htmx-counter/{count}", examples.HtmxCounterPage)
	mux.HandleFunc("/app/examples/inline-styles", id(sess(examples.InlineStylesHandler), true))
	mux.HandleFunc("/app/examples/lipsum-hx", id(sess(examples.LoremIpsumHxHandler), true))
	mux.HandleFunc("/app/examples/markdown", id(sess(examples.MarkdownHandler), true))
	mux.HandleFunc("/app/examples/smtp", id(sess(examples.SmtpHandler), true))
	mux.HandleFunc("/app/examples/surreal", id(sess(examples.SurrealHandler), true))
	mux.HandleFunc("/app/examples/ui-playground", id(sess(examples.UIPlaygroundHandler), true))
	mux.HandleFunc("/app/examples/upload", id(sess(examples.UploadHandler), true))
	mux.HandleFunc("/auth/login", id(sess(auth.LoginHandler), true))
	mux.HandleFunc("/auth/logout", id(sess(auth.LogoutHandler), true))
	mux.HandleFunc("/", handlers.IndexHandler)
}

