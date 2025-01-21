package repository

import (
	"fmt"
	"net/http"
	"strconv"

	. "github.com/go-jet/jet/v2/sqlite"

	. "previous/basic"
)

type ColFilter struct {
	DisplayName string
	ColumnName  string
}

type Filter struct {
	Search string
	Pagination Pagination
	OrderBy string
	OrderDescending bool
}

type Pagination struct {
	CurrentPage  int
	NextPage     int
	PreviousPage int
	TotalPages   int
	TotalItems   int
	ItemsPerPage int
}

func ParseFilterFromRequest(r *http.Request) Filter {
	filter := Filter{}

	filter.Search = r.URL.Query().Get("search")
	filter.OrderBy = r.URL.Query().Get("orderBy")
	filter.OrderDescending, _ = strconv.ParseBool(r.URL.Query().Get("desc"))
	filter.Pagination.CurrentPage, _ = strconv.Atoi(r.URL.Query().Get("pageNum"))
	filter.Pagination.ItemsPerPage, _ = strconv.Atoi(r.URL.Query().Get("itemsPerPage"))

	return filter
}

func QueryParamsFromPagenum(pageNum int, f Filter) string {
	f.Pagination.CurrentPage = pageNum

	return QueryParamsFromFilter(f)
}

func QueryParamsFromOrderBy(orderBy string, direction bool, f Filter) string {
	f.OrderBy = orderBy
	f.OrderDescending = direction

	return QueryParamsFromFilter(f)
}

func QueryParamsFromFilter(f Filter) string {
	return fmt.Sprintf(
		"?search=%s&orderBy=%s&desc=%t&pageNum=%d&itemsPerPage=%d",
		f.Search,
		f.OrderBy,
		f.OrderDescending,
		f.Pagination.CurrentPage,
		f.Pagination.ItemsPerPage,
	)
}

func GetColumnFromStringName(column string, cl ColumnList) (Column, bool) {
	for _, col := range cl {
		if col.Name() == column {
			return col, true
		}
	}

	return nil, false
}

func GetStringNamesFromColumns(cl ColumnList) []string {
	list := []string{}

	for _, col := range cl {
		list = append(list, col.Name())
	}

	return list
}

func GetFriendlyNamesFromColumns(cl ColumnList) []string {
	list := []string{}

	for _, col := range cl {
		list = append(list, SnakeCaseToTitleCase(col.Name()))
	}

	return list
}

func (p *Pagination) ProcessPageNum() {
	if p.ItemsPerPage == 0 {
		p.TotalPages = 1
	} else {
		p.TotalPages = p.TotalItems / p.ItemsPerPage

		// Gofmt sucks ass wtf is this ???
		// It won't let me put a space between the modulo sign
		if p.TotalItems%p.ItemsPerPage != 0 {
			p.TotalPages++
		}
	}

	if p.TotalPages == 0 {
		p.TotalPages = 1
	}

	if p.CurrentPage < 1 {
		p.CurrentPage = 1
		p.PreviousPage = 1
	} else {
		p.PreviousPage = p.CurrentPage - 1
	}

	if p.CurrentPage >= p.TotalPages {
		p.CurrentPage = p.TotalPages
		p.NextPage = p.TotalPages
	} else {
		p.NextPage = p.CurrentPage + 1
	}
}
