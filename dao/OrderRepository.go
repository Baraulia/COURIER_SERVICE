package dao

import (
	"database/sql"
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
	Id                int    `json:"id"`
	IdCourier         int    `json:"courier_id,omitempty"`
	DeliveryTime      string `json:"delivery_time,omitempty"`
	CustomerAddress   string `json:"customer_address,omitempty"`
	Status            string `json:"status"`
	OrderDate         string `json:"order_date"`
	RestaurantAddress string `json:"restaurant_address"`
	Picked            bool   `json:"picked"`
}

type DetailedOrder struct {
	IdDeliveryService  int    `json:"delivery_service_id,omitempty"`
	Id                 int    `json:"id"`
	IdCourier          int    `json:"courier_id,omitempty"`
	DeliveryTime       string `json:"delivery_time,omitempty"`
	CustomerAddress    string `json:"customer_address,omitempty"`
	Status             string `json:"status,omitempty"`
	OrderDate          string `json:"order_date,omitempty"`
	RestaurantAddress  string `json:"restaurant_address,omitempty"`
	Picked             bool   `json:"picked,omitempty"`
	CourierName        string `json:"name,omitempty"`
	CourierPhoneNumber string `json:"phone_number,omitempty"`
}

func (r *OrderPostgres) AssigningOrderToCourierInDB(order Order) error {
	log.Println("connected to db")
	s := "UPDATE delivery SET courier_id = $1 WHERE id = $2"
	log.Println(s)
	insert, err := r.db.Query(s, order.IdCourier, order.Id)
	defer insert.Close()
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}
