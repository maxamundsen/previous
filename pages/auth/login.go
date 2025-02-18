package auth

import (
	. "previous/components"

	. "maragu.dev/gomponents"
	. "maragu.dev/gomponents/html"

	"previous/auth"
	"previous/config"
	"previous/middleware"

	"log"
	"net/http"
	"strconv"
	"time"
)

// @Identity
// @Protected
// @CookieSession
func LoginPage(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		LoginView("").Render(w)
	} else if r.Method == http.MethodPost {
		username := r.FormValue("username")
		password := r.FormValue("password")
		rememberMe, _ := strconv.ParseBool(r.FormValue("rememberMe"))

		userid, authResult := auth.Authenticate(username, password)

		if !authResult {
			log.Println("Failed login attempt. Username: " + username)
			LoginView("Username or password incorrect.").Render(w)
			return
		}

		log.Println("Successful login. Username: " + username)

		// build identity info
		identity := auth.NewIdentity(userid, rememberMe)

		// serialize and send as cookie
		middleware.PutIdentityCookie(w, r, identity)

		params := r.URL.Query()
		location := params.Get("redirect")

		if len(params["redirect"]) > 0 {
			http.Redirect(w, r, location, http.StatusFound)
			return
		}

		defaultPath := config.IDENTITY_DEFAULT_PATH
		http.Redirect(w, r, defaultPath, http.StatusFound)
	}
}

func LoginView(errorMsg string) Node {
	currentYear := time.Time.Year(time.Now())

	return RootLayout("Login",
		Body(Attr("x-data", "{ clickedLogin: false }"), Class("h-full bg-neutral-50"),
			Div(Class("flex flex-col justify-normal px-6 py-5 lg:px-8 pt-24"),
				Div(Class("sm:mx-auto sm:w-full sm:max-w-sm mb-3"),
					A(Href("/"), Img(Class("mx-auto h-32 w-auto"), Src("/images/logo.svg"), Alt("Previous"))),
				),
				Div(Class("sm:mx-auto sm:w-full sm:max-w-sm"),
					If(errorMsg != "",
						Div(Class("mt-5 sm:mx-auto sm:w-full sm:max-w-sm"),
							P(Class("text-sm/6 text-red-500"), Text(errorMsg)),
						),
					),
					Form(Attr("x-on:submit", "clickedLogin = true;"), Class("mt-5 space-y-6"), Action(""), Method("POST"), AutoComplete("off"),
						H2(Class("mt-10 font-bold text-2xl/9 tracking-tight text-neutral-950"), Text("Log In")),
						Div(
							Label(For("username"), Class("block text-sm/6 font-medium text-neutral-900"), Text("Username")),
							Div(Class("mt-2"),
								Input(Placeholder("admin"), Name("username"), Type("text"), Required(), Class("block w-full border-0 p-1.5 text-neutral-900 shadow-sm ring-1 ring-inset ring-neutral-300 placeholder:text-neutral-400 sm:text-sm/6")),
							),
						),
						Div(
							Label(For("password"), Class("block text-sm/6 font-medium text-neutral-900"), Text("Password")),
							Div(Class("mt-2"),
								Input(Placeholder("admin"), Name("password"), Type("password"), Required(), Class("block w-full border-0 p-1.5 text-neutral-900 shadow-sm ring-1 ring-inset ring-neutral-300 placeholder:text-neutral-400 sm:text-sm/6")),
							),
						),
						Div(
							Button(Attr("x-text", `clickedLogin ? "Authenticating..." : $el.innerText`), Attr("x-bind:disabled", "clickedLogin"), Type("submit"), Class("flex w-full justify-center bg-neutral-950 px-3 py-1.5 text-sm/6 font-semibold text-white shadow-sm hover:bg-neutral-900 focus-visible:outline focus-visible:outline-2 focus-visible:outline-offset-2 focus-visible:outline-neutral-600"), Text("Log in")),
						),
					),
					P(Class("mt-10 text-sm text-neutral-500"), Text("© "), ToText(currentYear), Text(" Max Amundsen")),
				),
			),
		),
	)
}
