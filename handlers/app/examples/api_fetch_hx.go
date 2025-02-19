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

// Struct for deserializing json from 3rd party API
// http://api.open-notify.org/astros.json
type astroModel struct {
	Message string   `json:"message"`
	People  []person `json:"people"`
	Number  int      `json:"number"`
}

type person struct {
	Name  string `json:"name"`
	Craft string `json:"craft"`
}

func ApiFetchHxHandler(w http.ResponseWriter, r *http.Request) {
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

	func() Node {
		return Group{
			If(errorMsg != "", P(Class("text-sm text-red-500"), Text(errorMsg))),
			P(Class("text-sm text-blue-500"), Text("Fetch took: "), ToText(elapsed)),
			P(Class("my-5"), Text("Message:")),
			Code(Class("text-pink-600"), ToText(jsonOutput.Message)),
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
					Map(jsonOutput.People, func(p person) Node {
						return Tr(
							TdTW(Text(p.Name)),
							TdTW(Text(p.Craft)),
						)
					}),
				),
			),
		}
	}().Render(w)
}