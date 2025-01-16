package repository

import . "github.com/go-jet/jet/v2/sqlite"

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
