package examples

import (
	. "maragu.dev/gomponents"
	. "maragu.dev/gomponents/html"
	. "webdawgengine/pages/components"

	"webdawgengine/snailmail"
	"webdawgengine/middleware"
	"webdawgengine/models"

	"bytes"
	"net/http"
)

func SmtpController(w http.ResponseWriter, r *http.Request) {
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

func SmtpView(errorMsg string, successMsg string, identity models.Identity) Node {
	return AppLayout("SMTP Client Example", identity,
		If(errorMsg != "", P(Class("text-red-600"), Text(errorMsg))),
		If(successMsg != "", P(Class("text-red-600"), Text(successMsg))),
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
			ButtonGray(Type("submit"), Class("btn btn-primary"), Text("Send mail")),
		),
	)
}
