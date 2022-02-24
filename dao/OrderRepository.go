package dao

import (
	"database/sql"
	"fmt"
	"github.com/Baraulia/COURIER_SERVICE/model"
	"log"
)

type OrderPostgres struct {
	db *sql.DB
}

func NewDeliveryPostgres(db *sql.DB) *OrderPostgres {
	return &OrderPostgres{db: db}
}

func (r *OrderPostgres) GetActiveOrdersFromDB(id int) ([]model.Order, error) {
	var Orders []model.Order

	insertValue := `Select delivery_service_id,id,courier_id,delivery_time,customer_address,status,order_date,restaurant_address,picked from delivery where courier_id = $1 and status = 'ready to delivery'`
	get, err := r.db.Query(insertValue, id)
	if err != nil {
		log.Println("Error with getting list of orders: " + err.Error())
		return nil, err
	}

	for get.Next() {
		var order model.Order
		err = get.Scan(&order.IdDeliveryService, &order.Id, &order.IdCourier, &order.DeliveryTime, &order.CustomerAddress, &order.Status, &order.OrderDate, &order.RestaurantAddress, &order.Picked)
		Orders = append(Orders, order)
	}
	return Orders, nil
}

func (r *OrderPostgres) GetActiveOrderFromDB(id int) (*model.Order, error) {
	var Ord model.Order

	insertValue := `Select delivery_service_id,id,courier_id,delivery_time,customer_address,status,order_date,restaurant_address,picked from delivery where id = $1 AND status = 'ready_to_delivery'`
	get, err := r.db.Query(insertValue, id)
	if err != nil {
		log.Println("Error with getting order by id: " + err.Error())
		return nil, err
	}

	for get.Next() {
		var order model.Order
		err = get.Scan(&order.IdDeliveryService, &order.Id, &order.IdCourier, &order.DeliveryTime, &order.CustomerAddress, &order.Status, &order.OrderDate, &order.RestaurantAddress, &order.Picked)
		Ord = order
	}
	return &Ord, nil
}

func (r *OrderPostgres) ChangeOrderStatusInDB(status string, id uint16) (uint16, error) {
	UpdateValue := `UPDATE delivery SET status = $1 WHERE id = $2`
	_, err := r.db.Exec(UpdateValue, status, id)
	if err != nil {
		log.Println("Error with getting order by id: " + err.Error())
		return 0, fmt.Errorf("updateOrder: error while scanning for order:%w", err)
	}
	return id, nil
}

func (r *OrderPostgres) GetCourierCompletedOrdersWithPage_fromDB(limit, page, idCourier int) ([]model.DetailedOrder, int) {
	var Orders []model.DetailedOrder
	transaction, err := r.db.Begin()
	if err != nil {
		log.Fatal(err)
	}
	defer transaction.Commit()
	res, err := transaction.Query(fmt.Sprintf("SELECT delivery.order_date, delivery.courier_id,delivery.id, delivery.delivery_service_id, delivery.delivery_time, delivery.status, delivery.customer_address, delivery.restaurant_address, couriers.name, couriers.phone_number FROM delivery JOIN couriers ON couriers.id_courier=delivery.courier_id WHERE delivery.status='completed' and delivery.courier_id=%d LIMIT %d OFFSET %d", idCourier, limit, limit*(page-1)))
	if err != nil {
		log.Fatal(err)
	}
	for res.Next() {
		var order model.DetailedOrder
		err = res.Scan(&order.OrderDate, &order.IdCourier, &order.IdOrder, &order.IdDeliveryService, &order.DeliveryTime, &order.Status, &order.CustomerAddress, &order.RestaurantAddress, &order.CourierName, &order.CourierPhoneNumber)
		if err != nil {
			panic(err)
		}
		Orders = append(Orders, order)
	}

	var Ordersss []model.Order
	resl, err := transaction.Query(fmt.Sprintf("SELECT courier_id FROM delivery WHERE status='completed' and courier_id=%d ", idCourier))
	if err != nil {
		log.Println(err)
	}
	for resl.Next() {
		var order1 model.Order
		err = resl.Scan(&order1.IdCourier)
		if err != nil {
			panic(err)
		}

		Ordersss = append(Ordersss, order1)
	}
	return Orders, len(Ordersss)
}

func (r *OrderPostgres) GetAllOrdersOfCourierServiceWithPage_fromDB(limit, page, idService int) ([]model.Order, int) {
	var Orders []model.Order
	transaction, err := r.db.Begin()
	if err != nil {
		log.Fatal(err)
	}
	defer transaction.Commit()
	res, err := transaction.Query(fmt.Sprintf("SELECT courier_id, id, delivery_time, status, customer_address FROM delivery WHERE delivery_service_id=%d LIMIT %d OFFSET %d", idService, limit, limit*(page-1)))
	if err != nil {
		panic(err)
	}
	for res.Next() {
		var order model.Order
		err = res.Scan(&order.IdCourier, &order.Id, &order.DeliveryTime, &order.Status, &order.CustomerAddress)
		if err != nil {
			panic(err)
		}

		Orders = append(Orders, order)
	}

	var Ordersss []model.Order
	resl, err := transaction.Query(fmt.Sprintf("SELECT courier_id FROM delivery WHERE delivery_service_id=%d ", idService))
	if err != nil {
		panic(err)
	}
	for resl.Next() {
		var order model.Order
		err = resl.Scan(&order.IdCourier)
		if err != nil {
			panic(err)
		}

		Ordersss = append(Ordersss, order)
	}
	return Orders, len(Ordersss)
}

func (r *OrderPostgres) GetCourierCompletedOrdersByMouthWithPage_fromDB(limit, page, idCourier, Month, Year int) ([]model.Order, int) {
	var Orders []model.Order
	transaction, err := r.db.Begin()
	if err != nil {
		log.Fatal(err)
	}
	log.Println("connected to db")

	defer transaction.Commit()
	res, err := transaction.Query(fmt.Sprintf("SELECT courier_id ,id ,delivery_service_id ,delivery_time ,order_date ,status ,customer_address, restaurant_address FROM delivery where courier_id=%d and Extract(MONTH from order_date )=%d and Extract(Year from order_date )=%d LIMIT %d OFFSET %d ", idCourier, Month, Year, limit, limit*(page-1)))
	if err != nil {
		panic(err)
	}
	for res.Next() {
		var order model.Order
		err = res.Scan(&order.IdCourier, &order.Id, &order.IdDeliveryService, &order.DeliveryTime, &order.OrderDate, &order.Status, &order.CustomerAddress, &order.RestaurantAddress)
		if err != nil {
			panic(err)
		}

		Orders = append(Orders, order)
	}
	var Ordersss []model.Order
	resl, err := transaction.Query(fmt.Sprintf("SELECT courier_id FROM delivery WHERE courier_id=%d and Extract(MONTH from order_date )=%d", idCourier, Month))
	if err != nil {
		panic(err)
	}
	for resl.Next() {
		var order model.Order
		err = resl.Scan(&order.IdCourier)
		if err != nil {
			panic(err)
		}

		Ordersss = append(Ordersss, order)
	}

	return Orders, len(Ordersss)
}
