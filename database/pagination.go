package database

type Pagination struct {
	CurrentPage  int
	NextPage     int
	PreviousPage int
	TotalPages   int
	TotalItems   int
}

func ProcessPageNum(pageNum int, itemsPerPage int, totalItems int) Pagination {
	var nextPage int
	var previousPage int
	var totalPages int

	if itemsPerPage == 0 {
		totalPages = 1
	} else {
		totalPages = totalItems / itemsPerPage

		// Gofmt sucks ass wtf is this ???
		// It won't let me put a space between the modulo sign
		if totalItems%itemsPerPage != 0 {
			totalPages++
		}
	}

	if totalPages == 0 {
		totalPages = 1
	}

	if pageNum < 1 {
		pageNum = 1
		previousPage = 1
	} else {
		previousPage = pageNum - 1
	}

	if pageNum >= totalPages {
		pageNum = totalPages
		nextPage = totalPages
	} else {
		nextPage = pageNum + 1
	}

	return Pagination{
		CurrentPage:  pageNum,
		NextPage:     nextPage,
		PreviousPage: previousPage,
		TotalPages:   totalPages,
		TotalItems:   totalItems,
	}
}
