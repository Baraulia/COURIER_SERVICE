package controller

import (
	"CourierService/dao"
	"CourierService/model"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

     var Orders []dao.Order
func GetCourierCompletedOrders(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	page,er:= strconv.Atoi(r.URL.Query().Get("page"))
	if er!=nil{
		b, _ := json.Marshal(fmt.Sprintf("Error: %s",er))
		w.Write(b)
	}
	limit,er:= strconv.Atoi(r.URL.Query().Get("limit"))
	if er!=nil{
		b, _ := json.Marshal(fmt.Sprintf("Error: %s",er))
		w.Write(b)
	}
	idCourier,er:= strconv.Atoi(r.URL.Query().Get("idcourier"))
	if er!=nil{
		b, _ := json.Marshal(fmt.Sprintf("Error: %s",er))
		w.Write(b)
	}
	if er==nil && page!=0 && limit!=0{
		Orders,err := model.GetCourierCompletedOrders(Orders,limit,page,idCourier)
		if err!=nil{
			b, _ := json.Marshal(fmt.Sprintf("Error: %s",err))
			w.Write(b)
		}
	json.NewEncoder(w).Encode(Orders) }else {
		b, _ := json.Marshal(fmt.Sprintf("Error: %s",er))
		w.Write(b)
	}
}


func GetAllOrdersOfCourierService(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	page,er:= strconv.Atoi(r.URL.Query().Get("page"))
	if er!=nil{
		b, _ := json.Marshal(fmt.Sprintf("Error: %s",er))
		w.Write(b)
	}
	limit,er:= strconv.Atoi(r.URL.Query().Get("limit"))
	if er!=nil{
		b, _ := json.Marshal(fmt.Sprintf("Error: %s",er))
		w.Write(b)
	}
	idService,er:= strconv.Atoi(r.URL.Query().Get("iddeliveryservice"))
	if er!=nil{
		b, _ := json.Marshal(fmt.Sprintf("Error: %s",er))
		w.Write(b)
	}
	if er==nil && page!=0 && limit!=0{
		Orders, err := model.GetAllOrdersOfCourierService(Orders, limit, page, idService)
		if err != nil {
			b, _ := json.Marshal(fmt.Sprintf("Error: %s", err))
			w.Write(b)
		}
		json.NewEncoder(w).Encode(Orders)
	}else {
		b, _ := json.Marshal(fmt.Sprintf("Error: %s",er))
		w.Write(b)}
}
//SELECT id_courier,id_order,id_delivery_service,delivery_time,order_date, status, customer_address FROM delivery where id_courier=1 and status='completed' and Extract(MONTH from order_date )=1;
//SELECT id_courier,id_order,id_delivery_service,delivery_time,order_date, customer_address FROM delivery where id_courier=1 and Extract(MONTH from order_date )=1;

func GetCourierCompletedOrdersByMonth(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	page,er:= strconv.Atoi(r.URL.Query().Get("page"))
	if er!=nil{
		b, _ := json.Marshal(fmt.Sprintf("Error: %s",er))
		w.Write(b)
	}
	limit,er:= strconv.Atoi(r.URL.Query().Get("limit"))
	if er!=nil{
		b, _ := json.Marshal(fmt.Sprintf("Error: %s",er))
		w.Write(b)
	}
	idCourier,er:= strconv.Atoi(r.URL.Query().Get("idcourier"))
	if er!=nil{
		b, _ := json.Marshal(fmt.Sprintf("Error: %s",er))
		w.Write(b)
	}
	Month,er:= strconv.Atoi(r.URL.Query().Get("month"))
	if er!=nil {
		b, _ := json.Marshal(fmt.Sprintf("Error: %s",er))
		w.Write(b)
	}
	fmt.Println(limit,page,idCourier,Month)
	if er==nil && Month!=0 && page!=0 && limit!=0{
		Orders,err := model.GetCourierCompletedOrdersByMonth(Orders,limit,page,idCourier,Month)
		if err!=nil{
			b, _ := json.Marshal(fmt.Sprintf("Error: %s",err))
			w.Write(b)
		}
		json.NewEncoder(w).Encode(Orders) }else {
		b, _ := json.Marshal(fmt.Sprintf("Error: %s",er))
		w.Write(b)
	}
}