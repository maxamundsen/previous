package app

import (
	"previous/auth"
	. "previous/basic"
	. "previous/components"

	. "maragu.dev/gomponents"
	. "maragu.dev/gomponents/html"
)

func AppLayout(title string, identity auth.Identity, children ...Node) Node {
	navbarDropdown := func(name string, items []StrPair) Node {
		return Div(
			InlineStyle("me{cursor: pointer; position: relative; margin-left: $(3);}"),
			Div(
				InlineStyle(`me{ cursor: pointer; display: flex; position: relative; padding-top: $(2); padding-bottom: $(2); padding-left: $(3); font-size: var(--text-sm); line-height: var(--text-sm--line-height); cursor: pointer; font-weight: var(--font-weight-medium); color: var(--color-neutral-100); }`),
				InlineStyle("me:hover{color: var(--color-white);}"),
				Button(
					InlineScript(`

					`),
					Div(InlineStyle("me{cursor: pointer; display: flex; align-items: center;}"),
						Span(Text(name+" ")),
						Icon(ICON_CHEVRON_DOWN, 16),
					),
				),
			),
			Div(
				InlineStyle(`me{position: absolute; right: 0; z-index: 10; padding-top: $(1); padding-bottom: $(1); margin-top: $(2); width: $(48); background-color: var(--color-white); transform-origin: top right; box-shadow: var(--shadow-lg);}`),
				TabIndex("-1"),
				Map(items, func(item StrPair) Node {
					return A(InlineStyle(`me{display: block; padding-top: $(2); padding-bottom: $(2); padding-left: $(4); padding-right: $(4); font-size: var(--text-sm); line-height: $(5); color: var(--color-neutral-700); } me:hover{background: var(--color-neutral-100);}`), Href(item.Value), TabIndex("-1"), Text(item.Key))
				}),
			),
		)
	}

	navbarLink := func(name string, url string, newPage bool) Node {
		return A(
			InlineStyle(`me{ padding-left: $(3); padding-right: $(3); padding-top: $(2); padding-bottom: $(2); font-size: var(--text-sm); font-weight: var(--font-weight-medium); color: var(--color-neutral-100);}`),
			InlineStyle("me:hover{color: var(--color-white);}"),
			Href(url),
			Text(name),
			If(newPage, Target("_blank")),
		)
	}

	return RootLayout(title+" | Previous",
		Body(InlineStyle("me{background-color: var(--color-neutral-50); height: 100%;}"),
			Div(InlineStyle("me{min-height: 100%}"),
				Nav(InlineStyle("me{background-color: var(--color-neutral-800);}"),
					Div(InlineStyle("me{margin-left: auto; margin-right: auto; max-width: var(--container-7xl);}"),
						Div(InlineStyle("me{display: flex; height: $(16); align-items: center; justify-content: space-between;}"),
							Div(InlineStyle("me{align-items: center; display: flex;}"),
								Div(InlineStyle("me{flex-shrink: 0;}"),
									A(Href("/"), Img(InlineStyle("me{height: $(12); width: $(12);}"), Src("/images/logo.svg"), Alt("Previous"))),
								),
								Div(InlineStyle("@media lg-{ me{display: block;}}"),
									Div(InlineStyle(`me{margin-left: $(10); display: flex; align-items: baseline;} me:not(:last-child){ margin-left: $(4); }`),
										navbarLink("Dashboard", "/app/dashboard", false),
										navbarDropdown(
											"Examples",
											[]StrPair{
												{Key: "Auto Table", Value: "/app/examples/autotable"},
												{Key: "Chart.js", Value: "/app/examples/charts"},
												{Key: "Form Submission", Value: "/app/examples/forms"},
												{Key: "HTMX", Value: "/app/examples/htmx"},
												{Key: "Surreal.js", Value: "/app/examples/surreal"},
												{Key: "UI Elements", Value: "/app/examples/ui-playground"},
												{Key: "File Uploading", Value: "/app/examples/upload"},
												{Key: "SMTP Client", Value: "/app/examples/smtp"},
												{Key: "HTML Sanitization", Value: "/app/examples/html-sanitization"},
												{Key: "Markdown Rendering", Value: "/app/examples/markdown"},
												{Key: "Server-side API Fetch", Value: "/app/examples/api-fetch"},
												{Key: "Inline Styles", Value: "/app/examples/inline-styles"},
											},
										),

										navbarDropdown(
											"API",
											[]StrPair{
												{Key: "Test", Value: "/api/test"},
												{Key: "Account", Value: "/api/account"},
											},
										),

										navbarLink("Documentation", "https://github.com/maxamundsen/Previous/wiki", true),
									),
								),
							),
							Div(InlineStyle("me{display: none;} @media md { me{ display: block; }}"),
								Div(InlineStyle("me{ margin-left: $(4); display: flex; align-items: center;} @media md { me{ margin-left: $(6) }}"),
									Div(InlineStyle("me{position: relative; margin-left: $(3)}"),
										Div(
											Button(
												InlineStyle("me{cursor: pointer; position: relative; display: flex; max-width: var(--container-xs); background-color: var(--color-neutral-800); font-size: var(--text-sm);}"),
												Type("button"),
												Span(InlineStyle("me{position: absolute;}"),
													Img(InlineStyle("me{height: $(8); width: $(8); border-radius: var(--radius-full)}"), Src("/images/profile_picture.png"), Alt("profile picture")),
												),
											),
										),
										Div(
											InlineStyle("me{position: absolute; right: 0; z-index: 10; margin-top: $(2); width: $(48); transform-origin: top right; box-shadow: var(--shadow-lg); background: var(--color-white); padding: $(1);}"),
											TabIndex("-1"),
											A(Href("/app/account"), InlineStyle("me{display: block; padding-left: $(4); padding-right: $(4); padding-top: $(2); padding-bottom: $(2); color: var(--color-neutral-700);} me:hover{background: var(--color-neutral-100);}"), TabIndex("-1"), Text("Your Profile")),
											A(Href("/auth/logout"), InlineStyle("me{display: block; padding-left: $(4); padding-right: $(4); padding-top: $(2); padding-bottom: $(2); color: var(--color-neutral-700);} me:hover{background: var(--color-neutral-100);}"), TabIndex("-1"), Text("Log out")),
										),
									),
								),
							),
							Div(InlineStyle("me{margin-right: $(2); display: flex;} @media md{ me{display: none;}}"),
								Button(
									InlineStyle("me{position: relative; display: inline-flex; justify-items: center; padding: $(2); color: var(--color-neutral-400)}"),
									InlineStyle("me:hover{color: var(--color-white); background-color: var(--color-neutral-900);}"),
									Type("button"),
									Span(InlineStyle("me{position: absolute;}")),
									Icon(ICON_MENU, 24),
								),
							),
						),
					),
					Div(InlineStyle("@media md { me {display: none; }}"),
						Div(Class("space-y-1 px-2 pb-3 pt-2 sm:px-3"),
							A(Href("/app/dashboard"), Class("block hover:bg-neutral-900 px-3 py-2 text-base font-medium text-white"), Text("Dashboard")),
						),
						Div(Class("border-t border-neutral-700 pb-3 pt-4"),
							Div(Class("flex items-center px-5"),
								Div(Class("flex-shrink-0"),
									Img(Class("h-10 w-10 rounded-full"), Src("/images/profile_picture.png"), Alt("profile picture")),
								),
								Div(Class("ml-3"),
									Div(Class("text-base/5 font-medium text-white"), Text(identity.User.Firstname+" "+identity.User.Lastname)),
									Div(Class("text-sm font-medium text-neutral-400"), Text(identity.User.Email)),
								),
							),
							Div(Class("mt-3 space-y-1 px-2"),
								A(Href("/auth/logout"), Class("block px-3 py-2 text-base font-medium text-neutral-200 hover:bg-neutral-900 hover:text-white"), Text("Log out")),
							),
						),
					),
				),
				Header(InlineStyle("me{background-color: var(--color-white); box-shadow: var(--shadow-md);}"),
					Div(InlineStyle("me{margin-left: auto; margin-right: auto; max-width: var(--container-7xl); padding: $(4);} @media lg { me{ padding-left: $(8); padding-right: $(8);}}"),
						H1(InlineStyle("me{font-size: var(--text-3xl); font-weight: var(--font-weight-bold); color: var(--color-neutral-950); letter-spacing: var(--tracking-tight);}"), Text(title)),
					),
				),
				Main(
					Div(InlineStyle("me{margin-left: auto; margin-right: auto; max-width: var(--container-7xl); padding: $(6) $(4);} @media sm { me{padding-left: $(6); padding-right: $(6); }} @media lg { me{padding-left: $(8); padding-right: $(8);}}"),
						Group(children),
					),
				),
			),
		),
	)
}
