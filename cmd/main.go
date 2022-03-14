package main

import (
	"context"
	"github.com/Baraulia/COURIER_SERVICE/GRPC/grpcClient"
	"github.com/Baraulia/COURIER_SERVICE/GRPC/grpcServer"
	"github.com/Baraulia/COURIER_SERVICE/controller"
	"github.com/Baraulia/COURIER_SERVICE/dao"
	"github.com/Baraulia/COURIER_SERVICE/server"
	"github.com/Baraulia/COURIER_SERVICE/service"
	_ "github.com/lib/pq"
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
		log.Fatal("failed to initialize dao:", err.Error())
	}
	grpcCli := grpcClient.NewGRPCClient(os.Getenv("HOST"))
	repository := dao.NewRepository(database)
	services := service.NewService(repository, grpcCli)
	handlers := controller.NewHandler(services)
	port := os.Getenv("API_SERVER_PORT")

	serv := new(server.Server)

	go func() {
		err := serv.Run(port, handlers.InitRoutesGin())
		if err != nil {
			log.Fatalf("Error occured while running http server: %s", err.Error())
		}
	}()
	go func() {
		grpcServer.NewGRPCServer(services)
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit
	if err := serv.Shutdown(context.Background()); err != nil {
		log.Fatalf("Error occured while shutting down http server: %s", err.Error())
	}

}
