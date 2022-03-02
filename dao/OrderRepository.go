package dao

import (
	"database/sql"
	"fmt"
	courierProto "github.com/Baraulia/COURIER_SERVICE/GRPCserver"
	"github.com/sirupsen/logrus"
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

func (r *OrderPostgres) GetActiveOrdersFromDB(id int) ([]Order, error) {
	var Orders []Order

	insertValue := `Select delivery_service_id,id,courier_id,delivery_time,customer_address,status,order_date,restaurant_address,picked from delivery where courier_id = $1 and status = 'ready to delivery'`
	get, err := r.db.Query(insertValue, id)
	if err != nil {
		log.Println("Error with getting list of orders: " + err.Error())
		return nil, err
	}

	for get.Next() {
		var order Order
		err = get.Scan(&order.IdDeliveryService, &order.Id, &order.IdCourier, &order.DeliveryTime, &order.CustomerAddress, &order.Status, &order.OrderDate, &order.RestaurantAddress, &order.Picked)
		Orders = append(Orders, order)
	}
	return Orders, nil
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
		err = res.Scan(&order.IdCourier, &order.Id, &order.DeliveryTime, &order.Status, &order.CustomerAddress)
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

func (r *OrderPostgres) GetCourierCompletedOrdersByMouthWithPage_fromDB(limit, page, idCourier, Month, Year int) ([]Order, int) {
	var Orders []Order
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
		var order Order
		err = res.Scan(&order.IdCourier, &order.Id, &order.IdDeliveryService, &order.DeliveryTime, &order.OrderDate, &order.Status, &order.CustomerAddress, &order.RestaurantAddress)
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

func (r *OrderPostgres) CreateOrder(order *courierProto.OrderCourierServer) error {
	_, err := r.db.Exec("INSERT INTO delivery (delivery_service_id, customer_address, restraunt_address, delivery_time, restaurant_name) VALUES ($1, $2, $3, $4, $5)", order.CourierServiceID, order.ClientAddress, order.RestaurantAddress, order.DeliveryTime, order.RestaurantName)
	if err != nil {
		logrus.Errorf("CreateOrder:%s", err)
		return fmt.Errorf("CreateOrder:%w", err)
	}
	return nil
}

func (r *OrderPostgres) GetServices() (*courierProto.ServiceResponse, error) {
	var Services *courierProto.ServiceResponse
	insertValue := `SELECT id, name, email, photo, description, phone_number, manager_id, status FROM delivery_service`
	get, err := r.db.Query(insertValue)
	if err != nil {
		logrus.Fatalln("Error of getting list of services :", err)
		return nil, fmt.Errorf("Error of getting list of services :%w", err)
	}

	for get.Next() {
		var service *courierProto.DeliveryService
		err = get.Scan(&service.ServiceId, &service.ServiceName, &service.ServiceEmail, &service.ServicePhoto, &service.ServiceDescription, &service.ServicePhone, &service.ServiceManagerId, &service.ServiceStatus)
		Services.Services = append(Services.Services, service)
	}
	return Services, nil
}
