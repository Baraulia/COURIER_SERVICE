package main

import (
	"database/sql"
	"github.com/Baraulia/COURIER_SERVICE/controller"
	"github.com/Baraulia/COURIER_SERVICE/dao"
	_ "github.com/Baraulia/COURIER_SERVICE/dao"
	"github.com/Baraulia/COURIER_SERVICE/model"
	_ "github.com/lib/pq"
	"log"
	"net/http"
)

func main(){
	log.Println("Start...")
	var db *sql.DB
	repos := dao.NewRepository(db)
	services := model.NewService(repos)
	handlers := controller.NewHandler(services)
	s := &http.Server{
		Addr:    ":8080",
		Handler: handlers.InitRoutes(),
	}
	s.ListenAndServe()
/*	r := mux.NewRouter()
	r.HandleFunc("/orders/completed", controller.GetCourierCompletedOrders).Methods("GET")
	r.HandleFunc("/orders", controller.GetAllOrdersOfCourierService).Methods("GET")
	r.HandleFunc("/orders/bymonth",controller.GetCourierCompletedOrdersByMonth).Methods("GET")
	http.ListenAndServe("127.0.0.1:8080", r)   */

}

