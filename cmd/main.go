package main

import (
	"github.com/Baraulia/COURIER_SERVICE/Controller"
	_ "github.com/Baraulia/COURIER_SERVICE/dao"
	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
	"net/http"
)

func main(){
	r := mux.NewRouter()
	r.HandleFunc("/orders/completed", controller.GetCourierCompletedOrders).Methods("GET")
	r.HandleFunc("/orders", controller.GetAllOrdersOfCourierService).Methods("GET")
	r.HandleFunc("/orders/bymonth",controller.GetCourierCompletedOrdersByMonth).Methods("GET")
	http.ListenAndServe("127.0.0.1:8080", r)

}

