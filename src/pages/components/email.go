package components

import (
	. "maragu.dev/gomponents"
	. "maragu.dev/gomponents/html"
)

func ExampleEmailComponent(body string) Node {
	return EmailRoot(
		H1(Text("This email is automatically generated.")),
		P(Text(body)),
	)
}
