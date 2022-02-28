package dao

import (
	"database/sql"
	"fmt"
	courierProto "github.com/Baraulia/COURIER_SERVICE/GRPCserver"
	"github.com/Baraulia/COURIER_SERVICE/model"
	"github.com/sirupsen/logrus"
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
		logrus.Fatalln("Error of getting list of active orders :", err)
		return nil, fmt.Errorf("Error of getting list of active orders :%w", err)
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
		logrus.Fatalln("Error with getting order by id :", err)
		return nil, fmt.Errorf("Error with getting order by id :%w", err)
	}

	for get.Next() {
		var order model.Order
		err = get.Scan(&order.IdDeliveryService, &order.Id, &order.IdCourier, &order.DeliveryTime, &order.CustomerAddress, &order.Status, &order.OrderDate, &order.RestaurantAddress, &order.Picked)
		Ord = order
	}
	return &Ord, nil
}

func (r *OrderPostgres) GetOrderStatusByID(id int) (*courierProto.OrderStatusResponse, error) {
	var Ord courierProto.OrderStatusResponse

	insertValue := `SELECT id, status from delivery where id = $1`
	get, err := r.db.Query(insertValue, id)
	if err != nil {
		logrus.Fatalln("Error with getting order by id :", err)
		return nil, fmt.Errorf("Error with getting order by id :%w", err)
	}

	for get.Next() {
		var order courierProto.OrderStatusResponse
		err = get.Scan(&order.OrderId, &order.Status)
		Ord = order
	}
	return &Ord, nil
}

func (r *OrderPostgres) ChangeOrderStatusInDB(status string, id uint16) (uint16, error) {
	UpdateValue := `UPDATE delivery SET status = $1 WHERE id = $2`
	_, err := r.db.Exec(UpdateValue, status, id)
	if err != nil {
		logrus.Fatalln("Error with getting order by id: " + err.Error())
		return 0, fmt.Errorf("updateOrder: error while scanning for order:%w", err)
	}
	return id, nil
}

func (r *OrderPostgres) GetCourierCompletedOrdersWithPageFromDB(limit, page, idCourier int) ([]model.DetailedOrder, int) {
	var Orders []model.DetailedOrder
	transaction, err := r.db.Begin()
	if err != nil {
		logrus.Fatalln("GetAllCompletedOrders: can not starts transaction:%s", err)
		return nil, 0
	}
	defer transaction.Commit()
	res, err := transaction.Query(fmt.Sprintf("SELECT delivery.order_date, delivery.courier_id,delivery.id, delivery.delivery_service_id, delivery.delivery_time, delivery.status, delivery.customer_address, delivery.restaurant_address, couriers.name, couriers.phone_number FROM delivery JOIN couriers ON couriers.id_courier=delivery.courier_id WHERE delivery.status='completed' and delivery.courier_id=%d LIMIT %d OFFSET %d", idCourier, limit, limit*(page-1)))
	if err != nil {
		logrus.Fatalln("GetAllCompletedOrders: can not executes a query:%s", err)
		return nil, 0
	}

	for res.Next() {
		var order model.DetailedOrder
		err1 := res.Scan(&order.OrderDate, &order.IdCourier, &order.IdOrder, &order.IdDeliveryService, &order.DeliveryTime, &order.Status, &order.CustomerAddress, &order.RestaurantAddress, &order.CourierName, &order.CourierPhoneNumber); if err1 != nil {
			logrus.Fatalln("Error while scanning for orders:%s", err1)
			return nil, 0
		}
		Orders = append(Orders, order)
	}

	var Ordersss []model.Order
	resl, err := transaction.Query(fmt.Sprintf("SELECT courier_id FROM delivery WHERE status='completed' and courier_id=%d ", idCourier))
	if err != nil {
		logrus.Fatalln("GetAllCompletedOrders: can not executes a query:%s", err)
		return nil, 0
	}
	for resl.Next() {
		var order1 model.Order
		err2 := resl.Scan(&order1.IdCourier); if err2 != nil {
			logrus.Fatalln("Error while scanning for orders:%s", err2)
			return nil, 0
		}
		Ordersss = append(Ordersss, order1)
	}
	return Orders, len(Ordersss)
}

