package db

import (
	"fmt"
)

type Order struct {
	IdDeliveryService uint16 `json:"delivery_service_id"`
	Id                uint16 `json:"id"`
	IdCourier         uint16 `json:"courier_id"`
	DeliveryTime      string `json:"delivery_time"`
	CustomerAddress   string `json:"customer_address"`
	Status            string `json:"status"`
	OrderDate         string `json:"order_date"`
}

func GetActiveOrdersFromDB(Orders *[]Order) {
	db := ConnectDB()
	defer db.Close()

	insertValue := `Select * from delivery where status = 'ready to delivery'`
	get, err := db.Query(insertValue)
	if err != nil {
		fmt.Println(err)
	}

	for get.Next() {
		var order Order
		err = get.Scan(&order.IdDeliveryService, &order.Id, &order.IdCourier, &order.DeliveryTime, &order.CustomerAddress, &order.Status, &order.OrderDate)
		*Orders = append(*Orders, order)
	}
}

func GetActiveOrderFromDB(Orders *Order, id int) {
	db := ConnectDB()
	defer db.Close()

	insertValue := `Select * from delivery where id = $1 AND status = 'ready to delivery'`
	get, err := db.Query(insertValue, id)
	if err != nil {
		fmt.Println(err)
	}

	for get.Next() {
		var order Order
		err = get.Scan(&order.IdDeliveryService, &order.Id, &order.IdCourier, &order.DeliveryTime, &order.CustomerAddress, &order.Status, &order.OrderDate)
		*Orders = order
	}
}
