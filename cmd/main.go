package main

import (
	"github.com/Baraulia/COURIER_SERVICE/Controllers"
	"github.com/Baraulia/COURIER_SERVICE/db"
	"github.com/Baraulia/COURIER_SERVICE/service"
	_ "github.com/lib/pq"
	"log"
	"net/http"
)

// @title Courier Service
// @description Courier Service for Food Delivery Application

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
