package examples

import (
	. "maragu.dev/gomponents"
	. "maragu.dev/gomponents/html"
	. "webdawgengine/pages/components"

	"webdawgengine/snailmail"

	"bytes"
	"net/http"
)

func SmtpController(w http.ResponseWriter, r *http.Request) {
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
			SmtpView(err.Error(), "").Render(w)
		} else {
			SmtpView("", "Successfully sent mail!").Render(w)
		}
	} else {
		SmtpView("", "").Render(w)
	}
}

func SmtpView(errorMsg string, successMsg string) Node {
	return AppLayout("SMTP Client Example",
		If(errorMsg != "", Div(Class("alert alert-danger"), Text(errorMsg))),
		If(successMsg != "", Div(Class("alert alert-success"), Text(successMsg))),
		Form(Method("post"), AutoComplete("off"),
			Label(Class("form-label"), Text("To:")),
			Input(Class("form-control"), Type("email"), Name("to")),
			Br(),
			Label(Class("form-label"), Text("Subject:")),
			Input(Class("form-control"), Type("text"), Name("subject")),
			Br(),
			Label(Class("form-label"), Text("Body:")),
			Textarea(Class("form-control"), Name("body")),
			Br(),
			Button(Type("submit"), Class("btn btn-primary"), Text("Send mail")),
		),
	)
}
