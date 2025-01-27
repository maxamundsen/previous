package repository

import (
	"previous/.jet/model"

	. "previous/.jet/table"

	. "github.com/go-jet/jet/v2/sqlite"
)

type OrderRepository struct{}

func (o OrderRepository) Fetch() ([]model.Order, error) {
	orders := []model.Order{}

	stmt := SELECT(Order.AllColumns).FROM(Order)
	err := stmt.Query(db, &orders)

	return orders, err
}

func (o OrderRepository) Filter(f Filter) ([]model.Order, error) {
	orders := []model.Order{}

	stmt := SELECT(Order.AllColumns).FROM(Order)

	// where filters
	purchaserSearch := GetSearchFilterValueFromColName(f.Search, Order.PurchaserName.Name())
	if purchaserSearch != "" {
		stmt.WHERE(Order.PurchaserName.LIKE(String("%" + purchaserSearch + "%")))
	}

	// order by
	obCol, exists := GetColumnFromStringName(f.OrderBy, Order.AllColumns)

	if exists {
		if f.OrderDescending {
			stmt.ORDER_BY(obCol.DESC())
		} else {
			stmt.ORDER_BY(obCol.ASC())
		}
	}

	// pagination
	if f.Pagination.MaxItemsPerPage > 0 {
		stmt.LIMIT(int64(f.Pagination.MaxItemsPerPage))
		stmt.OFFSET(int64((f.Pagination.CurrentPage - 1) * f.Pagination.MaxItemsPerPage))
	}

	err := stmt.Query(db, &orders)
	return orders, err
}
