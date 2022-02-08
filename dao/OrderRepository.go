package dao

import (
	"fmt"
	"log"
	"time"
)

type Order struct {
	IdDeliveryService int `json:"id_delivery_service,omitempty"`
	IdCourier int `json:"id_courier,omitempty"`
	IdOrder int `json:"id_order"`
	DeliveryTime string `json:"delivery_time,omitempty"`
	CustomerAddress string `json:"customer_address,omitempty"`
	Status string `json:"status"`
	OrderDate time.Time `json:"order_date"`
}

func GetCourierCompletedOrdersWithPage_fromDB(Orders *[]Order,limit,page,idCourier int) int{
	db:=OpenDB()
	defer db.Close()  //"Select id_courier,id_order, id_delivery_service,delivery_time,status, customer_address from delivery where status =`new` and status =`in progress` and status =`reade to delivery`")

	res,err:=db.Query(fmt.Sprintf("SELECT id_courier,id_order,id_delivery_service,delivery_time,status,customer_address FROM delivery WHERE status='completed' and id_courier=%d LIMIT %d OFFSET %d",idCourier,limit,limit*(page-1)))
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

func GetAllOrdersOfCourierServiceWithPage_fromDB(Orders *[]Order,limit,page,idService int) int{
	db:=OpenDB()
	defer db.Close()
	res,err:=db.Query(fmt.Sprintf("SELECT id_courier,id_order,delivery_time,status,customer_address FROM delivery WHERE id_delivery_service=%d LIMIT %d OFFSET %d",idService,limit,limit*(page-1)))
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


func GetCourierCompletedOrdersByMouthWithPage_fromDB(Orders *[]Order,limit,page,idCourier,Month int) int{
	db:=OpenDB()
	log.Println("connected to db")

	defer db.Close()
	res,err:=db.Query(fmt.Sprintf("SELECT id_courier ,id_order ,id_delivery_service ,delivery_time ,order_date ,status ,customer_address FROM delivery where id_courier=%d and Extract(MONTH from order_date )=%d LIMIT %d OFFSET %d ",idCourier,Month,limit,limit*(page-1)))
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