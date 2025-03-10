package examples

import (
	. "previous/ui"

	. "maragu.dev/gomponents"
	. "maragu.dev/gomponents/html"

	"previous/database"
	"previous/orders"

	"net/http"
)

func AutoTableHxHandler(w http.ResponseWriter, r *http.Request) {
	filter := database.ParseFilterFromRequest(r)
	filter.Pagination.Enabled = true

	// fetch entities from filter function
	// this first counts the possible items before pagination
	searchItems, _ := orders.Filter(database.NewFilterFromSearch(filter.Search))

	// this query gets the data AFTER pagination
	orderList, _ := orders.Filter(filter)

	// generate page numbers according to total length of data
	filter.Pagination.GeneratePagination(len(searchItems), len(orderList))

	// Declare columns
	cols := []database.ColInfo{
		{DbName: "id", DisplayName: "ID"},
		{DbName: "product_id", DisplayName: "Product ID"},
		{DbName: "purchaser_name", DisplayName: "Customer", Sortable: true},
		{DbName: "purchaser_email", DisplayName: "Customer Email", Sortable: true},
		{DbName: "price", DisplayName: "Price (USD)", Sortable: true, DisplayPosition: database.COL_POS_RIGHT},
	}

	// Generate HTML
	elId := "order_table"
	AutoTable(
		elId,
		r.URL.Path,
		cols,
		filter,
		orderList,
		AutotableSearchGroup(
			AutotableSearch(
				Placeholder("Search Customer Name..."),
				BindSearch(elId, "purchaser_name"),
				AutoFocus(),
			),
			AutotableSearch(
				Placeholder("Search Customer Email..."),
				BindSearch(elId, "purchaser_email"),
			),
			AutotableSearch(
				Placeholder("Price Min"),
				BindSearch(elId, "price_left"),
			),
			AutotableSearch(
				Placeholder("Price Max"),
				BindSearch(elId, "price_right"),
			),
		),
		Map(orderList, func(order orders.Order) Node {
			return Tr(
				TdLeft(B(ToText(order.ID))),
				TdLeft(ToText(order.ProductID)),
				TdLeft(ToText(order.PurchaserName)),
				TdLeft(ToText(order.PurchaserEmail)),
				TdMoney(int64(order.Price)),
			)
		}),
		nil,
		AutoTableOptions{
			Compact:   false,
			Shadow:    true,
			Hover:     false,
			Alternate: false,
			BorderX:   true,
			BorderY:   false,
		},
	).Render(w)
}
