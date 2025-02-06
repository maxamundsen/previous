package repository

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"

	. "github.com/go-jet/jet/v2/sqlite"

	. "previous/basic"
)

const (
	BETWEEN_LEFT_URL_KEY_PREFIX  = "betweenLeft_"
	BETWEEN_RIGHT_URL_KEY_PREFIX = "betweenRight_"
	ORDER_BY_URL_KEY             = "orderBy"
	ORDER_DESC_URL_KEY           = "desc"
	PAGE_NUM_URL_KEY             = "pageNum"
	ITEMS_PER_PAGE_URL_KEY       = "itemsPerPage"
	SEARCH_URL_KEY_PREFIX        = "search_"
	FILTER_DEFAULT_MAX_ITEMS     = 10
)

type ColInfo struct {
	DisplayName string
	DbName      string
	Sortable    bool
}

type Filter struct {
	Search          map[string]string
	Pagination      Pagination
	OrderBy         string
	OrderDescending bool
}

type Pagination struct {
	Enabled         bool
	CurrentPage     int
	NextPage        int
	PreviousPage    int
	TotalPages      int
	TotalItems      int
	MaxItemsPerPage int
	ItemsThisPage   int
	ViewRangeLower  int
	ViewRangeUpper  int
}

func ParseFilterFromRequest(r *http.Request) Filter {
	filter := Filter{}
	filter.Search = make(map[string]string)

	for k, v := range r.URL.Query() {
		if strings.HasPrefix(k, SEARCH_URL_KEY_PREFIX) {
			qValue := strings.Join(v, "")

			if qValue != "" {
				filter.Search[strings.TrimPrefix(k, SEARCH_URL_KEY_PREFIX)]  = qValue
			}
		}
	}

	filter.OrderBy = r.URL.Query().Get(ORDER_BY_URL_KEY)
	filter.OrderDescending, _ = strconv.ParseBool(r.URL.Query().Get(ORDER_DESC_URL_KEY))
	filter.Pagination.CurrentPage, _ = strconv.Atoi(r.URL.Query().Get(PAGE_NUM_URL_KEY))
	filter.Pagination.MaxItemsPerPage, _ = strconv.Atoi(r.URL.Query().Get(ITEMS_PER_PAGE_URL_KEY))

	if filter.Pagination.MaxItemsPerPage == 0 {
		filter.Pagination.MaxItemsPerPage = FILTER_DEFAULT_MAX_ITEMS
	}

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
	output := fmt.Sprintf(
		"?"+ORDER_BY_URL_KEY+"=%s&"+ORDER_DESC_URL_KEY+"=%t&"+PAGE_NUM_URL_KEY+"=%d&"+ITEMS_PER_PAGE_URL_KEY+"=%d",
		f.OrderBy,
		f.OrderDescending,
		f.Pagination.CurrentPage,
		f.Pagination.MaxItemsPerPage,
	)

	for k, v := range f.Search {
		output += "&" + SEARCH_URL_KEY_PREFIX + k + "=" + v
	}

	return output
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

func (p *Pagination) GeneratePagination(totalItemsInSet int, itemsDisplayedThisPage int) {
	p.TotalItems = totalItemsInSet
	p.ItemsThisPage = itemsDisplayedThisPage

	if p.MaxItemsPerPage == 0 {
		p.MaxItemsPerPage = FILTER_DEFAULT_MAX_ITEMS
	}

	if p.MaxItemsPerPage == 0 {
		p.TotalPages = 1
	} else {
		p.TotalPages = p.TotalItems / p.MaxItemsPerPage

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
		p.ViewRangeLower = p.MaxItemsPerPage*p.CurrentPage - p.MaxItemsPerPage + 1
	} else {
		p.ViewRangeLower = 0
	}
	p.ViewRangeUpper = p.MaxItemsPerPage*p.CurrentPage - p.MaxItemsPerPage + p.ItemsThisPage

	if p.CurrentPage >= p.TotalPages {
		p.CurrentPage = p.TotalPages
		p.NextPage = p.TotalPages
	} else {
		p.NextPage = p.CurrentPage + 1
	}
}
