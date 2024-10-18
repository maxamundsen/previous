package examples

import (
	. "maragu.dev/gomponents"
	. "maragu.dev/gomponents/html"
	. "webdawgengine/pages/components"

	"encoding/json"
	"io"
	"net/http"
	"time"
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

func ApiFetchController(w http.ResponseWriter, r *http.Request) {
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

	ApiFetchView(errorMsg, jsonOutput).Render(w)
}

func ApiFetchView(errorMsg string, model astroModel) Node {
	return AppLayout("API Fetch Example",
		A(Href("http://api.open-notify.org/astros.json"), Text("http://api.open-notify.org/astros.json")),
		Br(),
		P(Text("Message:")),
		Code(ToText(model.Message)),
		Br(),
		Br(),
		Table(Class("table"),
			THead(
				Tr(
					Th(Text("Person")),
					Th(Text("Spacecraft")),
				),
			),
			TBody(
				Map(model.People, func(p person) Node {
					return Tr(
						Td(Text(p.Name)),
						Td(Text(p.Craft)),
					)
				}),
			),
		),
	)
}
