package main

import (
	"context"
	"github.com/Baraulia/COURIER_SERVICE/GRPCserver/grpcServer"
	"github.com/Baraulia/COURIER_SERVICE/controller"
	"github.com/Baraulia/COURIER_SERVICE/dao"
	"github.com/Baraulia/COURIER_SERVICE/server"
	"github.com/Baraulia/COURIER_SERVICE/service"
	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"
	"log"
	"os"
	"os/signal"
	"syscall"
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
	repository := dao.NewRepository(database)
	services := service.NewService(repository)
	handlers := controller.NewHandler(services)
	port := os.Getenv("API_SERVER_PORT")

	serv := new(server.Server)

	go func() {
		err := serv.Run(port, handlers.InitRoutes())
		if err != nil {
			logrus.Panicf("Error occured while running http server: %s", err.Error())
		}
	}()
	go func() {
		grpcServer.NewGRPCServer(services)
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit
	if err := serv.Shutdown(context.Background()); err != nil {
		logrus.Panicf("Error occured while shutting down http server: %s", err.Error())
	}

}

