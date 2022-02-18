package dao

import (
	"database/sql"
	"fmt"
	"log"
)

type OrderPostgres struct {
	db *sql.DB
}

func NewOrderPostgres(db *sql.DB) *OrderPostgres {
	return &OrderPostgres{db: db}
}

type Order struct {
	IdDeliveryService int    `json:"delivery_service_id,omitempty"`
	IdOrder           int    `json:"id"`
	IdCourier         int    `json:"courier_id,omitempty"`
	DeliveryTime      string `json:"delivery_time,omitempty"`
	CustomerAddress   string `json:"customer_address,omitempty"`
	Status            string `json:"status"`
	OrderDate         string `json:"order_date"`
	RestaurantAddress string `json:"restaurant_address"`
	Picked            bool   `json:"picked"`
}

func (r *OrderPostgres) AssigningOrderToCourier_InDB(order Order) error {
	transaction, err := r.db.Begin()
	if err != nil {
		log.Fatal(err)
	}
	log.Println("connected to db")

	defer transaction.Commit()
	s := fmt.Sprintf("UPDATE delivery SET courier_id = %d WHERE id = %d", order.IdCourier, order.IdOrder)
	log.Println(s)
	insert, err1 := transaction.Query(s)
	if err1 != nil {
		log.Println(err1)
	}
	defer insert.Close()

	return nil
}
