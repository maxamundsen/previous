package examples

import (
	. "previous/components"
	. "previous/handlers/app"

	. "maragu.dev/gomponents"
	. "maragu.dev/gomponents/html"

	"previous/middleware"

	"net/http"
)

func InlineScriptingHandler(w http.ResponseWriter, r *http.Request) {
	identity := middleware.GetIdentity(r)

	AppLayout("Inline Scripting Example", *identity,
		InlineScript(`
			// initialize a script global variable (this is usually discouraged since components are supposed to be
			// "reusable", however this component is a page, so it will only ever render once per request

			var count = 0;
			var text = "";
		`),

		Prose(
			InlineStyle(`
				$me {
					padding: $(5);
					background: $color(white);
					border: 1px solid $color(neutral-200);
					box-shadow: var(--shadow-md);
					margin-bottom: $(5);
				}
			`),
			Markdown(`
Contrary to popular belief, you can achieve extremely powerful "interactivity", using built-in JavaScript facilities.
Vanilla JS is typically separated from the content, and you need to pass around IDs and class names in order to reference elements on the page.

However, we can use a few shorthand functions (thanks [surreal.js](https://github.com/gnat/surreal)) that wrap around vanilla functionality to localize scripts to our content:
			`),
		),
		ButtonGray(
			Text("Count: 0"),
			InlineScript(`
				let btn = me();
				btn.on("click", () => { count++; btn.innerHTML = "Count: " + count;});
			`),
		),
		Div(
			InlineStyle("$me { margin-top: $(5);}"),
			FormInput(Placeholder("Type stuff here...")),
			P(InlineStyle("$me { color: red; }")),
			InlineScript(`
				let input = me("input", me());
				let par = me("p", me());

				input.on("keyup", () => {
					console.log("You typed text!");
					text = input.value;
					par.innerHTML = text;
				})
			`),
		),
	).Render(w)
}
