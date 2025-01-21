package components

import (
	"previous/repository"

	. "maragu.dev/gomponents"
	. "maragu.dev/gomponents/html"

)

func AutoTableClickHandler(url string, tableId string) Node {
	return HxOnClick(url, "outerHTML", "#" + tableId, "#" + tableId)
}

func AutoTable[E any](tableId string, url string, f repository.Filter, entities []E, nf func(E) Node, cols []string) Node {
	return Group{
		Form(AutoComplete("off"), Attr("hx-get", url), Attr("hx-trigger", "keyup delay:100ms from:input, click from:select"), Attr("hx-swap", "outerHTML"), Attr("hx-target", "#"+tableId), Attr("hx-select", "#"+tableId),
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
		Div(ID(tableId),
			Div(Class("relative flex flex-col w-full h-full overflow-scroll text-gray-700 bg-white shadow-md rounded-lg bg-clip-border"),
				Table(Class("table-fixed min-w-max"),
					THead(
						Tr(
							Map(cols, func(s string) Node {
								return Th(AutoTableClickHandler(url + repository.QueryParamsFromOrderBy(s, !f.OrderDescending, f), tableId), Class("p-4 border-b border-slate-200 transition-colors cursor-pointer bg-slate-50 hover:bg-slate-100"),
									P(Class("flex items-center justify-between gap-2 font-sans text-sm font-normal leading-none text-slate-500"), Text(s),
										Icon("sort-arrows", 16),
									),
								)
							}),
						),
					),
					TBody(
						Map(entities, nf),
					),
				),
				Div(Class("flex justify-between items-center px-4 py-3"),
					Div(Class("text-sm text-slate-500"), Text("Showing "), ToText(f.Pagination.CurrentPage), Text(" of "), ToText(f.Pagination.TotalPages),
						B(Text(" Total Items: ")), ToText(f.Pagination.TotalItems),
					),
					Div(Class("flex space-x-1"),
						Button(Class("px-3 py-1 min-w-9 min-h-9 text-sm font-normal text-slate-500 bg-white border border-slate-200 rounded hover:bg-slate-50 hover:border-slate-400 transition duration-200 ease"), Text("Prev"),
							AutoTableClickHandler(url + repository.QueryParamsFromPagenum(f.Pagination.PreviousPage, f), tableId),
						),

						Button(Class("px-3 py-1 min-w-9 min-h-9 text-sm font-normal text-slate-500 bg-white border border-slate-200 rounded hover:bg-slate-50 hover:border-slate-400 transition duration-200 ease"), Text("Next"),
							AutoTableClickHandler(url + repository.QueryParamsFromPagenum(f.Pagination.NextPage, f), tableId),
						),
					),
				),
			),
		),
	}
}
