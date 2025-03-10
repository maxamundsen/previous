package orders

import (
	"previous/database"
	"previous/finance"
)

type Order struct {
	ID             int32  `db:"id"`
	ProductID      int32  `db:"product_id"`
	Price          int32  `db:"price"`
	PurchaserName  string `db:"purchaser_name"`
	PurchaserEmail string `db:"purchaser_email"`
}

func Fetch() ([]Order, error) {
	qb := &database.QueryBuilder{}
	qb.BaseSQL = "SELECT * FROM orders"

	return database.Select[Order](qb, database.DB)
}

func Filter(f database.Filter) ([]Order, error) {
	qb := &database.QueryBuilder{}
	qb.BaseSQL = "SELECT * FROM orders"

	// search filters
	purchaserSearch := f.Search["purchaser_name"]
	if purchaserSearch != "" {
		qb.Where = append(qb.Where, database.QueryFilter{
			Column: "purchaser_name", Operator: database.LIKE, Parameter: database.Wildcard(purchaserSearch),
		})
	}

	emailSearch := f.Search["purchaser_email"]
	if emailSearch != "" {
		qb.Where = append(qb.Where, database.QueryFilter{
			Column: "purchaser_email", Operator: database.LIKE, Parameter: database.Wildcard(emailSearch),
		})
	}

	// between filters
	priceSearchLeft_string := f.Search["price_left"]
	priceSearchRight_string := f.Search["price_right"]

	priceSearchLeft := int32(finance.MoneyToInt64(priceSearchLeft_string))
	priceSearchRight := int32(finance.MoneyToInt64(priceSearchRight_string))

	if priceSearchLeft_string != "" {
		qb.Where = append(qb.Where, database.QueryFilter{
			Column: "price", Operator: database.GE, Parameter: priceSearchLeft,
		})
	}

	if priceSearchRight_string != "" {
		qb.Where = append(qb.Where, database.QueryFilter{
			Column: "price", Operator: database.LE, Parameter: priceSearchRight,
		})
	}

	// pagination, orderby
	database.SetBuilderFromFilter(qb, f)

	return database.Select[Order](qb, database.DB)
}
