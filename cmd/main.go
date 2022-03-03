package main

import (
	"github.com/Baraulia/COURIER_SERVICE/controller"
	"github.com/Baraulia/COURIER_SERVICE/dao"
	"github.com/Baraulia/COURIER_SERVICE/service"
	_ "github.com/lib/pq"
	"log"
	"net/http"
	"os"
)

// @title Courier Service
// @description Courier Service for Food Delivery Application
func main() {
	log.Println("Start...")
	database, err := dao.NewPostgresDB(dao.PostgresDB{
		Host:     os.Getenv("HOST"),
		Port:     os.Getenv("DB_PORT"),
		User:     os.Getenv("DB_USER"),
		Password: os.Getenv("DB_PASSWORD"),
		DBName:   os.Getenv("DB_DATABASE"),
		SSLMode:  os.Getenv("DB_SSL_MODE")})
	if err != nil {
		log.Fatal("failed to initialize dao:", err.Error())
	}
	repos := dao.NewRepository(database)
	services := service.NewService(repos)
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
