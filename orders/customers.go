package orders

import (
	"previous/database"
	"sort"
	"strings"
)

type Customer struct {
	Firstname string
	Lastname  string
	Email     string
	Phone     int
}

var CustomerList = []Customer{
	{Firstname: "John", Lastname: "Doe", Email: "john@example.com", Phone: 1234567},
	{Firstname: "Jane", Lastname: "Doe", Email: "jane@example.com", Phone: 1234567},
	{Firstname: "Sally", Lastname: "Smith", Email: "sally@example.com", Phone: 1234567},
	{Firstname: "Max", Lastname: "Amundsen", Email: "max@example.com", Phone: 1234567},
	{Firstname: "Peter", Lastname: "Piper", Email: "peter@example.com", Phone: 1234567},
	{Firstname: "Alice", Lastname: "Wonderland", Email: "alice@example.com", Phone: 1234567},
	{Firstname: "Bob", Lastname: "Pinciotti", Email: "bob@example.com", Phone: 1234567},
	{Firstname: "Matthew", Lastname: "Mark", Email: "matthew@example.com", Phone: 1234567},
	{Firstname: "Luke", Lastname: "John", Email: "luke@example.com", Phone: 1234567},
	{Firstname: "Ringo", Lastname: "Starr", Email: "ringo@example.com", Phone: 1234567},
}

func FetchCustomers() ([]Customer, error) {
	return CustomerList, nil
}

func FilterCustomers(f database.Filter) ([]Customer, error) {
	customers := CustomerList

	output := []Customer{}

	// Search Filters
	for _, v := range customers {
		search := strings.ToLower(f.Search["name"])
		searchable := strings.ToLower(v.Firstname + " " + v.Lastname)

		if search != "" {
			if strings.Contains(searchable, search) {
				output = append(output, v)
			}
		} else {
			output = append(output, v)
		}
	}

	// Order By
	orderByFunc := func (customerList []Customer, desc bool) []Customer {
		sort.Slice(customerList, func(i, j int) bool {
			if desc {
				return customerList[i].Lastname > customerList[j].Lastname
			} else {
				return customerList[i].Lastname < customerList[j].Lastname
			}
		})
		return customerList
	}

	output = orderByFunc(output, f.OrderDescending)

	// pagination
	output = database.PaginateSlice(output, f)

	return output, nil
}