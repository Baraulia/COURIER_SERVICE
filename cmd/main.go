package main

import (
	"github.com/Baraulia/COURIER_SERVICE/controller"
	"github.com/Baraulia/COURIER_SERVICE/dao"
	"github.com/Baraulia/COURIER_SERVICE/service"
	_ "github.com/lib/pq"
	"log"
	"net/http"
)

// @title Courier Service
// @description Courier Service for Food Delivery Application
func main() {
	log.Println("Start...")
	database, err := dao.NewPostgresDB(dao.PostgresDB{
		"159.223.1.135",
		"5434",
		"courierteam1",
		"qwerty",
		"courier_db",
		"disable"})
	if err != nil {
		log.Fatal("failed to initialize dao:", err.Error())
	}
	repository := dao.NewRepository(database)
	services := service.NewService(repository)
	handlers := controller.NewHandler(services)
	s := &http.Server{
		Addr:    ":8080", // ":"+host,
		Handler: handlers.InitRoutesGin(),
	}
	err = s.ListenAndServe()
	if err != nil {
		log.Println("failed to initialize port:", err.Error())
	}
}
