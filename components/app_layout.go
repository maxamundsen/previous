package components

import (
	"previous/auth"

	. "maragu.dev/gomponents"
	. "maragu.dev/gomponents/html"
)

const (
	LAYOUT_SECTION_DASHBOARD = iota
	LAYOUT_SECTION_EXAMPLES = iota
	LAYOUT_SECTION_ACCOUNT = iota
	LAYOUT_SECTION_API = iota
)

func AppLayout(title string, section int, identity auth.Identity, session map[string]interface{}, children ...Node) Node {
	if session["APP_LAYOUT_VERTICAL"] == true {
		return AppLayoutVertical(title, section, identity, session, children...)
	} else {
		return AppLayoutHorizontal(title, section, identity, session, children...)
	}
}

func AppLayoutVertical(title string, section int, identity auth.Identity, session map[string]interface{}, children ...Node) Node {
	return RootLayout(title+" | Previous",
		Body(InlineStyle("$me{background-color: $color(neutral-50); height: 100%;}"),
			Nav(
				InlineStyle(`
					$me {
						position: fixed;
						top: 0;
						left: 0;
						background: $color(neutral-800);
						height: $16;
						width: 100%;
						z-index: 50;
					}
				`),
				Text("test"),
			),

			Button(
				InlineStyle(`
					$me {
						display: inline-flex;
						align-items: center;
						padding: $2;
						margin-top: $2;
						margin-left: $2;
						color: $color(neutral-700);
					}

					$me:hover {
						background: $color(neutral-100);
					}

					@media $sm {
						$me {
							display: none;
						}
					}
				`),
				Type("button"),
				Span(Text("Open sidebar")),
				Icon(ICON_MENU, 24),
			),
			Aside(
				InlineStyle(`
					$me {
						position: fixed;
						border-right: 1px solid $color(neutral-200);
						box-shadow: var(--shadow-sm);
						background: $color(white);
						top: $16;
						left: 0;
						z-index: 40;
						width: $64;
						height: 100vh;
						translate: -100% $2;
					}

					@media $sm {
						$me {
							translate: 1 $2;
						}
					}
				`),
				Div(
					InlineStyle(`$me { padding: $6 $4; overflow-y: auto; }`),
					Ul(
						InlineStyle("$me:not(:last-child) { padding-top: $1; padding-bottom: $1; }"),
						A(
							InlineStyle(`
								$me {
									display: block;
									padding: $2 $4;
									font-size: var(--text-sm);
									color: $color(neutral-700);
								}

								$me:hover {
									color: $color(neutral-950);
									text-decoration: underline;
								}
							`),
							If(section == LAYOUT_SECTION_DASHBOARD,
								InlineStyle(`
									$me {
										color: $color(neutral-950);
										font-weight: var(--font-weight-medium);
									}
								`),
							),
							Href("/app/dashboard"),
							Text("Dashboard"),
						),
						// Map(DocList, func(doc Document) Node {
						// 	if len(doc.SubList) > 0 {
						// 		return Li(
						// 			Details(Class("group [&_summary::-webkit-details-marker]:hidden"), If(doc.DisplayId == displayId, Attr("open")),
						// 				Summary(Class("flex cursor-pointer items-center justify-between px-4 py-2 text-neutral-700 hover:text-neutral-950 hover:underline"),
						// 					Span(Class("text-sm" ), Text(doc.Title)),
						// 					Span(Class("shrink-0 transition duration-300 group-open:-rotate-180"),
						// 						Icon(ICON_CHEVRON_UP, 16),
						// 					),
						// 				),
						// 				Ul(Class("mt-2 space-y-1 px-4"),
						// 					Map(doc.SubList, func(subdoc Document) Node {
						// 						return Li(
						// 							A(Href("/docs/"+subdoc.Slug), Classes{"block px-4 py-2 text-sm text-neutral-700": true, "hover:text-neutral-950 hover:underline": title != subdoc.Title, "text-neutral-950 font-medium": title == subdoc.Title}, Text(subdoc.Title)),
						// 						)
						// 					}),
						// 				),
						// 			),
						// 		)
						// 	} else {
						// 		return Li(
						// 			A(Href("/docs/"+doc.Slug), Classes{"block px-4 py-2 text-sm text-neutral-700 hover:text-neutral-950 hover:underline": true, "text-neutral-950 font-medium": displayId == doc.DisplayId}, Text(doc.Title)),
						// 		)
						// 	}
						// }),
					),
				),
			),
			Main(
				InlineStyle(`
					$me {
						margin: $16 0 $5 $5;
						padding: $10;
					}

					@media $sm {
						$me {
							margin-left: $64;
						}
					}
				`),
				Group(children),
			),
		),
	)
}

