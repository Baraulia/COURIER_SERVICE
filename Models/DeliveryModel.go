package Models

import "github.com/Baraulia/COURIER_SERVICE/db"

func GetOrders(Orders []db.Order) []db.Order {
	Orders = []db.Order{}
	db.GetActiveOrdersFromDB(&Orders)
	return Orders
}

func GetOneOrder(Order db.Order, id string) db.Order {
	Order = db.Order{}
	db.GetActiveOrderFromDB(&Order, id)
	return Order
}
