package main

import (
	// "previous/middleware"
	"net/http"
	"previous/handlers"
	"previous/handlers/api"
	api_auth "previous/handlers/api/auth"
	"previous/handlers/app"
	"previous/handlers/app/examples"
	"previous/handlers/auth"
	"previous/middleware"
)

func mapRoutes(mux *http.ServeMux) {
	id := middleware.LoadIdentity
	sess := middleware.LoadSession
	cors := middleware.EnableCors

	mux.HandleFunc("/api/account", cors(id(api.AccountApiHandler, true)))
	mux.HandleFunc("/api/auth/login", id(api_auth.LoginHandler, false))
	mux.HandleFunc("/api/test", api.TestHandler)
	mux.HandleFunc("/app/account", id(sess(app.AccountHandler), true))
	mux.HandleFunc("/app/dashboard", id(sess(app.DashboardHandler), true))
	mux.HandleFunc("/app/examples/api-fetch", id(sess(examples.ApiFetchHandler), true))
	mux.HandleFunc("/app/examples/api-fetch-hx", id(examples.ApiFetchHxHandler, true))
	mux.HandleFunc("/app/examples/autotable", id(sess(examples.AutoTableHandler), true))
	mux.HandleFunc("/app/examples/autotable-hx", id(sess(examples.AutoTableHxHandler), true))
	mux.HandleFunc("/app/examples/forms", id(sess(examples.FormHandler), true))
	mux.HandleFunc("/app/examples/charts", id(sess(examples.ChartHandler), true))
	mux.HandleFunc("/app/examples/html-sanitization", id(sess(examples.HtmlSanitizationHandler), true))
	mux.HandleFunc("/app/examples/htmx", id(sess(examples.HtmxHxHandler), true))
	mux.HandleFunc("/app/examples/htmx-counter/{count}", examples.HtmxCounterHandler)
	mux.HandleFunc("/app/examples/inline-styles", id(sess(examples.InlineStylesHandler), true))
	mux.HandleFunc("/app/examples/lipsum-hx", id(sess(examples.LoremIpsumHxHandler), true))
	mux.HandleFunc("/app/examples/markdown", id(sess(examples.MarkdownHandler), true))
	mux.HandleFunc("/app/examples/smtp", id(sess(examples.SmtpHandler), true))
	mux.HandleFunc("/app/examples/inline-scripting", id(sess(examples.InlineScriptingHandler), true))
	mux.HandleFunc("/app/examples/ui-playground", id(sess(examples.UIPlaygroundHandler), true))
	mux.HandleFunc("/app/examples/upload", id(sess(examples.UploadHandler), true))
	mux.HandleFunc("/auth/login", id(sess(auth.LoginHandler), true))
	mux.HandleFunc("/auth/logout", id(sess(auth.LogoutHandler), true))
	mux.HandleFunc("/", handlers.IndexHandler)
}
