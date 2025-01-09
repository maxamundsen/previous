package app

import (
	. "maragu.dev/gomponents"
	. "maragu.dev/gomponents/html"
	. "previous/components"

	"previous/middleware"
	"previous/auth"

	"net/http"
)

// @Identity
// @Protected
// @CookieSession
func ApiDemoController(w http.ResponseWriter, r *http.Request) {
	identity := middleware.GetIdentity(r)

	ApiDemoView(*identity).Render(w)
}

func ApiDemoView(identity auth.Identity) Node {
	return AppLayout("API Demo", identity,
		Div(
			Attr("x-data", `
				{
					username: "",
					password: "",
					token: "",
					response: "",
					error: "",
					route: "/api/test",

					getToken() {
						fetch("/api/auth/login", {
							headers : {
								'Content-Type': 'application/json'
							},
							method: 'POST',
							body: JSON.stringify({
								Username: this.username,
								Password: this.password,
							})
						})
					   .then(resp => {
					        if (!resp.ok) {
					            throw new Error(resp.status);
					        }
					        return resp.text();
					   })
					   .then(text => {this.token = text; this.error = "";})
					   .catch(error => this.error = error)
					},

					sendRequest() {
						fetch(this.route, {
						  headers: {Authorization: 'Bearer ' + this.token}
						})
					   .then(resp => {
					        if (!resp.ok) {
					            throw new Error(resp.status);
					        }
					        return resp.json();
					   })
					   .then(json => { this.response = JSON.stringify(json); this.error = ""; })
					   .catch( error => this.error = error)
					},
				}
			`),
			Form(Attr("x-on:submit.prevent", "getToken()"), AutoComplete("off"),
				FormLabel(Text("Username:")),
				FormInput(Type("text"), Attr("x-model", "username")),
				FormLabel(Text("Password:")),
				FormInput(Type("password"), Attr("x-model", "password")),
				Br(),
				ButtonGray(Type("submit"), Text("Generate Token")),
				Div(Class("mt-7 text-red-500"), Attr("x-text", "error")),
			),
			Hr(),
			Form(Attr("x-on:submit.prevent", "sendRequest()"), AutoComplete("off"),
				FormLabel(Text("Local route:")),
				FormSelect(Attr("x-model", "route"),
					Option(Value("/api/test"), Text("/api/test"), Selected()),
					Option(Value("/api/account"), Text("/api/account - Authorized")),
				),
				Br(),
				FormLabel(Text("Bearer token:")),
				FormInput(Attr("x-model", "token")),
				Br(),
				ButtonGray(Type("submit"), Text("Submit Request")),
			),
			Br(),
			FormLabel(Text("Result:")),
			FormTextarea(Attr("x-text", "response")),
		),
	)
}