func AppLayoutHorizontal(title string, section int, identity auth.Identity, session map[string]interface{}, children ...Node) Node {
	navbarDropdown := func(dropdownHeader Node, items [][2]string) Node {
		return Div(
			InlineStyle("$me{cursor: pointer; position: relative; margin-left: $3;}"),
			Div(
				Class("button"),
				InlineStyle(`$me{ cursor: pointer; display: flex; position: relative; padding-top: $2; padding-bottom: $2; padding-left: $3; font-size: var(--text-sm); line-height: var(--text-sm--line-height); font-weight: var(--font-weight-medium); color: $color(neutral-100); }`),
				InlineStyle("$me:hover{color: $color(white);}"),
				Button(
					Div(InlineStyle("$me{cursor: pointer; display: flex; align-items: center;}"),
						dropdownHeader, Span(Text(" ")),
						Icon(ICON_CHEVRON_DOWN, 16),
					),
				),
			),
			Div(
				Class("dropdown"),
				InlineStyle(`$me{display: none; border-radius: var(--radius-sm); position: absolute; right: 0; z-index: 10; padding-top: $1; padding-bottom: $1; margin-top: $2; width: $48; background-color: $color(white); transform-origin: top right; box-shadow: var(--shadow-lg);}`),
				TabIndex("-1"),
				Map(items, func(item [2]string) Node {
					return A(InlineStyle(`$me{display: block; padding-top: $2; padding-bottom: $2; padding-left: $4; padding-right: $4; font-size: var(--text-sm); line-height: $5; color: $color(neutral-700); } $me:hover{background: $color(neutral-100);}`), Href(item[1]), TabIndex("-1"), Text(item[0]))
				}),
			),
			InlineScript(`
				let button = me(".button", me());
				let dropdown = me(".dropdown", me());

				button.on("click", ev => { toggleShowHide(dropdown) });
				onClickOutsideOrEscape(me(), () => { hide(dropdown) });
			`),
		)
	}

	navbarLink := func(name string, url string, newPage bool) Node {
		return A(
			InlineStyle(`$me{ padding-left: $3; padding-right: $3; padding-top: $2; padding-bottom: $2; font-size: var(--text-sm); font-weight: var(--font-weight-medium); color: $color(neutral-100);}`),
			InlineStyle("$me:hover{color: $color(white);}"),
			Href(url),
			Text(name),
			If(newPage, Target("_blank")),
		)
	}

	return RootLayout(title+" | Previous",
		Body(InlineStyle("$me{background-color: $color(neutral-50); height: 100%;}"),
			Div(InlineStyle("$me{min-height: 100%}"),
				Nav(InlineStyle("$me{background-color: $color(neutral-800);}"),
					Div(InlineStyle("$me{margin-left: auto; margin-right: auto; max-width: var(--container-7xl);}"),
						Div(InlineStyle("$me{display: flex; height: $16; align-items: center; justify-content: space-between;}"),
							Div(InlineStyle("$me{align-items: center; display: flex;}"),
								Div(InlineStyle("$me{flex-shrink: 0;}"),
									A(Href("/"), Img(InlineStyle("$me{height: $8; width: $8;}"), Src("/images/logo.svg"), Alt("Previous"))),
								),
								Div(InlineStyle("@media $lg-{ $me{display: block;}}"),
									Div(InlineStyle(`$me{margin-left: $1; display: flex; align-items: baseline;} $me:not(:last-child){ margin-left: $4; }`),
										navbarLink("Dashboard", "/app/dashboard", false),
										navbarDropdown(
											Text("Examples"),
											[][2]string{
												{"Form Submission", "/app/examples/forms"},
												{"File Uploading", "/app/examples/upload"},
												{"Auto Table", "/app/examples/autotable"},
												{"Inline Styles", "/app/examples/inline-styles"},
												{"Inline Scripting", "/app/examples/inline-scripting"},
												{"HTMX", "/app/examples/htmx"},
												{"UI Elements", "/app/examples/ui-playground"},
												{"Chart.js", "/app/examples/charts"},
												{"SMTP Client", "/app/examples/smtp"},
												{"HTML Sanitization", "/app/examples/html-sanitization"},
												{"Markdown Rendering", "/app/examples/markdown"},
												{"Server-side API Fetch", "/app/examples/api-fetch"},
											},
										),

										navbarDropdown(
											Text("API"),
											[][2]string{
												{"Test", "/api/test"},
												{"Account", "/api/account"},
											},
										),

										navbarLink("Documentation", "https://github.com/maxamundsen/Previous/wiki", true),
									),
								),
							),
							Div(InlineStyle("$me{display: none;} @media $md { $me{ display: block; }}"),
								Div(InlineStyle("$me{ margin-left: $4; display: flex; align-items: center;} @media $md { $me{ margin-left: $6;}}"),
									Div(InlineStyle("$me{position: relative; margin-left: $3;}"),
										navbarDropdown(
											Icon(ICON_USERS, 24),
											[][2]string{
												{"Your Profile", "/app/account"},
												{"Log Out", "/auth/logout"},
											},
										),
									),
								),
							),
							Div(InlineStyle("$me{margin-right: $2; display: flex;} @media $md{ $me{display: none;}}"),
								Button(
									InlineStyle("$me{position: relative; display: inline-flex; justify-items: center; padding: $2; color: $color(neutral-400)}"),
									InlineStyle("$me:hover{color: $color(white); background-color: $color(neutral-900);}"),
									Type("button"),
									Span(InlineStyle("$me{position: absolute;}")),
									Icon(ICON_MENU, 24),
								),
							),
						),
					),
					Div(InlineStyle("@media $md { $me {display: none; }}"),
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
				Header(InlineStyle("$me{background-color: $color(white); box-shadow: var(--shadow-sm);}"),
					Div(InlineStyle("$me{margin-left: auto; margin-right: auto; max-width: var(--container-7xl); padding: $4;} @media $lg { $me{ padding-left: $8; padding-right: $8;}}"),
						H1(InlineStyle("$me{font-size: var(--text-3xl); font-weight: var(--font-weight-bold); color: $color(neutral-950); letter-spacing: var(--tracking-tight);}"), Text(title)),
					),
				),
				Main(
					Div(InlineStyle("$me{margin-left: auto; margin-right: auto; max-width: var(--container-7xl); padding: $6 $4;} @media $sm { $me{padding-left: $6; padding-right: $6; }} @media $lg { $me{padding-left: $8; padding-right: $8;}}"),
						Group(children),
					),
				),
			),
		),
	)
}
