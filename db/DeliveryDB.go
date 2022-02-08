package db

import (
	"fmt"
)

type Order struct {
	IdDeliveryService uint16 `json:"id_delivery_service"`
	IdOrder           uint16 `json:"id_order"`
	IdCourier         uint16 `json:"id_courier"`
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
		err = get.Scan(&order.IdDeliveryService, &order.IdOrder, &order.IdCourier, &order.DeliveryTime, &order.CustomerAddress, &order.Status, &order.OrderDate)
		*Orders = append(*Orders, order)
	}
}

func GetActiveOrderFromDB(Orders *Order, id int) {
	db := ConnectDB()
	defer db.Close()

	insertValue := `Select * from delivery where id_order = $1 AND status = 'ready to delivery'`
	get, err := db.Query(insertValue, id)
	if err != nil {
		fmt.Println(err)
	}

	for get.Next() {
		var order Order
		err = get.Scan(&order.IdDeliveryService, &order.IdOrder, &order.IdCourier, &order.DeliveryTime, &order.CustomerAddress, &order.Status, &order.OrderDate)
		*Orders = order
	}
}
