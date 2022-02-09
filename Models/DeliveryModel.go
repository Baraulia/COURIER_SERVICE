package Models

import (
	"github.com/Baraulia/COURIER_SERVICE/db"
)

func GetOrders() []db.Order {
	var Orders []db.Order
	db.GetActiveOrdersFromDB(&Orders)
	return Orders
}

func GetOneOrder(id int) db.Order {
	Order := db.Order{}
	db.GetActiveOrderFromDB(&Order, id)
	return Order
}
