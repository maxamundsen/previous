package examples

import (
	. "maragu.dev/gomponents"
	. "maragu.dev/gomponents/html"
	. "previous/components"
	. "previous/pages/app"

	"previous/middleware"
	"previous/auth"

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

// @Identity
// @Protected
// @CookieSession
func ApiFetchController(w http.ResponseWriter, r *http.Request) {
	identity := middleware.GetIdentity(r)

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

	ApiFetchView(errorMsg, *identity, jsonOutput).Render(w)
}

func ApiFetchView(errorMsg string, identity auth.Identity, model astroModel) Node {
	return AppLayout("API Fetch Example", identity,
		PageLink("http://api.open-notify.org/astros.json", Text("http://api.open-notify.org/astros.json"), true),
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
	)
}