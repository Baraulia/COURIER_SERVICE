package main

import (
	"github.com/Baraulia/COURIER_SERVICE/controller"
	"github.com/Baraulia/COURIER_SERVICE/dao"
	"github.com/Baraulia/COURIER_SERVICE/server"
	"github.com/Baraulia/COURIER_SERVICE/service"
	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"
	"log"
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
		log.Fatal("failed to initialize db:", err.Error())
	}
	repos := dao.NewRepository(database)
	services := service.NewService(repos)
	handlers := controller.NewHandler(services)
	port := os.Getenv("API_SERVER_PORT")

	serv := new(server.Server)

	go func() {
		err := serv.Run(port, handlers.InitRoutesGin())
		if err != nil {
			logrus.Panicf("Error occured while running http server: %s", err.Error())
		}
	}()


}

