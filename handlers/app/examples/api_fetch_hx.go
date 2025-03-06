package examples

import (
	"encoding/json"
	"io"
	"net/http"
	"time"

	. "maragu.dev/gomponents"
	. "maragu.dev/gomponents/html"
	. "previous/ui"
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

	cols := []string{"Person", "Spacecraft"}

	func() Node {
		return Group{
			If(errorMsg != "", P(InlineStyle("$me { font-size: var(--text-sm); color: $color(red-500); }"), Text(errorMsg))),
			P(InlineStyle("$me { font-size: var(--text-sm); }"), Text("Fetch took: "), ToText(elapsed)),
			P(InlineStyle("$me { margin: $3 $0; }"), Text("Message:")),
			Code(InlineStyle("$me { color: $color(pink-600);}"), ToText(jsonOutput.Message)),

			Br(),
			Br(),

			AutoTableLite(
				cols,
				jsonOutput.People,
				Map(jsonOutput.People, func(p person) Node {
					return Tr(
						Td(Text(p.Name)),
						Td(Text(p.Craft)),
					)
				}),
				AutoTableOptions{
					Compact: true,
					BorderX: true,
					BorderY: true,
					Hover: true,
					Alternate: true,
				},
			),
		}
	}().Render(w)
}
