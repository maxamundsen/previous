package app

import (
	"previous/.metagen/pageinfo"
	"previous/auth"
	. "previous/components"
	. "previous/basic"

	. "maragu.dev/gomponents"
	. "maragu.dev/gomponents/html"
)

func AppLayout(title string, identity auth.Identity, children ...Node) Node {
	navbarDropdown := func (name string, items []StrPair) Node {
		return Div(
			InlineStyle("$this{cursor: pointer; position: relative; margin-left: $(3);}"),
			Attr("x-data", "{" + name + "DropdownOpen: false}"),
			Attr("x-on:click.outside", name+"DropdownOpen = false"),
			Div(
				InlineStyle(`$this { cursor: pointer; display: flex; position: relative; padding-top: $(2); padding-bottom: $(2); padding-left: $(3); font-size: var(--text-sm); line-height: var(--text-sm--line-height); cursor: pointer; font-weight: var(--font-weight-medium); color: var(--color-neutral-100); }`),
				InlineStyle("$this:hover{color: var(--color-white);}"),
				Attr("x-on:click", name+"DropdownOpen = !"+name+"DropdownOpen;"),
				Button(
					Div(InlineStyle("$this{cursor: pointer; display: flex; align-items: center;}"),
						Span(Text(name + " ")),
						Icon(ICON_CHEVRON_DOWN, 16),
					),
				),
			),
			Div(
				InlineStyle(`$this{position: absolute; right: 0; z-index: 10; padding-top: $(1); padding-bottom: $(1); margin-top: $(2); width: $(48); background-color: var(--color-white); transform-origin: top right; box-shadow: var(--shadow-lg);}`),
				Attr("x-cloak"),
				Attr("x-show", name + "DropdownOpen"),
				TabIndex("-1"),
				Map(items, func(item StrPair) Node {
					return A(InlineStyle(`$this{display: block; padding-top: $(2); padding-bottom: $(2); padding-left: $(4); padding-right: $(4); font-size: 0.875rem; line-height: 1.25rem; color: var(--color-neutral-700); } $this:hover{background: var(--color-neutral-100);}`), Href(item.Value), TabIndex("-1"), Text(item.Key))
				}),
			),
		)
	}

	navbarLink := func(name string, url string, newPage bool) Node {
		return A(
			InlineStyle(`$this{ padding-left: $(3); padding-right: $(3); padding-top: $(2); padding-bottom: $(2); font-size: var(--text-sm); font-weight: var(--font-weight-medium); color: var(--color-neutral-100);}`),
			InlineStyle("$this:hover{color: var(--color-white);}"),
			Href(url),
			Text(name),
			If(newPage, Target("_blank")),
		)
	}

	return RootLayout(title+" | Previous",
		Body(InlineStyle("$this{background-color: var(--color-neutral-50); height: 100%;}"),
			Div(InlineStyle("$this{min-height: 100%}"), Attr("x-data", "{ profileDropdownOpen: false, mobileMenuOpen: false }"),
				Nav(InlineStyle("$this{background-color: var(--color-neutral-800);}"),
					Div(InlineStyle("$this{margin-left: auto; margin-right: auto; max-width: var(--container-7xl);}"),
						Div(InlineStyle("$this{display: flex; height: $(16); align-items: center; justify-content: space-between;}"),
							Div(InlineStyle("$this{align-items: center; display: flex;}"),
								Div(InlineStyle("$this{flex-shrink: 0;}"),
									A(Href(pageinfo.Root.Index.Url()), Img(InlineStyle("$this{height: $(12); width: $(12);}"), Src("/images/logo.svg"), Alt("Previous"))),
								),
								Div(InlineStyle("@media $lg-{ $this{display: block;}}"),
									Div(InlineStyle(`$this{margin-left: $(10); display: flex; align-items: baseline;} $this:not(:last-child){ margin-left: $(4); }`),
										navbarLink("Dashboard", pageinfo.Root.App.Dashboard.Url(), false),
										navbarLink("Sitemap", pageinfo.Root.App.Sitemap.Url(), false),
										navbarDropdown(
											"Examples",
											[]StrPair{
												StrPair{Key: "Example Index Page", Value: pageinfo.Root.App.Examples.Index.Url()},
												StrPair{Key: "Auto Table", Value: pageinfo.Root.App.Examples.Autotable.Url()},
												StrPair{Key: "Form Submission", Value: pageinfo.Root.App.Examples.Forms.Url()},
												StrPair{Key: "HTMX", Value: pageinfo.Root.App.Examples.Htmx.Url()},
												StrPair{Key: "Alpine.js", Value: pageinfo.Root.App.Examples.Alpine.Url()},
												StrPair{Key: "UI Elements", Value: pageinfo.Root.App.Examples.Ui_playground.Url()},
												StrPair{Key: "File Uploading", Value: pageinfo.Root.App.Examples.Upload.Url()},
												StrPair{Key: "SMTP Client", Value: pageinfo.Root.App.Examples.Smtp.Url()},
												StrPair{Key: "HTML Sanitization", Value: pageinfo.Root.App.Examples.Html_sanitization.Url()},
												StrPair{Key: "Markdown Rendering", Value: pageinfo.Root.App.Examples.Markdown.Url()},
												StrPair{Key: "Server-side API Fetch", Value: pageinfo.Root.App.Examples.Api_fetch.Url()},
												StrPair{Key: "Inline Styles", Value: pageinfo.Root.App.Examples.Inline_styles.Url()},
												StrPair{Key: "Static Pages", Value: pageinfo.Root.App.Examples.Static.Url()},
											},
										),

										navbarDropdown(
											"API",
											[]StrPair{
												StrPair{Key: "Test", Value: pageinfo.Root.Api.Test.Url()},
												StrPair{Key: "Account", Value: pageinfo.Root.Api.Account.Url()},
												StrPair{Key: "Static", Value: pageinfo.Root.Api.Static.Url()},
											},
										),

										navbarLink("Documentation", "https://github.com/maxamundsen/Previous/wiki", true),
									),
								),
							),
							Div(InlineStyle("$this{display: none;} @media $md { $this{ display: block; }}"),
								Div(InlineStyle("$this{ margin-left: $(4); display: flex; align-items: center;} @media $md { $this{ margin-left: $(6) }}"),
									Div(InlineStyle("$this{position: relative; margin-left: $(3)}"),
										Div(
											Button(
												InlineStyle("$this{cursor: pointer; position: relative; display: flex; max-width: var(--container-xs); background-color: var(--color-neutral-800); font-size: var(--text-sm);}"),
												Attr("x-on:click", "profileDropdownOpen = !profileDropdownOpen"),
												Attr("x-on:click.outside", "profileDropdownOpen = false"),
												Type("button"),
												Span(InlineStyle("$this{position: absolute;}"),
													Img(InlineStyle("$this{height: $(8); width: $(8); border-radius: var(--radius-full)}"), Src("/images/profile_picture.png"), Alt("profile picture")),
												),
											),
										),
										Div(
											InlineStyle("$this{position: absolute; right: 0; z-index: 10; margin-top: $(2); width: $(48); transform-origin: top right; box-shadow: var(--shadow-lg); background: var(--color-white); padding: $(1);}"),
											Attr("x-cloak"),
											Attr("x-show", "profileDropdownOpen"),
											TabIndex("-1"),
											A(Href(pageinfo.Root.App.Account.Url()), InlineStyle("$this{display: block; padding-left: $(4); padding-right: $(4); padding-top: $(2); padding-bottom: $(2); color: var(--color-neutral-700);} $this:hover{background: var(--color-neutral-100);}"), TabIndex("-1"),Text("Your Profile")),
											A(Href(pageinfo.Root.Auth.Logout.Url()), InlineStyle("$this{display: block; padding-left: $(4); padding-right: $(4); padding-top: $(2); padding-bottom: $(2); color: var(--color-neutral-700);} $this:hover{background: var(--color-neutral-100);}"), TabIndex("-1"),Text("Log out")),
										),
									),
								),
							),
							Div(InlineStyle("$this{margin-right: $(2); display: flex;} @media $md{ $this{display: none;}}"),
								Button(
									InlineStyle("$this{position: relative; display: inline-flex; justify-items: center; padding: $(2); color: var(--color-neutral-400)}"),
									InlineStyle("$this:hover{color: var(--color-white); background-color var(--color-neutral-900);}"),
									Attr("x-on:click", "mobileMenuOpen = !mobileMenuOpen"),
									Type("button"),
									Span(InlineStyle("$this{position: absolute;}")),
									Icon(ICON_MENU, 24),
								),
							),
						),
					),
					Div(Attr("x-show", "mobileMenuOpen"), Attr("x-on:click.outside", "mobileMenuOpen = false"), Attr("x-cloak"), InlineStyle("@media $md { $this {display: none; }}"),
						Div(Class("space-y-1 px-2 pb-3 pt-2 sm:px-3"),
							A(Href(pageinfo.Root.App.Dashboard.Url()), Class("block hover:bg-neutral-900 px-3 py-2 text-base font-medium text-white"), Text("Dashboard")),
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
								A(Href(pageinfo.Root.Api.Account.Url()), Class("block px-3 py-2 text-base font-medium text-neutral-200 hover:bg-neutral-900 hover:text-white"), Text("Your Profile")),
								A(Href(pageinfo.Root.Auth.Logout.Url()), Class("block px-3 py-2 text-base font-medium text-neutral-200 hover:bg-neutral-900 hover:text-white"), Text("Log out")),
							),
						),
					),
				),
				Header(InlineStyle("$this{background-color: var(--color-white); box-shadow: var(--shadow-md);}"),
					Div(InlineStyle("$this{margin-left: auto; margin-right: auto; max-width: var(--container-7xl); padding: $(4);} @media $lg { $this{ padding-left: $(8); padding-right: $(8);}}"),
						H1(InlineStyle("$this{font-size: var(--text-3xl); font-weight: var(--font-weight-bold); color: var(--color-neutral-950); letter-spacing: var(--tracking-tight);}"), Text(title)),
					),
				),
				Main(
					Div(InlineStyle("$this{margin-left: auto; margin-right: auto; max-width: var(--container-7xl); padding: $(6) $(4);} @media $sm { $this{padding-left: $(6); padding-right: $(6); }} @media $lg { $this{padding-left: $(8); padding-right: $(8);}}"),
						Group(children),
					),
				),
			),
		),
	)
}
