package main

import (
	_ "github.com/lib/pq"
	"log"
	"net/http"
	"os"
	"stlab.itechart-group.com/go/food_delivery/courier_service/controller"
	"stlab.itechart-group.com/go/food_delivery/courier_service/dao"
	"stlab.itechart-group.com/go/food_delivery/courier_service/model"
)

// @title COURIER_SERVICE
// @description API Server
func main() {
	log.Println("Start...")
	db, err := dao.NewPostgresDB(dao.PostgresDB{
		"159.223.1.135",
		"5434",
		"courierteam1",
		"qwerty",
		"courier_db",
		"disable"})
	if err != nil {
		log.Fatal("failed to initialize db:", err.Error())
	}
	repos := dao.NewRepository(db)
	services := model.NewService(repos)
	handlers := controller.NewHandler(services)
	host := os.Getenv("API_SERVER_PORT")
	s := &http.Server{
		Addr:    ":" + host,
		Handler: handlers.InitRoutesGin(),
	}
	err = s.ListenAndServe()
	if err != nil {
		log.Println("failed to initialize port:", err.Error())
	}

}
