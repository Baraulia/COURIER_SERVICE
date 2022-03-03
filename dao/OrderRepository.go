package dao

import (
	"database/sql"
	"fmt"
	"log"
	"time"
)

type OrderPostgres struct {
	db *sql.DB
}

func NewDeliveryPostgres(db *sql.DB) *OrderPostgres {
	return &OrderPostgres{db: db}
}

type Order struct {
	IdDeliveryService int       `json:"delivery_service_id,omitempty"`
	Id                int       `json:"id"`
	IdCourier         int       `json:"courier_id,omitempty"`
	DeliveryTime      time.Time `json:"delivery_time,omitempty"`
	CustomerAddress   string    `json:"customer_address,omitempty"`
	Status            string    `json:"status"`
	OrderDate         string    `json:"order_date"`
	RestaurantAddress string    `json:"restaurant_address"`
	Picked            bool      `json:"picked"`
}

type DetailedOrder struct {
	IdDeliveryService  int       `json:"delivery_service_id,omitempty"`
	IdOrder            int       `json:"id"`
	IdCourier          int       `json:"courier_id,omitempty"`
	DeliveryTime       time.Time `json:"delivery_time,omitempty"`
	CustomerAddress    string    `json:"customer_address,omitempty"`
	Status             string    `json:"status"`
	OrderDate          string    `json:"order_date,omitempty"`
	RestaurantAddress  string    `json:"restaurant_address,omitempty"`
	Picked             bool      `json:"picked"`
	CourierName        string    `json:"name"`
	CourierPhoneNumber string    `json:"phone_number"`
}

func (r *OrderPostgres) GetActiveOrderFromDB(id int) (Order, error) {
	var Ord Order

	insertValue := `Select delivery_service_id,id,courier_id,delivery_time,customer_address,status,order_date,restaurant_address,picked from delivery where id = $1 AND status = 'ready to delivery'`
	get, err := r.db.Query(insertValue, id)
	if err != nil {
		log.Println("Error with getting order by id: " + err.Error())
		return Order{}, err
	}

	for get.Next() {
		var order Order
		err = get.Scan(&order.IdDeliveryService, &order.Id, &order.IdCourier, &order.DeliveryTime, &order.CustomerAddress, &order.Status, &order.OrderDate, &order.RestaurantAddress, &order.Picked)
		Ord = order
	}
	return Ord, nil
}

func (r *OrderPostgres) ChangeOrderStatusInDB(id uint16) (uint16, error) {

	UpdateValue := `UPDATE "delivery" SET "status" = $1 WHERE "id" = $2`
	_, err := r.db.Exec(UpdateValue, "completed", id)
	if err != nil {
		log.Println("Error with getting order by id: " + err.Error())
		return 0, fmt.Errorf("updateOrder: error while scanning for order:%w", err)
	}
	return id, nil
}
