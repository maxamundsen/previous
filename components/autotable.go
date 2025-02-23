package components

import (
	. "previous/basic"

	"previous/repository"

	. "maragu.dev/gomponents"
	hx "maragu.dev/gomponents-htmx"
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
func AutoTable[E any](tableId string, url string, cols []repository.ColInfo, f repository.Filter, entities []E, aboveTable Node, nf func(E) Node, belowTable Node) Node {
	paginationButton := func(icon string, page int) Node {
		return Button(
			InlineStyle(`
				$me {
					padding-left: $(3);
					padding-right: $(3);
					padding-top: $(1);
					padding-bottom: $(1);
					min-height: $(9);
					font-size: var(--text-sm);
					font-weight: var(--font-weight-normal);
					color: $color(neutral-800);
					cursor: pointer;
				}
			`),
			Icon(icon, 16),
			hx.Get(url+repository.QueryParamsFromPagenum(page, f)),
			hx.Swap(CSSID(tableId)),
			hx.Target(CSSID(tableId)),
			hx.Select(CSSID(tableId)),
			hx.Trigger("click"),
		)
	}

	return Group{
		Div(ID(tableId+TABLE_ABOVE_PREFIX),
			aboveTable,
		),
		Div(ID(tableId),
			Form(ID(tableId+FORM_BIND_SUFFIX),
				AutoComplete("off"),
				hx.Get(url),
				hx.Trigger("keyup delay:100ms from:(#"+tableId+TABLE_ABOVE_PREFIX+" input), change from:(#"+tableId+TABLE_ABOVE_PREFIX+" input[type=date]), change from:(#"+tableId+TABLE_ABOVE_PREFIX+" input[type=datetime-local]), change from:(#"+tableId+TABLE_ABOVE_PREFIX+" select)"),
				hx.Swap("outerHTML"),
				hx.Target(CSSID(tableId)),
				hx.Select(CSSID(tableId)),
				Input(Type("hidden"), Name(repository.ORDER_BY_URL_KEY), Value(f.OrderBy)),
				Input(Type("hidden"), Name(repository.ORDER_DESC_URL_KEY), Value(ToString(f.OrderDescending))),
				Input(Type("hidden"), Name(repository.ITEMS_PER_PAGE_URL_KEY), Value(ToString(f.Pagination.MaxItemsPerPage))),
			),
			Div(
				InlineStyle(`
					$me {
						position: relative;
						display: flex;
						flex-direction: column;
						width: 100%;
						height: 100%;
						color: $color(gray-700);
						background-color: $color(white);
						box-shadow: var(--shadow-md);
					}
				`),
				Div(
					InlineStyle(`
						$me {
							overflow: scroll;
							display: flex;
							flex-direction: column;
							width: 100%;
							height: 100%;
						}
					`),
					Table(
						InlineStyle("$me { table-layout: fixed; width: 100%; }"),
						THead(
							Tr(
								Map(cols, func(col repository.ColInfo) Node {
									if col.Sortable {
										return Th(
											InlineStyle(`
												$me {
													padding: $(4);
													border-bottom: 1px solid $color(neutral-200);
													cursor: pointer;
													background-color: $color(neutral-100);
													transition-property: color, background-color, border-color, text-decoration-color, fill, stroke, --tw-gradient-from, --tw-gradient-via, --tw-gradient-to;
													transition-timing-function: var(--default-transition-timing-function);
													transition-duration: var(--default-transition-duration);
												}

												$me:hover {
													background-color: $color(neutral-200);
												}
											`),
											hx.Get(url+repository.QueryParamsFromOrderBy(col.DbName, !f.OrderDescending && (col.DbName == f.OrderBy), f)),
											hx.Swap(CSSID(tableId)),
											hx.Target(CSSID(tableId)),
											hx.Select(CSSID(tableId)),
											hx.Trigger("click"),
											P(
												InlineStyle(`
													$me {
														display: flex;
														align-items: center;
														justify-content: space-between;
														gap: $(2);
														font-size: var(--text-sm);
														color: $color(neutral-500);
													}
												`),
												Text(col.DisplayName),
												IfElse(f.OrderBy == col.DbName,
													Group{
														InlineStyle("$me { font-weight: var(--font-weight-bold) ; }"),
														IfElse(f.OrderDescending,
															Icon(ICON_ARROW_DOWN_WIDE_NARROW, 16),
															Icon(ICON_ARROW_UP_WIDE_NARROW, 16),
														),
													},
													InlineStyle("$me { font-weight: var(--font-weight-normal); }"),
												),
											),
										)
									} else {
										return Th(
											InlineStyle(`
												$me {
													padding: $(4);
													border-bottom: 1px solid $color(neutral-200);
													background-color: $color(neutral-100);
												}
											`),
											P(
												InlineStyle(`
													$me {
														display: flex;
														align-items: center;
														justify-content: space-between;
														gap: $(2);
														font-size: var(--text-sm);
														font-weight: var(--font-weight-normal);
														color: $color(neutral-500);
													}
												`),
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
								Td(InlineStyle("$me { padding: $(4); }"),
									P(InlineStyle("$me {display: block; font-size: var(--text-sm); color: $color(neutral-800);}"), Text("Dataset contains no entries.")),
								),
							),
						),
					),
				),
				If(f.Pagination.Enabled,
					Div(InlineStyle("$me { display: flex; justify-content: space-between; align-items: center; padding: $(3) $(4); }"),
						Div(
							InlineStyle("$me { display: flex; align-items: center; font-size: var(--text-sm); color: $color(neutral-500); }"),
							B(Icon(ICON_LIST_ORDERED, 16)),
							Span(InlineStyle("$me {margin-right: $(3);}")), ToText(f.Pagination.ViewRangeLower), Text("-"), ToText(f.Pagination.ViewRangeUpper), Text(" of "), ToText(f.Pagination.TotalItems),
						),

						Div(InlineStyle("$me { display: flex; }"),
							Div(InlineStyle("$me { margin-right: $(3); align-content: center; font-size: var(--text-sm); color: $color(neutral-500);}"),
								Span(Text("Items per page:")),
							),
							Form(
								ID(tableId+FORM_PAGINATION_SUFFIX),
								hx.Get(url),
								hx.Trigger("change from:(#"+tableId+FORM_PAGINATION_SUFFIX+" select)"),
								hx.Target(CSSID(tableId)),
								hx.Select(CSSID(tableId)),
								hx.Swap("outerHTML"),

								MapMapWithKey(f.Search, func(s string, v string) Node {
									return Input(Type("hidden"), Name(repository.SEARCH_URL_KEY_PREFIX+s), Value(v))
								}),

								Input(Type("hidden"), Name(repository.ORDER_BY_URL_KEY), Value(f.OrderBy)),
								Input(Type("hidden"), Name(repository.ORDER_DESC_URL_KEY), Value(ToString(f.OrderDescending))),
								Div(
									InlineStyle(`
										$me > select {
											background-color: $color(gray-50);
											padding: $(3);
											border: 1px solid $color(gray-50);
											font-size: var(--text-sm);
											display: block;
											box-shadow: var(--shadow-sm);
											margin-right: $(3);
										}
									`),
									Select(
										Name(repository.ITEMS_PER_PAGE_URL_KEY),
										Option(If(f.Pagination.MaxItemsPerPage == 5, Selected()), Text("5"), Value("5")),
										Option(If(f.Pagination.MaxItemsPerPage == 10, Selected()), Text("10"), Value("10")),
										Option(If(f.Pagination.MaxItemsPerPage == 25, Selected()), Text("25"), Value("25")),
										Option(If(f.Pagination.MaxItemsPerPage == 50, Selected()), Text("50"), Value("50")),
										Option(If(f.Pagination.MaxItemsPerPage == 100, Selected()), Text("100"), Value("100")),
									),
								),
							),

							paginationButton(ICON_CHEVRON_FIRST, 1),
							paginationButton(ICON_CHEVRON_LEFT, f.Pagination.PreviousPage),

							Div(
								InlineStyle("$me { font-size: var(--text-sm); display: flex; justify-content: center; color: $color(neutral-500); padding: $(3); min-height: $(9); }"),
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

func AutotableRow(children ...Node) Node {
	return Tr(InlineStyle("$me:hover{ background-color: $color(neutral-50);} $me {border-bottom: 1px solid $color(neutral-200); }"),
		Group(children),
	)
}

func AutotableItem(children ...Node) Node {
	return Td(InlineStyle("$me { padding: $(4); }"),
		P(InlineStyle("$me { width: 100%; display: block; font-size: var(--text-sm); color: $color(neutral-800); }"),
			Group(children),
		),
	)
}

func AutotableItemBold(children ...Node) Node {
	return Td(InlineStyle("$me { padding: $(4); }"),
		P(InlineStyle("$me { width: 100%; display: block; font-weight: var(--font-weight-semibold); font-size: var(--text-sm); color: $color(neutral-800); }"),
			Group(children),
		),
	)
}

func AutotableSearchGroup(children ...Node) Node {
	return Div(InlineStyle("$me { width: 100%; display: flex; justify-content: space-between; margin-bottom: $(3); margin-top: $(1); }"),
		Div(InlineStyle("$me { width: 100%; position: relative; }"),
			Div(InlineStyle("$me { position: relative; display: flex; flex-direction: row; align-items: center; gap: $(1);}"),
				Group(children),
			),
		),
	)
}

func AutotableSearch(c ...Node) Node {
	return Div(
		InlineStyle(`
			$me {
				width: 100%;
			}

			$me > input {
				background-color: $color(white);
				width: 100%;
				padding-right: $(11);
				padding-left: $(3);
				padding-top: $(2);
				padding-bottom: $(2);
				height: $(10);
				font-size: var(--text-sm);
				border: 1px solid $color(neutral-200);
				box-shadow: var(--shadow-sm);
			}

			$me > input:placeholder {
				color: $color(neutral-400);
			}
		`),
		Input(Group(c)),
	)
}
