package examples

import (
	"previous/.jet/model"
	"previous/.jet/table"

	// "github.com/go-jet/jet/v2/sqlite"

	. "previous/components"

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
	filter.OrderBy = "id"
	filter.OrderDescending = true

	orders, _ := repository.OrderRepository{}.Filter(filter)
	filter.Pagination.TotalItems = repository.OrderRepository{}.Count()

	filter.Pagination.ProcessPageNum()

	cols := repository.GetFriendlyNamesFromColumns(table.Order.AllColumns)

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
