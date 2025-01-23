package components

import (
	. "previous/basic"

	"previous/repository"

	. "maragu.dev/gomponents"
	. "maragu.dev/gomponents/components"
	. "maragu.dev/gomponents/html"
)

func AutoTable[E any](tableId string, url string, f repository.Filter, entities []E, nf func(E) Node, cols []repository.ColInfo) Node {
	paginationButton := func(icon string, page int) Node {
		return Button(Class("px-3 py-1 min-h-9 text-sm font-normal text-neutral-800 transition duration-200 ease"), Icon(icon, 16),
			Attr("hx-get", url+repository.QueryParamsFromPagenum(page, f)),
			Attr("hx-swap", "#"+tableId),
			Attr("hx-target", "#"+tableId),
			Attr("hx-select", "#"+tableId),
			Attr("hx-trigger", "click"),
		)
	}

	return Group{
		Form(AutoComplete("off"), Attr("hx-get", url), Attr("hx-trigger", "keyup delay:100ms from:input, click from:select"), Attr("hx-swap", "outerHTML"), Attr("hx-target", "#"+tableId), Attr("hx-select", "#"+tableId),
			Div(Class("w-full flex justify-between items-center mb-3 mt-1 pl-3"),
				Div(
					H3(Class("text-lg font-semibold text-neutral-800"), Text("Example Orders")),
					P(Class("text-neutral-500"), Text("An example `orders` table with filtering, pagination, and sorting.")),
				),
				Div(Class("ml-3"),
					Div(Class("w-full max-w-sm min-w-[200px] relative"),
						Div(Class("relative"),
							Input(Class("bg-white w-full pr-11 h-10 pl-3 py-2 bg-transparent placeholder:text-neutral-400 text-neutral-700 text-sm border border-neutral-200 transition duration-200 ease focus:outline-none focus:border-neutral-400 hover:border-neutral-400 shadow-sm focus:shadow-md"), Placeholder("Search..."), Name("search"), AutoFocus()),
							Input(Type("hidden"), Name("orderBy"), Value(f.OrderBy)),
							Input(Type("hidden"), Name("orderDescending"), Value(ToString(f.OrderDescending))),
							Input(Type("hidden"), Name("itemsPerPage"), Value(ToString(f.Pagination.MaxItemsPerPage))),
							Button(Class("absolute h-8 w-8 right-1 top-1 my-auto px-2 flex items-center bg-white "), Type("button"),
								Icon(ICON_SEARCH, 24),
							),
						),
					),
				),
			),
		),
		Div(ID(tableId),
			Div(Class("relative flex flex-col w-full h-full text-gray-700 bg-white shadow-md"),
				Div(Class("overflow-scroll flex flex-col w-full h-full"),
					Table(Class("table-fixed"),
						THead(
							Tr(
								Map(cols, func(col repository.ColInfo) Node {
									return Th(Class("p-4 border-b border-neutral-200 transition-colors cursor-pointer bg-neutral-100 hover:bg-neutral-200"),
										Attr("hx-get", url+repository.QueryParamsFromOrderBy(col.DbName, !f.OrderDescending && (col.DbName == f.OrderBy), f)),
										Attr("hx-swap", "#"+tableId),
										Attr("hx-target", "#"+tableId),
										Attr("hx-select", "#"+tableId),
										Attr("hx-trigger", "click"),
										P(Classes{"flex items-center justify-between gap-2 font-sans text-sm leading-none text-neutral-500": true, "font-bold": f.OrderBy == col.DbName, "font-normal": f.OrderBy != col.DbName}, Text(col.DisplayName),
											If(f.OrderBy == col.DbName,
												IfElse(f.OrderDescending,
													Icon(ICON_ARROW_DOWN_WIDE_NARROW, 16),
													Icon(ICON_ARROW_UP_WIDE_NARROW, 16),
												),
											),
										),
									)
								}),
							),
						),
						TBody(
							Map(entities, nf),
						),
					),
				),
				Div(Class("flex justify-between items-center px-4 py-3"),
					Div(Class("text-sm text-neutral-500"),
						B(Icon(ICON_LIST_ORDERED, 16)),
						Span(Class("mr-3")), ToText(f.Pagination.ViewRangeLower), Text("-"), ToText(f.Pagination.ViewRangeUpper), Text(" of "), ToText(f.Pagination.TotalItems),
					),

					Div(Class("flex"),
						Div(Class("content-center text-sm text-neutral-500"),
							Span(Text("Items per page:")),
						),
						Form(Class("py-3 px-3 min-h-9 block"),
							Attr("hx-get", url),
							Attr("hx-trigger", "click from:select"),
							Attr("hx-target", "#"+tableId),
							Attr("hx-select", "#"+tableId),
							Attr("hx-swap", "outerHTML"),
							Input(Type("hidden"), Name("search"), Value(f.Search)),
							Input(Type("hidden"), Name("orderBy"), Value(f.OrderBy)),
							Input(Type("hidden"), Name("orderDescending"), Value(ToString(f.OrderDescending))),
							Select(
								Class("bg-gray-50 py-3 px-3 min-h-9 border border-gray-300 text-gray-900 text-sm block p-1.5 shadow-sm"),
								Name("itemsPerPage"),
								Option(If(f.Pagination.MaxItemsPerPage == 5, Selected()), Text("5")),
								Option(If(f.Pagination.MaxItemsPerPage == 10, Selected()), Text("10")),
								Option(If(f.Pagination.MaxItemsPerPage == 25, Selected()), Text("25")),
								Option(If(f.Pagination.MaxItemsPerPage == 50, Selected()), Text("50")),
								Option(If(f.Pagination.MaxItemsPerPage == 100, Selected()), Text("100")),
							),
						),

						paginationButton(ICON_CHEVRON_FIRST, 1),
						paginationButton(ICON_CHEVRON_LEFT, f.Pagination.PreviousPage),

						Div(Class("text-sm content-center text-neutral-500 py-3 px-3 min-h-9"),
							Text("Page "),
							ToText(f.Pagination.CurrentPage),
							Text(" of "),
							ToText(f.Pagination.TotalPages),
						),

						paginationButton(ICON_CHEVRON_RIGHT, f.Pagination.NextPage),
						paginationButton(ICON_CHEVRON_LAST, f.Pagination.TotalPages),
					),
				),
			),
		),
	}
}
