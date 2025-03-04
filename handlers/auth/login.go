package auth

import (
	. "previous/ui"

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

func LoginHandler(w http.ResponseWriter, r *http.Request) {
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
		Body(
			InlineStyle(`
				$me {
					height: 100%;
					background: $color(neutral-50);
				}
			`),
			Div(
				InlineStyle(`
					$me {
						display: flex;
						flex-direction: column;
						justify-content: normal;
						padding-right: $6;
						padding-left: $6;
						padding-bottom: $5;
						padding-top: $24;
					}

					@media $md {
						$me {
							padding-right: $8;
							padding-left: $8;
						}
					}
				`),
				Div(
					InlineStyle(`
						$me {
							margin-bottom: $3;
						}

						@media $sm {
							$me {
								margin-right: auto;
								margin-left: auto;
								width: 100%;
								max-width: var(--container-sm);
							}
						}
					`),
					A(Href("/"),
						Img(
							InlineStyle("$me { margin-right: auto; margin-left: auto; height: $32; width: auto; }"),
							Src("/images/logo.svg"),
							Alt("Previous"),
						),
					),
				),
				Div(
					InlineStyle(`
						@media $sm {
							$me {
								margin-right: auto;
								margin-left: auto;
								width: 100%;
								max-width: var(--container-sm);
							}
						}
					`),
					If(errorMsg != "",
						Div(
							InlineStyle(`
								$me {
									margin-top: $5;
								}

								@media $sm {
									$me {
										margin-right: auto;
										margin-left: auto;
										width: 100%;
										max-width: var(--container-sm);
									}
								}
							`),
							P(InlineStyle("$me { font-size: var(--text-sm); color: $color(red-500); }"), Text(errorMsg)),
						),
					),
					Form(InlineStyle("$me { margin-top: $5; }"), Action(""), Method("POST"), AutoComplete("off"),
						H2(
							InlineStyle(`
								$me {
									margin-top: $10;
									margin-bottom: $5;
									font-weight: var(--font-weight-bold);
									font-size: var(--text-2xl);
									letter-spacing: var(--tracking-tight);
									color: $color(neutral-950);
								}
							`),
							Text("Log In"),
						),
						Div(
							Label(
								InlineStyle("$me { display: block; font-size: var(--text-xs); font-weight: var(--font-weight-normal); color: $color(neutral-700);}"),
								For("username"),
								Text("Username"),
							),
							Div(InlineStyle("$me { margin-top: $2; }"),
								FormInput(Placeholder("admin"), Name("username"), Type("text"), Required()),
							),
						),
						Div(
							Label(
								InlineStyle("$me { margin-top: $5; display: block; font-size: var(--text-xs); font-weight: var(--font-weight-normal); color: $color(neutral-700);}"),
								For("password"),
								Text("Password"),
							),
							Div(InlineStyle("$me { margin-top: $2; }"),
								FormInput(Placeholder("admin"), Name("password"), Type("password"), Required()),
							),
						),
						Div(
							InlineStyle(`
								$me {
									margin-top: $5;
								}
							`),
							Button(
								InlineStyle(`
									$me {
										cursor: pointer;
										width: 100%;
										padding-top: $2;
										padding-bottom: $2;
										padding-left: $5;
										padding-right: $5;
										color: $color(white);
										background-color: $color(neutral-800);
										border-radius: var(--radius-xs);
										text-align: center;
										font-size: var(--text-sm);
									}

									$me:hover {
										background-color: $color(neutral-950);
									}
								`),
								Type("submit"),
								Text("Log in"),
							),
						),
						InlineScript(`
							let form = me();
							let btn = me("button", me());

							form.on("submit", () => { btn.innerHTML = "Authenticating..."; });
						`),
					),
					P(
						InlineStyle("$me { margin-top: $10; font-size: var(--text-sm); color: $color(neutral-500);}"),
						Text("© "),
						ToText(currentYear),
						Text(" Max Amundsen"),
					),
				),
			),
		),
	)
}
