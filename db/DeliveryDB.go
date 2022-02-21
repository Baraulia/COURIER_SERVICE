package db

import (
	"database/sql"
	"fmt"
	"log"
)

type DeliveryPostgres struct {
	db *sql.DB
}

func NewDeliveryPostgres(db *sql.DB) *DeliveryPostgres {
	return &DeliveryPostgres{db: db}
}

type Order struct {
	IdDeliveryService uint16 `json:"delivery_service_id"`
	Id                uint16 `json:"id"`
	IdCourier         uint16 `json:"courier_id"`
	DeliveryTime      string `json:"delivery_time"`
	CustomerAddress   string `json:"customer_address"`
	Status            string `json:"status"`
	OrderDate         string `json:"order_date"`
}

func (r *DeliveryPostgres) GetActiveOrdersFromDB(Orders *[]Order) error {
	db := ConnectDB()
	defer db.Close()

	insertValue := `Select * from delivery where status = 'ready to delivery'`
	get, err := db.Query(insertValue)
	if err != nil {
		log.Println("Error with getting list of orders: " + err.Error())
		return err
	}

	for get.Next() {
		var order Order
		err = get.Scan(&order.IdDeliveryService, &order.Id, &order.IdCourier, &order.DeliveryTime, &order.CustomerAddress, &order.Status, &order.OrderDate)
		*Orders = append(*Orders, order)
	}
	return nil
}

func (r *DeliveryPostgres) GetActiveOrderFromDB(Orders *Order, id int) error {
	db := ConnectDB()
	defer db.Close()

	insertValue := `Select * from delivery where id = $1 AND status = 'ready to delivery'`
	get, err := db.Query(insertValue, id)
	if err != nil {
		log.Println("Error with getting order by id: " + err.Error())
		return err
	}

	for get.Next() {
		var order Order
		err = get.Scan(&order.IdDeliveryService, &order.Id, &order.IdCourier, &order.DeliveryTime, &order.CustomerAddress, &order.Status, &order.OrderDate)
		*Orders = order
	}
	return nil
}

func (r *DeliveryPostgres) ChangeOrderStatusInDB(id uint16) (uint16, error) {
	db := ConnectDB()
	defer db.Close()

	UpdateValue := `UPDATE "delivery" SET "status" = $1 WHERE "id" = $2`
	_, err := db.Exec(UpdateValue, "completed", id)
	if err != nil {
		log.Println("Error with getting order by id: " + err.Error())
		return 0, fmt.Errorf("updateOrder: error while scanning for order:%w", err)
	}
	return id, nil
}
