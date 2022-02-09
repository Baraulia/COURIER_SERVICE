package controller

import (
	"CourierService/dao"
	"CourierService/model"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

func JSONError(w http.ResponseWriter, err interface{}, code int) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.Header().Set("X-Content-Type-Options", "nosniff")
	w.WriteHeader(code)
	json.NewEncoder(w).Encode(err)
}

     var Orders []dao.Order
func GetCourierCompletedOrders(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	page,er:= strconv.Atoi(r.URL.Query().Get("page"))
	if er!=nil{
		b, _ := json.Marshal(fmt.Sprintf("Error: %s",er))
		JSONError(w,b,http.StatusNoContent)
	}
	limit,er:= strconv.Atoi(r.URL.Query().Get("limit"))
	if er!=nil{
		b, _ := json.Marshal(fmt.Sprintf("Error: %s",er))
		JSONError(w,b,http.StatusNoContent)
	}
	idCourier,er:= strconv.Atoi(r.URL.Query().Get("idcourier"))
	if er!=nil{
		b, _ := json.Marshal(fmt.Sprintf("Error: %s",er))
		JSONError(w,b,http.StatusNoContent)
	}
	if er==nil && page!=0 && limit!=0{
		Orders,err := model.GetCourierCompletedOrders(limit,page,idCourier)
		if err!=nil{
			b, _ := json.Marshal(fmt.Sprintf("Error: %s",err))
			JSONError(w,b,http.StatusExpectationFailed)
		}
	json.NewEncoder(w).Encode(Orders) }else {
		b, _ := json.Marshal(fmt.Sprintf("Error: %s",er))
		JSONError(w,b,http.StatusExpectationFailed)
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
	if er==nil && page!=0 && limit!=0 {
		Orders, err := model.GetAllOrdersOfCourierService( limit, page, idService)
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
		Orders,err := model.GetCourierCompletedOrdersByMonth(	limit,page,idCourier,Month)
		if err!=nil{
			b, _ := json.Marshal(fmt.Sprintf("Error: %s",err))
			w.Write(b)
		}
		json.NewEncoder(w).Encode(Orders) }else {
		b, _ := json.Marshal(fmt.Sprintf("Error: %s",er))
		w.Write(b)
	}
}