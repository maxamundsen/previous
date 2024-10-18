package auth

import (
	. "maragu.dev/gomponents"
	. "maragu.dev/gomponents/html"
	. "webdawgengine/pages/components"

	"webdawgengine/auth"
	"webdawgengine/config"
	"webdawgengine/middleware"

	"log"
	"net/http"
	"strconv"
	"time"
)

func LoginController(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		LoginView("").Render(w)
	} else if r.Method == http.MethodPost {
		username := r.FormValue("username")
		password := r.FormValue("password")
		rememberMe, _ := strconv.ParseBool(r.FormValue("rememberMe"))

		user, authResult := auth.Authenticate(username, password)

		if !authResult {
			log.Println("Failed login attempt. Username: " + username)
			LoginView("Username or password incorrect.").Render(w)
			return
		}

		log.Println("Successful login. Username: " + username)

		// build identity info
		identity := middleware.NewIdentity(user.Id, user.SecurityStamp, true, rememberMe)

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

	return Root("Login",
		Div(Class("container-sm"), Style("margin-top: 20px;"),
			Div(Style("text-align: center;"),
				A(Href("/"),
					Img(Src("/images/logo.png"), Alt("logo")),
				),
			),
			If(errorMsg != "", Div(Class("alert alert-danger"), Text(errorMsg))),
			Br(),
			Form(Action("/auth/login"), Method("POST"), AutoComplete("off"), ID("loginForm"),
				H1(Class("h3 mb-3 fw-normal"), Text("Log In")),
				Label(For("Username"), Text("Username")),
				Input(Name("username"), Type("text"), Class("form-control"), ID("Username"), AutoFocus(), Required()),
				Br(),
				Label(For("password"), Text("Password")),
				Input(Name("password"), Type("password"), Class("form-control"), ID("password"), Required()),
				Br(),
				Input(Class("form-check-input"), Type("checkbox"), Value("true"), ID("flexCheckDefault"), Name("rememberMe")),
				Label(Class("form-check-label"), For("flexCheckDefault"), Text("Remember me")),
				Br(),
				Br(),
				Button(Class("btn btn-primary w-100 py-2"), ID("loginBtn"), Type("submit"), Text("Sign in")),
				P(Class("mt-5 mb-3 text-body-secondary"), Text("© "), ToText(currentYear), Text(" Example Company")),
			),
		),
		Raw(`<script>
		    // Provide feedback when form is submitted, since it takes a few seconds
		    // for the Bcrypt algo to validate hashes
		    document.getElementById("loginForm").addEventListener("submit", function (event) {
		        var loginBtn = document.getElementById("loginBtn");
		        loginBtn.innerText = "Authenticating...";
		        loginBtn.disabled = true;

		        setTimeout(function () {
		            loginBtn.innerText = "Login";
		            loginBtn.disabled = false;
		        }, 99999999);
		    });
		</script>`),
	)
}
