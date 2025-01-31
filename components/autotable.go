package components

import (
	. "previous/basic"

	"previous/repository"

	. "maragu.dev/gomponents"
	. "maragu.dev/gomponents/components"
	. "maragu.dev/gomponents/html"
)

const TABLE_ABOVE_PREFIX = "_tableAbove"
const FORM_BIND_SUFFIX = "_form"
const FORM_PAGINATION_SUFFIX = "_paginationForm"

// FILTER COMPONENTS

// Autotable bind search to col
func BindSearch(elId string, identifier string) Node {
	return Group{
		FormAttr(elId + FORM_BIND_SUFFIX),
		Name(repository.SEARCH_URL_KEY_PREFIX + identifier),
	}
}

// THE TABLE
// Note that "aboveTable" node is not swapped with HTMX, but "belowTable" is.
func AutoTable[E any](tableId string, url string, f repository.Filter, entities []E, aboveTable Node, nf func(E) Node, cols []repository.ColInfo, belowTable Node) Node {
	paginationButton := func(icon string, page int) Node {
		return Button(Class("px-3 py-1 min-h-9 text-sm font-normal text-neutral-800 transition duration-200 ease cursor-pointer"), Icon(icon, 16),
			Attr("hx-get", url+repository.QueryParamsFromPagenum(page, f)),
			Attr("hx-swap", "#"+tableId),
			Attr("hx-target", "#"+tableId),
			Attr("hx-select", "#"+tableId),
			Attr("hx-trigger", "click"),
		)
	}

	return Group{
		Div(ID(tableId+TABLE_ABOVE_PREFIX),
			aboveTable,
		),
		Div(ID(tableId),
			Form(ID(tableId+FORM_BIND_SUFFIX),
				AutoComplete("off"),
				Attr("hx-get", url),
				Attr("hx-trigger", "keyup delay:100ms from:(#"+tableId+TABLE_ABOVE_PREFIX+" input), change from:(#"+tableId+TABLE_ABOVE_PREFIX+" input[type=date]), change from:(#"+tableId+TABLE_ABOVE_PREFIX+" input[type=datetime-local]), change from:(#"+tableId+TABLE_ABOVE_PREFIX+" select)"),
				Attr("hx-swap", "outerHTML"), Attr("hx-target", "#"+tableId),
				Attr("hx-select", "#"+tableId),
				Input(Type("hidden"), Name(repository.ORDER_BY_URL_KEY), Value(f.OrderBy)),
				Input(Type("hidden"), Name(repository.ORDER_DESC_URL_KEY), Value(ToString(f.OrderDescending))),
				Input(Type("hidden"), Name(repository.ITEMS_PER_PAGE_URL_KEY), Value(ToString(f.Pagination.MaxItemsPerPage))),
			),
			Div(Class("relative flex flex-col w-full h-full text-gray-700 bg-white shadow-md"),
				Div(Class("overflow-scroll flex flex-col w-full h-full"),
					Table(Class("table-fixed"),
						THead(
							Tr(
								Map(cols, func(col repository.ColInfo) Node {
									if col.Sortable {
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
									} else {
										return Th(Class("p-4 border-b border-neutral-200 transition-colors bg-neutral-100"),
											P(Class("flex items-center justify-between gap-2 font-sans text-sm leading-none text-neutral-500 font-normal"),
												Text(col.DisplayName),
											),
										)
									}
								}),
							),
						),
						TBody(
							IfElse(len(entities) > 0,
								Map(entities, nf),
								Td(Class("p-4 py-5"),
									P(Class("block text-sm text-neutral-800"), Text("Dataset contains no entries.")),
								),
							),
						),
					),
				),
				If(f.Pagination.Enabled,
					Div(Class("flex justify-between items-center px-4 py-3"),
						Div(Class("text-sm text-neutral-500"),
							B(Icon(ICON_LIST_ORDERED, 16)),
							Span(Class("mr-3")), ToText(f.Pagination.ViewRangeLower), Text("-"), ToText(f.Pagination.ViewRangeUpper), Text(" of "), ToText(f.Pagination.TotalItems),
						),

						Div(Class("flex"),
							Div(Class("content-center text-sm text-neutral-500"),
								Span(Text("Items per page:")),
							),
							Form(ID(tableId+FORM_PAGINATION_SUFFIX), Class("py-3 px-3 min-h-9 block"),
								Attr("hx-get", url),
								Attr("hx-trigger", "change from:(#"+tableId+FORM_PAGINATION_SUFFIX+" select)"),
								Attr("hx-target", "#"+tableId),
								Attr("hx-select", "#"+tableId),
								Attr("hx-swap", "outerHTML"),

								MapMapWithKey(f.Search, func(s string, v string) Node {
									return Input(Type("hidden"), Name(repository.SEARCH_URL_KEY_PREFIX+s), Value(v))
								}),

								Input(Type("hidden"), Name(repository.ORDER_BY_URL_KEY), Value(f.OrderBy)),
								Input(Type("hidden"), Name(repository.ORDER_DESC_URL_KEY), Value(ToString(f.OrderDescending))),
								Select(
									Class("bg-gray-50 py-3 px-3 min-h-9 border border-gray-300 text-gray-900 text-sm block p-1.5 shadow-sm"),
									Name(repository.ITEMS_PER_PAGE_URL_KEY),
									Option(If(f.Pagination.MaxItemsPerPage == 5, Selected()), Text("5"), Value("5")),
									Option(If(f.Pagination.MaxItemsPerPage == 10, Selected()), Text("10"), Value("10")),
									Option(If(f.Pagination.MaxItemsPerPage == 25, Selected()), Text("25"), Value("25")),
									Option(If(f.Pagination.MaxItemsPerPage == 50, Selected()), Text("50"), Value("50")),
									Option(If(f.Pagination.MaxItemsPerPage == 100, Selected()), Text("100"), Value("100")),
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
			belowTable,
		),
	}
}
