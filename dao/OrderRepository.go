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
	Status             string `json:"status"`
	OrderDate          string `json:"order_date,omitempty"`
	RestaurantAddress  string `json:"restaurant_address,omitempty"`
	Picked             bool   `json:"picked"`
	CourierName        string `json:"name"`
	CourierPhoneNumber string `json:"phone_number"`
}

func (r *OrderPostgres) GetCourierCompletedOrdersWithPage_fromDB(limit, page, idCourier int) ([]DetailedOrder, int) {
	var Orders []DetailedOrder
	transaction, err := r.db.Begin()
	if err != nil {
		log.Fatal(err)
	}
	defer transaction.Commit()
	res, err := transaction.Query(fmt.Sprintf("SELECT delivery.order_date, delivery.courier_id,delivery.id,delivery.delivery_service_id,delivery.delivery_time,delivery.status,delivery.customer_address,delivery.restaurant_address,couriers.name,couriers.phone_number FROM delivery JOIN couriers ON couriers.id_courier=delivery.courier_id Where delivery.status='completed' and delivery.courier_id=%d LIMIT %d OFFSET %d", idCourier, limit, limit*(page-1)))
	if err != nil {
		log.Fatal(err)
	}
	for res.Next() {
		var order DetailedOrder
		err = res.Scan(&order.OrderDate, &order.IdCourier, &order.IdOrder, &order.IdDeliveryService, &order.DeliveryTime, &order.Status, &order.CustomerAddress, &order.RestaurantAddress, &order.CourierName, &order.CourierPhoneNumber)
		if err != nil {
			panic(err)
		}
		Orders = append(Orders, order)
	}

	var Ordersss []Order
	resl, err := transaction.Query(fmt.Sprintf("SELECT courier_id FROM delivery WHERE status='completed' and courier_id=%d ", idCourier))
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

func (r *OrderPostgres) GetAllOrdersOfCourierServiceWithPage_fromDB(limit, page, idService int) ([]Order, int) {
	var Orders []Order
	transaction, err := r.db.Begin()
	if err != nil {
		log.Fatal(err)
	}
	defer transaction.Commit()
	res, err := transaction.Query(fmt.Sprintf("SELECT courier_id,id,delivery_time,status,customer_address FROM delivery WHERE delivery_service_id=%d LIMIT %d OFFSET %d", idService, limit, limit*(page-1)))
	if err != nil {
		panic(err)
	}
	for res.Next() {
		var order Order
		err = res.Scan(&order.IdCourier, &order.IdOrder, &order.DeliveryTime, &order.Status, &order.CustomerAddress)
		if err != nil {
			panic(err)
		}

		Orders = append(Orders, order)
	}

	var Ordersss []Order
	resl, err := transaction.Query(fmt.Sprintf("SELECT courier_id FROM delivery WHERE delivery_service_id=%d ", idService))
	if err != nil {
		panic(err)
	}
	for resl.Next() {
		var order Order
		err = resl.Scan(&order.IdCourier)
		if err != nil {
			panic(err)
		}

		Ordersss = append(Ordersss, order)
	}
	return Orders, len(Ordersss)
}

func (r *OrderPostgres) GetCourierCompletedOrdersByMouthWithPage_fromDB(limit, page, idCourier, Month int) ([]Order, int) {
	var Orders []Order
	transaction, err := r.db.Begin()
	if err != nil {
		log.Fatal(err)
	}
	log.Println("connected to db")

	defer transaction.Commit()
	res, err := transaction.Query(fmt.Sprintf("SELECT courier_id ,id ,delivery_service_id ,delivery_time ,order_date ,status ,customer_address, restaurant_address FROM delivery where courier_id=%d and Extract(MONTH from order_date )=%d LIMIT %d OFFSET %d ", idCourier, Month, limit, limit*(page-1)))
	if err != nil {
		panic(err)
	}
	for res.Next() {
		var order Order
		err = res.Scan(&order.IdCourier, &order.IdOrder, &order.IdDeliveryService, &order.DeliveryTime, &order.OrderDate, &order.Status, &order.CustomerAddress, &order.RestaurantAddress)
		if err != nil {
			panic(err)
		}

		Orders = append(Orders, order)
	}
	var Ordersss []Order
	resl, err := transaction.Query(fmt.Sprintf("SELECT courier_id FROM delivery WHERE courier_id=%d and Extract(MONTH from order_date )=%d", idCourier, Month))
	if err != nil {
		panic(err)
	}
	for resl.Next() {
		var order Order
		err = resl.Scan(&order.IdCourier)
		if err != nil {
			panic(err)
		}

		Ordersss = append(Ordersss, order)
	}

	return Orders, len(Ordersss)
}
