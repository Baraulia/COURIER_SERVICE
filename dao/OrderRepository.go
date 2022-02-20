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

type DetailedOrder struct {
	IdDeliveryService  int    `json:"delivery_service_id,omitempty"`
	IdOrder            int    `json:"id"`
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

func (r *OrderPostgres) AssigningOrderToCourier_InDB(order Order) error {
	transaction, err := r.db.Begin()
	if err != nil {
		log.Fatal(err)
	}
	log.Println("connected to db")

	defer transaction.Commit()
	s := fmt.Sprintf("UPDATE delivery SET courier_id = %d WHERE id = %d", order.IdCourier, order.IdOrder)
	log.Println(s)
	insert, err := transaction.Query(s)
	defer insert.Close()
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}

func(r *OrderPostgres) GetAllServiceCompletedOrders_fromDB(limit, page, idService int) ([]DetailedOrder, int){
	var Orders []DetailedOrder
	transaction, err := r.db.Begin()
	if err != nil {
		log.Fatal(err)
	}
	defer transaction.Commit()
	res, err := transaction.Query(fmt.Sprintf("SELECT delivery.order_date, delivery.courier_id,delivery.id,delivery.customer_address,delivery.restaurant_address,couriers.name,couriers.phone_number FROM delivery JOIN couriers ON couriers.id_courier=delivery.courier_id Where delivery.status='completed' and delivery.delivery_service_id=%d LIMIT %d OFFSET %d", idService, limit, limit*(page-1)))
	if err != nil {
		log.Fatal(err)
	}
	for res.Next() {
		var order DetailedOrder
		err = res.Scan(&order.OrderDate, &order.IdCourier, &order.IdOrder, &order.CustomerAddress, &order.RestaurantAddress, &order.CourierName, &order.CourierPhoneNumber)
		if err != nil {
			panic(err)
		}
		Orders = append(Orders, order)
	}

	var Ordersss []Order
	resl, err := transaction.Query(fmt.Sprintf("SELECT courier_id FROM delivery WHERE status='completed' and delivery_service_id=%d ", idService))
	if err != nil {
		log.Println(err)
	}
	for resl.Next() {
		var order1 Order
		err = resl.Scan(&order1.IdCourier)
		if err != nil {
			panic(err)
		}

		Ordersss = append(Ordersss, order1)
	}
	return Orders, len(Ordersss)
}

func(r *OrderPostgres) GetDetailedOrdersById_FromDB(idOrder int) (DetailedOrder, error){
	var order DetailedOrder
	transaction, err := r.db.Begin()
	if err != nil {
		log.Fatal(err)
	}
	defer transaction.Commit()
	res, err := transaction.Query(fmt.Sprintf("SELECT delivery.order_date, delivery.courier_id,delivery.id,delivery.delivery_service_id,delivery.delivery_time,delivery.status,delivery.customer_address,delivery.restaurant_address,couriers.name,couriers.phone_number FROM delivery JOIN couriers ON couriers.id_courier=delivery.courier_id Where delivery.status='completed' and delivery.id=%d", idOrder))
	if err != nil {
		log.Fatal(err)
	}
	for res.Next() {
		err = res.Scan(&order.OrderDate, &order.IdCourier, &order.IdOrder, &order.IdDeliveryService, &order.DeliveryTime, &order.Status, &order.CustomerAddress, &order.RestaurantAddress, &order.CourierName, &order.CourierPhoneNumber)
		if err != nil {
			panic(err)
		}
	}
	return order, nil
}

