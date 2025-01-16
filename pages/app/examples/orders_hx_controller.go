package examples

import (
	"previous/.jet/model"

	. "maragu.dev/gomponents"
	. "maragu.dev/gomponents/html"
	. "previous/components"

	"previous/repository"

	"net/http"
	"strconv"
)

// @Identity
// @Protected
// @CookieSession
func OrdersHxController(w http.ResponseWriter, r *http.Request) {
	filter := repository.Filter{}

	filter.Search = r.URL.Query().Get("search")
	filter.OrderBy = r.URL.Query().Get("orderBy")
	filter.OrderDescending, _ = strconv.ParseBool(r.URL.Query().Get("desc"))
	filter.Pagination.CurrentPage, _ = strconv.Atoi(r.URL.Query().Get("pageNum"))
	filter.Pagination.ItemsPerPage, _ = strconv.Atoi(r.URL.Query().Get("itemsPerPage"))

	url := r.URL.Path

	OrdersHxView(url, filter).Render(w)
}

func OrdersHxView(url string, filter repository.Filter) Node {
	filter.Pagination.ItemsPerPage = 20
	filter.OrderBy = "id"
	filter.OrderDescending = true
	filter.Pagination.ProcessPageNum()

	orders, _ := repository.OrderRepository{}.Filter(filter)

	return Group{
		Form(AutoComplete("off"), Attr("hx-get", url), Attr("hx-trigger", "keyup delay:100ms from:input, click from:select"), Attr("hx-swap", "outerHTML"), Attr("hx-target", "#ordersTable"), Attr("hx-select", "#ordersTable"),
			Div(Class("w-full flex justify-between items-center mb-3 mt-1 pl-3"),
				Div(
					H3(Class("text-lg font-semibold text-slate-800"), Text("Example Orders")),
					P(Class("text-slate-500"), Text("An example `orders` table with filtering, pagination, and sorting.")),
				),
				Div(Class("ml-3"),
					Div(Class("w-full max-w-sm min-w-[200px] relative"),
						Div(Class("relative"),
							Input(Class("bg-white w-full pr-11 h-10 pl-3 py-2 bg-transparent placeholder:text-slate-400 text-slate-700 text-sm border border-slate-200 rounded transition duration-200 ease focus:outline-none focus:border-slate-400 hover:border-slate-400 shadow-sm focus:shadow-md"), Placeholder("Search for customer..."), Name("search"), AutoFocus()),
							Button(Class("absolute h-8 w-8 right-1 top-1 my-auto px-2 flex items-center bg-white rounded "), Type("button"),
								Icon("search", 24),
							),
						),
					),
				),
			),
		),
		Div(ID("ordersTable"),
			Div(Class("relative flex flex-col w-full h-full overflow-scroll text-gray-700 bg-white shadow-md rounded-lg bg-clip-border"),
				Table(Class("w-full text-left table-auto min-w-max"),
					THead(
						Tr(
							Th(HxOnClick(url, "outerHTML", "#ordersTable", "#ordersTable"), Class("p-4 border-b border-slate-200 transition-colors cursor-pointer bg-slate-50 hover:bg-slate-100"),
								P(Class("flex items-center justify-between gap-2 font-sans text-sm font-normal leading-none text-slate-500"), Text("ID"),
									Icon("sort-arrows", 16),
								),
							),
							Th(Class("p-4 border-b border-slate-200 transition-colors cursor-pointer bg-slate-50 hover:bg-slate-100"),
								P(Class("flex items-center justify-between gap-2 font-sans text-sm font-normal leading-none text-slate-500"), Text("Product ID"),
									Icon("sort-arrows", 16),
								),
							),
							Th(Class("p-4 border-b border-slate-200 transition-colors cursor-pointer bg-slate-50 hover:bg-slate-100"),
								P(Class("flex items-center justify-between gap-2 font-sans text-sm font-normal leading-none text-slate-500"), Text("Customer Name"),
									Icon("sort-arrows", 16),
								),
							),
							Th(Class("p-4 border-b border-slate-200 transition-colors cursor-pointer bg-slate-50 hover:bg-slate-100"),
								P(Class("flex items-center justify-between gap-2 font-sans text-sm font-normal leading-none text-slate-500"), Text("Customer Email"),
									Icon("sort-arrows", 16),
								),
							),
							Th(Class("p-4 border-b border-slate-200 transition-colors cursor-pointer bg-slate-50 hover:bg-slate-100"),
								P(Class("flex items-center justify-between gap-2 font-sans text-sm font-normal leading-none text-slate-500"), Text("Price"),
									Icon("sort-arrows", 16),
								),
							),
						),
					),
					TBody(
						Map(orders, func(order model.Order) Node {
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
						}),
					),
				),
				Div(Class("flex justify-between items-center px-4 py-3"),
					Div(Class("text-sm text-slate-500"), Text("Showing "), ToText(filter.Pagination.CurrentPage), Text(" of "), ToText(filter.Pagination.TotalPages),
						B(Text(" Total Items: ")), ToText(filter.Pagination.TotalItems),
					),
					Div(Class("flex space-x-1"),
						Button(Class("px-3 py-1 min-w-9 min-h-9 text-sm font-normal text-slate-500 bg-white border border-slate-200 rounded hover:bg-slate-50 hover:border-slate-400 transition duration-200 ease"), Text("Prev")),

						Button(Class("px-3 py-1 min-w-9 min-h-9 text-sm font-normal text-slate-500 bg-white border border-slate-200 rounded hover:bg-slate-50 hover:border-slate-400 transition duration-200 ease"), Text("Next")),
					),
				),
			),
		),
	}
}
