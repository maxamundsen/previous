package examples

import (
	"previous/.jet/model"
	"previous/.jet/table"

	. "previous/components"

	"github.com/go-jet/jet/v2/sqlite"

	. "maragu.dev/gomponents"
	. "maragu.dev/gomponents/html"

	"previous/repository"

	"net/http"
)

// @Identity
// @Protected
// @CookieSession
func OrdersHxController(w http.ResponseWriter, r *http.Request) {
	filter := repository.ParseFilterFromRequest(r)
	url := r.URL.Path
	OrdersHxView(url, filter).Render(w)
}

func OrdersHxView(url string, filter repository.Filter) Node {
	filter.Pagination.ItemsPerPage = 5

	// fetch entities from filter function
	orders, _ := repository.OrderRepository{}.Filter(filter)
	searchItems, _ := repository.OrderRepository{}.Filter(repository.Filter{Search: filter.Search})
	filter.Pagination.TotalItems = len(searchItems)

	// generate page numbers according to total length of data
	filter.Pagination.GeneratePageNumbers()

	// You can automatically generate friendly names from the SQL columns:
	cols := repository.GetColInfoFromJet(
		sqlite.ColumnList{
			table.Order.ID,
			table.Order.ProductID,
			table.Order.PurchaserName,
			table.Order.PurchaserEmail,
			table.Order.Price,
		},
	)

	// Or you can map them manually:
	cols = []repository.ColInfo{
		{DbName: table.Order.ID.Name(), DisplayName: "ID"},
		{DbName: table.Order.ProductID.Name(), DisplayName: "Product ID"},
		{DbName: table.Order.PurchaserName.Name(), DisplayName: "Customer"},
		{DbName: table.Order.PurchaserEmail.Name(), DisplayName: "Customer Email"},
		{DbName: table.Order.Price.Name(), DisplayName: "Price (USD)"},
	}

	return AutoTable(
		"order_table",
		url,
		filter,
		orders,
		func(order model.Order) Node {
			return Tr(Class("hover:bg-slate-50 border-b border-slate-200"),
				Td(Class("p-4 py-5"),
					P(Class("block font-semibold text-sm text-slate-800"), ToText(order.ID)),
				),
				Td(Class("p-4 py-5"),
					P(Class("block text-sm text-slate-800"), ToText(order.ProductID)),
				),
				Td(Class("p-4 py-5"),
					P(Class("block text-sm text-slate-800"), ToText(order.PurchaserName)),
				),
				Td(Class("p-4 py-5"),
					P(Class("block text-sm text-slate-800"), ToText(order.PurchaserEmail)),
				),
				Td(Class("p-4 py-5"),
					P(Class("block font-semibold text-sm text-slate-800"), Text("$"), FormatMoney(int64(order.Price))),
				),
			)
		},
		cols,
	)
}
