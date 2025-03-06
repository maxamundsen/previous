package examples

import (
	"net/http"
	"previous/database"
	"previous/orders"

	. "maragu.dev/gomponents/html"

	. "previous/ui"
)

// This version of autotable does not talk to a database, but uses the same API.
// By simply implementing a filter function that operates on an array rather than a database,
// you can use all the features autotable provides.
func AutoTableInMemoryHxHandler(w http.ResponseWriter, r *http.Request) {
	filter := database.ParseFilterFromRequest(r)
	filter.Pagination.Enabled = true

	// get items after search, before pagination
	searchItems, _ := orders.FilterCustomers(
		database.NewFilterFromSearch(filter.Search),
	)

	// get items after search, after pagination
	customers, _ := orders.FilterCustomers(filter)

	// generate pagination info
	filter.Pagination.GeneratePagination(len(searchItems), len(customers))

	cols := []database.ColInfo {
		{ DisplayName: "Firstname"},
		{ DbName: "Lastname", DisplayName: "Lastname", Sortable: true },
		{ DisplayName: "Email" },
		{ DisplayName: "Phone #" },
	}

	// Generate HTML
	elId := "customer_table"

	AutoTable(
		elId,
		r.URL.Path,
		cols,
		filter,
		customers,
		AutotableSearchGroup(
			AutotableSearch(
				Placeholder("Search Customer Name..."),
				BindSearch(elId, "name"),
				AutoFocus(),
			),
		),
		AutoTableAutoRowFunc(customers),
		nil,
		AutoTableOptions{
			Compact:   true,
			Shadow:    true,
			Hover:     false,
			Alternate: true,
			BorderX:   true,
			BorderY:   false,
		},
	).Render(w)
}