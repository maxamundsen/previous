package repository

import (
	"fmt"
	"net/http"
	"strconv"

	. "github.com/go-jet/jet/v2/sqlite"

	. "previous/basic"
)

type ColInfo struct {
	DisplayName string
	DbName      string
}

type Filter struct {
	Search          string
	Pagination      Pagination
	OrderBy         string
	OrderDescending bool
}

type Pagination struct {
	CurrentPage           int
	NextPage              int
	PreviousPage          int
	TotalPages            int
	TotalItems            int
	MaxItemsPerPage          int
	ItemsThisPage int
	ViewRangeLower        int
	ViewRangeUpper        int
}

func ParseFilterFromRequest(r *http.Request) Filter {
	filter := Filter{}

	filter.Search = r.URL.Query().Get("search")
	filter.OrderBy = r.URL.Query().Get("orderBy")
	filter.OrderDescending, _ = strconv.ParseBool(r.URL.Query().Get("desc"))
	filter.Pagination.CurrentPage, _ = strconv.Atoi(r.URL.Query().Get("pageNum"))
	filter.Pagination.MaxItemsPerPage, _ = strconv.Atoi(r.URL.Query().Get("itemsPerPage"))

	return filter
}

func QueryParamsFromPagenum(pageNum int, f Filter) string {
	f.Pagination.CurrentPage = pageNum

	return QueryParamsFromFilter(f)
}

func QueryParamsFromOrderBy(orderBy string, direction bool, f Filter) string {
	f.OrderBy = orderBy
	f.OrderDescending = direction
	f.Pagination.CurrentPage = 1

	return QueryParamsFromFilter(f)
}

func QueryParamsFromFilter(f Filter) string {
	return fmt.Sprintf(
		"?search=%s&orderBy=%s&desc=%t&pageNum=%d&itemsPerPage=%d",
		f.Search,
		f.OrderBy,
		f.OrderDescending,
		f.Pagination.CurrentPage,
		f.Pagination.MaxItemsPerPage,
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

func GetColInfoFromJet(cl ColumnList) []ColInfo {
	list := []ColInfo{}

	for _, col := range cl {
		newCol := ColInfo{
			DisplayName: SnakeCaseToTitleCase(col.Name()),
			DbName:      col.Name(),
		}

		list = append(list, newCol)
	}

	return list
}

func (p *Pagination) GeneratePagination() {
	if p.MaxItemsPerPage == 0 {
		p.TotalPages = 1
	} else {
		p.TotalPages = p.TotalItems / p.MaxItemsPerPage

		// Gofmt sucks ass wtf is this ???
		// It won't let me put a space between the modulo sign
		if p.TotalItems%p.MaxItemsPerPage != 0 {
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

	if p.TotalItems != 0 {
		p.ViewRangeLower = p.MaxItemsPerPage * p.CurrentPage - p.MaxItemsPerPage + 1
	} else {
		p.ViewRangeLower = 0
	}
	p.ViewRangeUpper = p.MaxItemsPerPage * p.CurrentPage - p.MaxItemsPerPage + p.ItemsThisPage

	if p.CurrentPage >= p.TotalPages {
		p.CurrentPage = p.TotalPages
		p.NextPage = p.TotalPages
	} else {
		p.NextPage = p.CurrentPage + 1
	}
}
