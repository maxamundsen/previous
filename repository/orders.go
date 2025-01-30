package repository

import (
	"previous/.jet/model"
	"previous/finance"

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

	condition := Bool(true)

	// search filters
	emailSearch := f.Search[Order.PurchaserEmail.Name()]
	if emailSearch != "" {
		condition = condition.AND(Order.PurchaserEmail.LIKE(String("%" + emailSearch + "%")))
	}

	purchaserSearch := f.Search[Order.PurchaserName.Name()]
	if purchaserSearch != "" {
		condition = condition.AND(Order.PurchaserName.LIKE(String("%" + purchaserSearch + "%")))
	}

	// between filters
	priceSearchLeft_string := f.Search[Order.Price.Name() + "_left"]
	priceSearchRight_string := f.Search[Order.Price.Name() + "_right"]

	priceSearchLeft := Int32(int32(finance.MoneyToInt64(priceSearchLeft_string)))
	priceSearchRight := Int32(int32(finance.MoneyToInt64(priceSearchRight_string)))

	if priceSearchLeft_string != "" {
		condition = condition.AND(Order.Price.GT_EQ(priceSearchLeft))
	}

	if priceSearchRight_string != "" {
		condition = condition.AND(Order.Price.LT_EQ(priceSearchRight))
	}

	stmt.WHERE(condition)

	// order by
	obCol, exists := GetColumnFromStringName(f.OrderBy, Order.AllColumns)

	if exists {
		if f.OrderDescending {
			stmt.ORDER_BY(obCol.DESC())
		} else {
			stmt.ORDER_BY(obCol.ASC())
		}
	} else {
		stmt.ORDER_BY(Order.ID.ASC())
	}

	// pagination
	if f.Pagination.Enabled {
		if f.Pagination.MaxItemsPerPage > 0 {
			stmt.LIMIT(int64(f.Pagination.MaxItemsPerPage))
			stmt.OFFSET(int64((f.Pagination.CurrentPage - 1) * f.Pagination.MaxItemsPerPage))
		}
	}

	// fmt.Println(stmt.DebugSql())

	err := stmt.Query(db, &orders)
	return orders, err
}
