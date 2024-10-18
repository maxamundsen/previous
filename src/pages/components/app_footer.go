package components

import (
	. "maragu.dev/gomponents"
	. "maragu.dev/gomponents/html"

	"time"
)

func AppFooter() Node {
	currentYear := time.Time.Year(time.Now())

	return Div(Class("container"),
		Footer(Class("text-secondary mt-5"),
			Text("© "),
			ToText(currentYear),
			Text(" Your Company"),
		),
	)
}