func (r *OrderPostgres) GetAllOrdersOfCourierServiceWithPageFromDB(limit, page, idService int) ([]model.Order, int) {
	var Orders []model.Order
	transaction, err := r.db.Begin()
	if err != nil {
		logrus.Fatalln("GetAllOrders: can not starts transaction:%s", err)
		return nil, 0
	}
	defer transaction.Commit()
	res, err := transaction.Query(fmt.Sprintf("SELECT courier_id, id, delivery_time, status, customer_address FROM delivery WHERE delivery_service_id=%d LIMIT %d OFFSET %d", idService, limit, limit*(page-1)))
	if err != nil {
		logrus.Fatalln("GetAllOrders: can not executes a query:%s", err)
		return nil, 0
	}
	for res.Next() {
		var order model.Order
		err1 := res.Scan(&order.IdCourier, &order.Id, &order.DeliveryTime, &order.Status, &order.CustomerAddress);if err1 != nil {
			logrus.Fatalln("Error while scanning for orders:%s", err1)
			return nil, 0
		}
		Orders = append(Orders, order)
	}

	var Ordersss []model.Order
	resl, err := transaction.Query(fmt.Sprintf("SELECT courier_id FROM delivery WHERE delivery_service_id=%d ", idService))
	if err != nil {
		logrus.Fatalln("GetAllOrders: can not executes a query:%s", err)
		return nil, 0
	}
	for resl.Next() {
		var order model.Order
		err2 := resl.Scan(&order.IdCourier);if err2 != nil {
			logrus.Fatalln("Error while scanning for orders:%s", err2)
			return nil, 0
		}
		Ordersss = append(Ordersss, order)
	}
	return Orders, len(Ordersss)
}

func (r *OrderPostgres) GetCourierCompletedOrdersByMouthWithPageFromDB(limit, page, idCourier, Month, Year int) ([]model.Order, int) {
	var Orders []model.Order
	transaction, err := r.db.Begin()
	if err != nil {
		logrus.Fatalln("GetCourierCompletedOrdersByMouth: can not starts transaction:%s", err)
		return nil, 0
	}

	defer transaction.Commit()
	res, err := transaction.Query(fmt.Sprintf("SELECT courier_id ,id ,delivery_service_id ,delivery_time ,order_date ,status ,customer_address, restaurant_address FROM delivery where courier_id=%d and Extract(MONTH from order_date )=%d and Extract(Year from order_date )=%d LIMIT %d OFFSET %d ", idCourier, Month, Year, limit, limit*(page-1)))
	if err != nil {
		logrus.Fatalln("GetCourierCompletedOrdersByMouth: can not executes a query:%s", err)
		return nil, 0
	}
	for res.Next() {
		var order model.Order
		err1 := res.Scan(&order.IdCourier, &order.Id, &order.IdDeliveryService, &order.DeliveryTime, &order.OrderDate, &order.Status, &order.CustomerAddress, &order.RestaurantAddress);if err1 != nil {
			logrus.Fatalln("Error while scanning for orders:%s", err1)
			return nil, 0
		}
		Orders = append(Orders, order)
	}
	var Ordersss []model.Order
	resl, err := transaction.Query(fmt.Sprintf("SELECT courier_id FROM delivery WHERE courier_id=%d and Extract(MONTH from order_date )=%d", idCourier, Month))
	if err != nil {
		logrus.Fatalln("GetCourierCompletedOrdersByMouth: can not executes a query:%s", err)
		return nil, 0
	}
	for resl.Next() {
		var order model.Order
		err2 := resl.Scan(&order.IdCourier);if err2 != nil {
			logrus.Fatalln("Error while scanning for orders:%s", err2)
			return nil, 0
		}
		Ordersss = append(Ordersss, order)
	}
	return Orders, len(Ordersss)
}
