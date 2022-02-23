package main

import (
	"github.com/Baraulia/COURIER_SERVICE/Controllers"
	"github.com/Baraulia/COURIER_SERVICE/db"
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
	database, err := db.NewPostgresDB(db.PostgresDB{
		Host:     os.Getenv("HOST"),
		Port:     os.Getenv("DB_PORT"),
		User:     os.Getenv("DB_USER"),
		Password: os.Getenv("DB_PASSWORD"),
		DBName:   os.Getenv("DB_DATABASE"),
		SSLMode:  os.Getenv("DB_SSL_MODE")})
	if err != nil {
		log.Fatal("failed to initialize db:", err.Error())
	}
	repos := db.NewRepository(database)
	services := service.NewService(repos)
	handlers := Controllers.NewHandler(services)
	host := os.Getenv("API_SERVER_PORT")
	s := &http.Server{
		Addr:    ":" + host,
		Handler: handlers.InitRoutes(),
	}
	err = s.ListenAndServe()
	if err != nil {
		log.Println("failed to initialize port:", err.Error())
	}

}

/*
func main() {
	log.Println("Start...")
	database, err := db.NewPostgresDB(db.PostgresDB{
		"159.223.1.135",
		"5434",
		"courierteam1",
		"qwerty",
		"courier_db",
		"disable"})
	if err != nil {
		log.Fatal("failed to initialize db:", err.Error())
	}
	repos := db.NewRepository(database)
	services := service.NewService(repos)
	handlers := Controllers.NewHandler(services)
	//host:=os.Getenv("API_SERVER_PORT")
	s := &http.Server{
		Addr:    ":8080", // ":"+host,
		Handler: handlers.InitRoutes(),
	}
	err = s.ListenAndServe()
	if err != nil {
		log.Println("failed to initialize port:", err.Error())
	}

}

*/
