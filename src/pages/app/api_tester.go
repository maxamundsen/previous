package app

import (
	. "maragu.dev/gomponents"
	. "maragu.dev/gomponents/html"
	. "webdawgengine/pages/components"

	"net/http"
)

func ApiTesterController(w http.ResponseWriter, r *http.Request) {
	ApiTesterView().Render(w)
}

func ApiTesterView() Node {
	return AppLayout("API Tester",
		Form(Attr("onsubmit", "GetToken(); return false;"), AutoComplete("off"),
			Label(Class("form-label"), Text("Username:")),
			Input(Type("text"), Class("form-control"), ID("username")),
			Label(Class("form-label"), Text("Password:")),
			Input(Type("password"), Class("form-control"), ID("password")),
			Br(),
			Button(ID("api_submit"), Type("submit"), Class("btn btn-primary"), Text("Generate Token")),
			Div(Class("text-danger"), ID("error")),
		),
		Hr(),
		Form(Attr("onsubmit", "SendRequest(); return false;"), AutoComplete("off"),
			Label(Class("form-label"), Text("Local route:")),
			Select(Class("form-select"), ID("route"),
				Option(Value("/api/test"), Text("/api/test")),
				Option(Value("/api/account"), Text("/api/account - Authorized")),
			),
			Br(),
			Label(Class("form-label"), Text("Bearer token:")),
			Input(Class("form-control"), ID("token")),
			Br(),
			Button(ID("api_submit"), Type("submit"), Class("btn btn-primary"), Text("Submit Request")),
		),
		Br(),
		Label(Class("form-label"), Text("Result:")),
		Textarea(Class("form-control"), ID("result")),

		Script(Raw(
			`
			function GetToken() {
				var username = document.getElementById("username").value;
				var password = document.getElementById("password").value;

				fetch("/api/auth/login", {
					headers : {
						'Content-Type': 'application/json'
					},
					method: 'POST',
					body: JSON.stringify({
						Username: username,
						Password: password,
					})
				})
			   .then(resp => {
			        if (!resp.ok) {
			            throw new Error(resp.status);
			        }
			        return resp.text();
			   })
			   .then(text => ShowToken(text))
			   .catch( error => ShowError(error))
			}

			function ShowToken(string) {
				ClearError();
				var token = document.getElementById("token");
				token.value = string;
			}

			function ClearError() {
				var errorDiv = document.getElementById("error");
				error.innerHTML = "";
			}

			function ShowError(string) {
				var errorDiv = document.getElementById("error");
				error.innerHTML = string;
			}

			function SendRequest() {
				var route = document.getElementById("route").value;
				var token = document.getElementById("token").value;

				fetch(route, {
				  headers: {Authorization: 'Bearer ' + token}
				})
			   .then(resp => {
			        if (!resp.ok) {
			            throw new Error(resp.status);
			        }
			        return resp.json();
			   })
			   .then(json => ShowResult(JSON.stringify(json)))
			   .catch( error => ShowResult(error))
			}

			function ShowResult(jsonString) {
				var resultbox = document.getElementById("result");
				resultbox.innerHTML = jsonString;
			}
			`),
		),
	)
}
