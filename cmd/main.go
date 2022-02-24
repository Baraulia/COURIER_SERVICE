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
		Host:     os.Getenv("HOST"),
		Port:     os.Getenv("DB_PORT"),
		User:     os.Getenv("DB_USER"),
		Password: os.Getenv("DB_PASSWORD"),
		DBName:   os.Getenv("DB_DATABASE"),
		SSLMode:  os.Getenv("DB_SSL_MODE")})
	if err != nil {
		log.Println("failed to initialize db:", err.Error())
	}
	repos := dao.NewRepository(db)
	services := model.NewService(repos)
	handlers := controller.NewHandler(services)
	port := os.Getenv("API_SERVER_PORT")
	s := &http.Server{
		Addr:    ":" + port,
		Handler: handlers.InitRoutesGin(),
	}
	err = s.ListenAndServe()
	if err != nil {
		log.Println("failed to initialize db:", err.Error())
	}

}
