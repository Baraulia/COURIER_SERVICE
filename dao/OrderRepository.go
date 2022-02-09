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

func NewOrderPostgres(db *sql.DB) *OrderPostgres {
	return &OrderPostgres{db: db}
}

type Order struct {
	IdDeliveryService int `json:"id_delivery_service,omitempty"`
	IdCourier int `json:"id_courier,omitempty"`
	IdOrder int `json:"id_order"`
	DeliveryTime string `json:"delivery_time,omitempty"`
	CustomerAddress string `json:"customer_address,omitempty"`
	Status string `json:"status"`
	OrderDate time.Time `json:"order_date"`
}

func (r *OrderPostgres) GetCourierCompletedOrdersWithPage_fromDB(Orders *[]Order,limit,page,idCourier int) int{
	db:=OpenDB()
	defer db.Close()  //"Select id_courier,id_order, id_delivery_service,delivery_time,status, customer_address from delivery where status =`new` and status =`in progress` and status =`reade to delivery`")

	res,err:=db.Query(fmt.Sprintf("SELECT courier_id,id,delivery_service_id,delivery_time,status,customer_address FROM delivery WHERE status='completed' and courier_id=%d LIMIT %d OFFSET %d",idCourier,limit,limit*(page-1)))
		if err!=nil{
		panic(err)
	}
	for res.Next(){
		var order Order
		err = res.Scan(&order.IdCourier, &order.IdOrder, &order.IdDeliveryService, &order.DeliveryTime, &order.Status, &order.CustomerAddress)
		if err!=nil{
			panic(err)
		}

		*Orders = append (*Orders, order)
	}
	return len(*Orders)
}

func (r *OrderPostgres)  GetAllOrdersOfCourierServiceWithPage_fromDB(Orders *[]Order,limit,page,idService int) int{
	db:=OpenDB()
	defer db.Close()
	res,err:=db.Query(fmt.Sprintf("SELECT courier_id,id,delivery_time,status,customer_address FROM delivery WHERE delivery_service_id=%d LIMIT %d OFFSET %d",idService,limit,limit*(page-1)))
	if err!=nil{
		panic(err)
	}
	for res.Next(){
		var order Order
		err = res.Scan(&order.IdCourier, &order.IdOrder, &order.DeliveryTime, &order.Status, &order.CustomerAddress)
		if err!=nil{
			panic(err)
		}

		*Orders = append (*Orders, order)
	}
	return len(*Orders)
}


func (r *OrderPostgres)  GetCourierCompletedOrdersByMouthWithPage_fromDB(Orders *[]Order,limit,page,idCourier,Month int) int{
	db:=OpenDB()
	log.Println("connected to db")

	defer db.Close()
	res,err:=db.Query(fmt.Sprintf("SELECT courier_id ,id ,delivery_service_id ,delivery_time ,order_date ,status ,customer_address FROM delivery where courier_id=%d and Extract(MONTH from order_date )=%d LIMIT %d OFFSET %d ",idCourier,Month,limit,limit*(page-1)))
	if err!=nil{
		panic(err)
	}
	for res.Next(){
		var order Order
		err = res.Scan(&order.IdCourier, &order.IdOrder,&order.IdDeliveryService, &order.DeliveryTime,&order.OrderDate, &order.Status, &order.CustomerAddress)
		if err!=nil{
			panic(err)
		}

		*Orders = append (*Orders, order)
	}
	return len(*Orders)
}