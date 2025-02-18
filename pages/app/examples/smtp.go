package examples

import (
	. "previous/components"
	. "previous/pages/app"

	. "maragu.dev/gomponents"
	. "maragu.dev/gomponents/html"

	"previous/auth"
	"previous/middleware"
	"previous/snailmail"

	"bytes"
	"net/http"
)

func SmtpPage(w http.ResponseWriter, r *http.Request) {
	identity := middleware.GetIdentity(r)

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
			SmtpView(err.Error(), "", *identity).Render(w)
		} else {
			SmtpView("", "Successfully sent mail!", *identity).Render(w)
		}
	} else {
		SmtpView("", "", *identity).Render(w)
	}
}

func SmtpView(errorMsg string, successMsg string, identity auth.Identity) Node {
	return AppLayout("SMTP Client Example", identity,
		Div(InlineStyle("me{padding: $(10); background: var(--color-white); border: 1px solid var(--color-neutral-200); box-shadow: var(--shadow-md); margin-bottom: $(5);}"),
			P(InlineStyle("me{font-weight: var(--font-weight-bold); color: var(--color-neutral-800);}"), Text("Note:")),
			P(Text("This demo requires you to connect a valid SMTP server. These options are set in the runtime configuration file.")),
		),
		If(errorMsg != "", P(InlineStyle("me{color: var(--color-red-600);}"), Text(errorMsg))),
		If(successMsg != "", P(InlineStyle("me{color: var(--color-green-600);}"), Text(successMsg))),
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
			ButtonGray(Type("submit"), Text("Send mail")),
		),
	)
}
