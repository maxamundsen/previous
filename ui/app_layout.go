package ui

import (
	"previous/auth"

	. "maragu.dev/gomponents"
	. "maragu.dev/gomponents/html"
)

const (
	LAYOUT_SECTION_DASHBOARD = iota
	LAYOUT_SECTION_EXAMPLES  = iota
	LAYOUT_SECTION_ACCOUNT   = iota
	LAYOUT_SECTION_API       = iota
)

type NavGroup struct {
	SectionId int
	Title     string
	URL       string
	SubGroup  []NavGroup
	NewTab    bool
}

var NavGroups = []NavGroup{
	{SectionId: LAYOUT_SECTION_DASHBOARD, Title: "Dashboard", URL: "/app/dashboard", SubGroup: nil},
	{
		Title: "Examples",
		SubGroup: []NavGroup{
			{SectionId: LAYOUT_SECTION_EXAMPLES, Title: "Form Submission", URL: "/app/examples/forms"},
			{SectionId: LAYOUT_SECTION_EXAMPLES, Title: "File Uploading", URL: "/app/examples/upload"},
			{SectionId: LAYOUT_SECTION_EXAMPLES, Title: "Cookie Sessions", URL: "/app/examples/session"},
			{SectionId: LAYOUT_SECTION_EXAMPLES, Title: "Auto Table", URL: "/app/examples/autotable"},
			{SectionId: LAYOUT_SECTION_EXAMPLES, Title: "Inline Styles", URL: "/app/examples/inline-styles"},
			{SectionId: LAYOUT_SECTION_EXAMPLES, Title: "Inline Scripting", URL: "/app/examples/inline-scripting"},
			{SectionId: LAYOUT_SECTION_EXAMPLES, Title: "HTMX", URL: "/app/examples/htmx"},
			{SectionId: LAYOUT_SECTION_EXAMPLES, Title: "UI Elements", URL: "/app/examples/ui-playground"},
			{SectionId: LAYOUT_SECTION_EXAMPLES, Title: "Chart.js", URL: "/app/examples/charts"},
			{SectionId: LAYOUT_SECTION_EXAMPLES, Title: "SMTP Client", URL: "/app/examples/smtp"},
			{SectionId: LAYOUT_SECTION_EXAMPLES, Title: "HTML Sanitization", URL: "/app/examples/html-sanitization"},
			{SectionId: LAYOUT_SECTION_EXAMPLES, Title: "Markdown Rendering", URL: "/app/examples/markdown"},
			{SectionId: LAYOUT_SECTION_EXAMPLES, Title: "WYSIWYG Editor", URL: "/app/examples/quilljs"},
			{SectionId: LAYOUT_SECTION_EXAMPLES, Title: "Server-side API Fetch", URL: "/app/examples/api-fetch"},
		},
	},
	{
		Title: "API",
		SubGroup: []NavGroup{
			{SectionId: LAYOUT_SECTION_API, Title: "Test", URL: "/api/test", NewTab: true},
			{SectionId: LAYOUT_SECTION_API, Title: "Account", URL: "/api/account", NewTab: true},
		},
	},
	{Title: "Documentation", URL: "https://github.com/maxamundsen/Previous/wiki", NewTab: true},
}

func AppLayout(title string, section int, identity auth.Identity, session map[string]interface{}, children ...Node) Node {
	if session["APP_LAYOUT_VERTICAL"] == true {
		return AppLayoutVertical(title, section, identity, session, children...)
	} else {
		return AppLayoutHorizontal(title, section, identity, session, children...)
	}
}

