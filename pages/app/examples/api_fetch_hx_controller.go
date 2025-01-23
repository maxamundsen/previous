package examples

import (
	"encoding/json"
	"io"
	"net/http"
	"time"

	. "previous/components"
	. "maragu.dev/gomponents"
	. "maragu.dev/gomponents/html"
)

// @Identity
// @Protected
func ApiFetchHxController(w http.ResponseWriter, r *http.Request) {
	start := time.Now()

	errorMsg := ""

	client := http.Client{
		Timeout: time.Second * 5,
	}

	req, reqErr := http.NewRequest(http.MethodGet, "http://api.open-notify.org/astros.json", nil)

	if reqErr != nil {
		errorMsg = reqErr.Error()
	}

	req.Header.Set("User-Agent", "Example-Api")

	res, resErr := client.Do(req)

	if resErr != nil {
		errorMsg = resErr.Error()
	}

	if res.Body != nil {
		defer res.Body.Close()
	}

	body, readErr := io.ReadAll(res.Body)

	if readErr != nil {
		errorMsg = readErr.Error()
	}

	jsonOutput := astroModel{}

	jsonErr := json.Unmarshal(body, &jsonOutput)

	if jsonErr != nil {
		errorMsg = jsonErr.Error()
	}

	end := time.Now()
	elapsed := end.Sub(start)

	ApiFetchHxView(errorMsg, jsonOutput, elapsed).Render(w)
}

func ApiFetchHxView(errorMsg string, model astroModel, duration time.Duration) Node {
	return Group{
		P(Class("text-sm text-blue-500"), Text("Fetch took: "), ToText(duration)),
		P(Class("my-5"), Text("Message:")),
		Code(Class("text-pink-600"), ToText(model.Message)),
		Br(),
		Br(),
		TableTW(
			THead(
				Tr(
					ThTW(Text("Person")),
					ThTW(Text("Spacecraft")),
				),
			),
			TBodyTW(
				Map(model.People, func(p person) Node {
					return Tr(
						TdTW(Text(p.Name)),
						TdTW(Text(p.Craft)),
					)
				}),
			),
		),
	}
}