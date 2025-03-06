package examples

import (
	. "maragu.dev/gomponents"
	. "maragu.dev/gomponents/html"

	. "previous/ui"

	// "previous/database"

	"previous/middleware"

	"net/http"
)

func AutoTableHandler(w http.ResponseWriter, r *http.Request) {
	identity := middleware.GetIdentity(r)
	session := middleware.GetSession(r)

	exampleCols := []string{"Col1", "Col2", "Col3", "Col4"}

	type exampleData struct {
		Field1 string
		Field2 string
		Field3 string
		Field4 string
	}

	data := []exampleData{
		{Field1: "hello", Field2: "world", Field3: "test", Field4: "1234"},
		{Field1: "hello", Field2: "world", Field3: "test", Field4: "1234"},
		{Field1: "hello", Field2: "world", Field3: "test", Field4: "1234"},
		{Field1: "hello", Field2: "world", Field3: "test", Field4: "1234"},
	}

	AppLayout("Auto Table", LAYOUT_SECTION_EXAMPLES, *identity, session,
		Card(
			P(Text("This codebase provides an API for generating filterable, sortable, and paginated datagrids such as the one shown below. You do not need to write a single line of JavaScript in order for this to work, as the \"interactivity\" is provided by HTMX.")),
			P(Text("Each interaction with an element of this table generates a dynamic SQL query, retrieves the result, and generates new HTML to swap out the table contents.")),
		),
		Br(),

		// Load the table from the designated handler.
		// The reason this is wrapped in hx-load is because a "full" autotable uses HTMX (swapping html content in place)
		// to achieve interactivity.
		HxLoad("/app/examples/autotable-hx"),

		Br(),
		Br(),

		Card(
			P(Text("The following table operates entirely in-memory. The filter function acts on an array, not against a database.")),
		),

		Br(),

		HxLoad("/app/examples/autotable-inmemory-hx"),

		Br(),
		Br(),

		Card(
			P(Text("This is the 'lite' version of the above table. It uses the same API, however it does not require implementing data filtration or pagination. It is intended to be used with fixed-volume data.")),
		),

		Br(),

		// This table can just be injected straight into the HTML since it does not require interactivity
		AutoTableLite(
			exampleCols,
			data,
			AutoTableAutoRowFunc(data),
			AutoTableOptions{
				Hover:   true,
				Shadow:  true,
				BorderY: true,
				BorderX: true,
			},
		),

		Br(),
		Br(),

		Card(
			P(Text("This is the 'easy' version. All you need to do is pass an array of structs, and some simple table settings. It literally cannot get any easier than this.")),
		),

		Br(),

		AutoTableEasy(
			[]exampleData {
				{Field1: "test", Field2: "more testing", Field3: "hellooooo", Field4: "fourth field"},
				{Field1: "test", Field2: "more testing", Field3: "hellooooo", Field4: "fourth field"},
				{Field1: "test", Field2: "more testing", Field3: "hellooooo", Field4: "fourth field"},
				{Field1: "test", Field2: "more testing", Field3: "hellooooo", Field4: "fourth field"},
			},
			AutoTableOptions{
				Shadow: true,
			},
		),

	).Render(w)
}