func AppLayoutVertical(title string, section int, identity auth.Identity, session map[string]interface{}, children ...Node) Node {
	// navbarDivider := Hr(
	// 	InlineStyle("$me { color: $color(neutral-600); }"),
	// )

	navLink := func(url string, navTitle string, newPage bool) Node {
		return Li(
			InlineStyle("$me { margin-bottom: $2; }"),
			A(
				Href(url),
				If(newPage, Target("_blank")),
				InlineStyle("$me { display: block; padding: $2 $4; font-size: var(--text-sm); font-weight: var(--font-weight-semibold); color: $color(neutral-200); letter-spacing: var(--tracking-tight); }"),
				IfElse(title != navTitle,
					InlineStyle(`
						$me:hover {
							background: $color(neutral-600);
						}
					`),
					InlineStyle(`
						$me {
							background: $color(neutral-600);
							border-left: 3px solid $color(red-700);
						}
					`),
				),
				Text(navTitle),
			),
		)
	}

	return RootLayout(title+" | Previous",
		Body(InlineStyle("$me{background-color: $color(neutral-50); height: 100%;}"),
			Nav(
				InlineStyle(`
					$me {
						border-top: 3px solid $color(red-700);
						position: fixed;
						top: 0;
						left: 0;
						background: $color(neutral-900);
						height: $16;
						width: 100%;
						z-index: 30;
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
						overflow-y: auto;
						border-right: 1px solid $color(neutral-200);
						box-shadow: var(--shadow-sm);
						background: $color(neutral-800);
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
					Ul(
						InlineStyle("$me:not(:last-child) { padding-top: $1; padding-bottom: $1;} $me { padding-top: $3; }"),
						Map(NavGroups, func(nav NavGroup) Node {
							if len(nav.SubGroup) > 0 {
								return Group{
									// Subsection Title
									Li(
										InlineStyle(`
												$me {
													display: flex;
													align-items: center;
													justify-content: space-between;
													padding: $2 $4;
													color: $color(neutral-200);
													border-bottom: 1px solid $color(neutral-600);
												}
											`),
										Span(InlineStyle("$me { font-size: var(--text-xs); text-transform: uppercase;}"), Text(nav.Title)),
									),
									// Each subsection item
									Ul(InlineStyle("$me:not(:last-child) { margin-top: $1; margin-botton: $2; }"),
										Map(nav.SubGroup, func(subnav NavGroup) Node {
											return navLink(subnav.URL, subnav.Title, subnav.NewTab)
										}),
									),
								}
							} else {
								return Group{
									navLink(nav.URL, nav.Title, nav.NewTab),
								}
							}
						}),
					),
				),
			),

			Div(
				InlineStyle(`
					$me {
						margin: $16 0 0 $5;
						padding: $4 $8;
						background: $color(white);
						box-shadow: var(--shadow-sm);
					}

					@media $sm {
						$me {
							margin-left: $64;
						}
					}
				`),
				H1(
					InlineStyle("$me { letter-spacing: var(--tracking-tight); font-size: var(--text-2xl); color: $color(black); font-weight: var(--font-weight-bold); }"),
					Text(title),
				),
			),

			Main(
				InlineStyle(`
					$me {
						margin: 0 0 $5 $5;
						padding: $4 $8;
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
	navbarDropdown := func(dropdownHeader Node, dropdownItems Node) Node {
		return Div(
			InlineStyle("$me{cursor: pointer; position: relative; margin-left: $3;}"),
			Div(
				Class("button"),
				InlineStyle(`
					$me {
						cursor: pointer;
						display: flex;
						position: relative;
						padding-top: $2;
						padding-bottom: $2;
						padding-left: $3;
						font-size: var(--text-sm);
						line-height: var(--text-sm--line-height);
						font-weight: var(--font-weight-medium);
						color: $color(neutral-300);
					}

					$me:hover{color: $color(white);}
				`),
				Button(
					Div(InlineStyle("$me{cursor: pointer; display: flex; align-items: center;}"),
						dropdownHeader, Span(Text(" ")),
						Icon(ICON_CHEVRON_DOWN, 16),
					),
				),
			),
			Div(
				Class("dropdown"),
				InlineStyle(`$me{display: none; position: absolute; right: 0; z-index: 10; padding-top: $1; padding-bottom: $1; margin-top: $2; width: $48; background-color: $color(white); transform-origin: top right; box-shadow: var(--shadow-lg);}`),
				TabIndex("-1"),
				dropdownItems,
			),
			InlineScript(`
				let button = me(".button", me());
				let dropdown = me(".dropdown", me());

				button.on("click", ev => { toggleShowHide(dropdown) });
				onClickOutsideOrEscape(me(), () => { hide(dropdown) });
			`),
		)
	}

	navbarDropdownItem := func(name string, url string, newPage bool) Node {
		return A(InlineStyle(`$me{display: block; padding-top: $2; padding-bottom: $2; padding-left: $4; padding-right: $4; font-size: var(--text-sm); line-height: $5; color: $color(neutral-700); } $me:hover{background: $color(neutral-100);}`), Href(url), TabIndex("-1"), Text(name), If(newPage, Target("_blank")))
	}

	navbarLink := func(name string, url string, newPage bool) Node {
		return A(
			InlineStyle(`$me{ padding-left: $3; padding-right: $3; padding-top: $2; padding-bottom: $2; font-size: var(--text-sm); font-weight: var(--font-weight-medium); color: $color(neutral-300);}`),
			InlineStyle("$me:hover{color: $color(white);}"),
			Href(url),
			Text(name),
			If(newPage, Target("_blank")),
		)
	}

	return RootLayout(title+" | Previous",
		Body(InlineStyle("$me{background-color: $color(neutral-50); height: 100%;}"),
			Div(InlineStyle("$me{min-height: 100%}"),
				Nav(InlineStyle("$me{background-color: $color(neutral-900); border-top: 3px solid $color(red-700);}"),
					Div(InlineStyle("$me{margin-left: auto; margin-right: auto; max-width: var(--container-7xl);}"),
						Div(InlineStyle("$me{display: flex; height: $16; align-items: center; justify-content: space-between;}"),
							Div(InlineStyle("$me{align-items: center; display: flex;}"),
								Div(InlineStyle("$me{flex-shrink: 0;}"),
									A(Href("/"), Img(InlineStyle("$me{height: $8; width: $8;}"), Src("/images/logo.svg"), Alt("Previous"))),
								),
								Div(InlineStyle("@media $lg-{ $me{display: block;}}"),
									Div(InlineStyle(`$me{margin-left: $1; display: flex; align-items: baseline;} $me:not(:last-child){ margin-left: $4; }`),
										Map(NavGroups, func(nav NavGroup) Node {
											if len(nav.SubGroup) > 0 {
												return navbarDropdown(
													Text(nav.Title),
													Map(nav.SubGroup, func(sub NavGroup) Node {
														return navbarDropdownItem(sub.Title, sub.URL, sub.NewTab)
													}),
												)
											} else {
												return navbarLink(nav.Title, nav.URL, nav.NewTab)
											}
										}),
									),
								),
							),
							Div(InlineStyle("$me{display: none;} @media $md { $me{ display: block; }}"),
								Div(InlineStyle("$me{ margin-left: $4; display: flex; align-items: center;} @media $md { $me{ margin-left: $6;}}"),
									Div(InlineStyle("$me{position: relative; margin-left: $3;}"),
										navbarDropdown(
											Icon(ICON_USERS, 24),
											Group{
												navbarDropdownItem("My Profile", "/app/account", false),
												navbarDropdownItem("Logout", "/auth/logout", false),
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
