package examples

import (
	. "previous/components"

	. "maragu.dev/gomponents"
	. "maragu.dev/gomponents/html"

	"previous/auth"
	"previous/middleware"
	"previous/snailmail"

	"bytes"
	"net/http"
)

func SmtpHandler(w http.ResponseWriter, r *http.Request) {
	identity := middleware.GetIdentity(r)
	session := middleware.GetSession(r)

	var errMsg string
	var successMsg string

	if r.Method == http.MethodPost {
		recipients := r.FormValue("to")
		subject := r.FormValue("subject")
		body := r.FormValue("body")

		var b bytes.Buffer

		ExampleEmailComponent(body).Render(&b)

		message := snailmail.Email{
			Recipients: []string{recipients},
			Subject:    subject,
			Body:       &b,
		}

		err := snailmail.SendMail(message, snailmail.TYPE_HTML)

		if err != nil {
			errMsg = err.Error()
		} else {
			successMsg = "Successfully sent mail!"
		}
	}

	SmtpView(errMsg, successMsg, *identity, session).Render(w)

}

func SmtpView(errorMsg string, successMsg string, identity auth.Identity, session map[string]interface{}) Node {
	return AppLayout("SMTP Client Example", LAYOUT_SECTION_EXAMPLES, identity, session,
		Card(InlineStyle("$me { margin-bottom: $5; }"),
			P(InlineStyle("$me{font-weight: var(--font-weight-bold); color: $color(neutral-800);}"), Text("Note:")),
			P(Text("This demo requires you to connect a valid SMTP server. These options are set in the runtime configuration file.")),
		),
		If(errorMsg != "", P(InlineStyle("$me{color: $color(red-600);}"), Text(errorMsg))),
		If(successMsg != "", P(InlineStyle("$me{color: $color(green-600);}"), Text(successMsg))),
		Form(Method("post"), AutoComplete("off"),
			FormLabel(Text("To:")),
			FormInput(Type("email"), Name("to")),
			Br(),
			FormLabel(Text("Subject:")),
			FormInput(Type("text"), Name("subject")),
			Br(),
			FormLabel(Text("Body:")),
			FormTextarea(Name("body")),
			Br(),
			ButtonUI(Type("submit"), Text("Send mail")),
		),
	)
}
