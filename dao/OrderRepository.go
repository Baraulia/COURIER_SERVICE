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
	IdDeliveryService int `json:"delivery_service_id,omitempty"`
	IdOrder int `json:"id"`
	IdCourier int `json:"courier_id,omitempty"`
	DeliveryTime string `json:"delivery_time,omitempty"`
	CustomerAddress string `json:"customer_address,omitempty"`
	Status string `json:"status"`
	OrderDate string `json:"order_date"`
}

func (r *OrderPostgres) GetCourierCompletedOrdersWithPage_fromDB(Orders *[]Order,limit,page,idCourier int) (int){
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
	var Ordersss []Order
	resl,err:=db.Query(fmt.Sprintf("SELECT courier_id FROM delivery WHERE status='completed' and courier_id=%d ",idCourier))
	if err!=nil{
		panic(err)
	}
	for resl.Next(){
		var order Order
		err = resl.Scan(&order.IdCourier)
		if err!=nil{
			panic(err)
		}

		*Orders = append (*Orders, order)
	}
	return len(Ordersss)
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
		err = res.Scan(&order.IdCourier,&order.IdOrder,&order.DeliveryTime,&order.Status,&order.CustomerAddress)
		if err!=nil{
			panic(err)
		}

		*Orders = append (*Orders, order)
	}

	var Ordersss []Order
	resl,err:=db.Query(fmt.Sprintf("SELECT courier_id FROM delivery WHERE delivery_service_id=%d ",idService))
	if err!=nil{
		panic(err)
	}
	for resl.Next(){
		var order Order
		err = resl.Scan(&order.IdCourier)
		if err!=nil{
			panic(err)
		}

		*Orders = append (*Orders, order)
	}
	return len(Ordersss)
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
	var Ordersss []Order
	resl,err:=db.Query(fmt.Sprintf("SELECT courier_id FROM delivery WHERE courier_id=%d and Extract(MONTH from order_date )=%d",idCourier,Month))
	if err!=nil{
		panic(err)
	}
	for resl.Next(){
		var order Order
		err = resl.Scan(&order.IdCourier)
		if err!=nil{
			panic(err)
		}

		*Orders = append (*Orders, order)
	}

	return len(Ordersss)
}