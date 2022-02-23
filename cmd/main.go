package main

import (
	"database/sql"
	"github.com/Baraulia/COURIER_SERVICE/Controllers"
	"github.com/Baraulia/COURIER_SERVICE/db"
	"github.com/Baraulia/COURIER_SERVICE/service"
	_ "github.com/lib/pq"
	"net/http"
)

// @title Courier Service
// @description Courier Service for Food Delivery Application

func main() {
	var database *sql.DB
	repos := db.NewRepository(database)
	services := service.NewService(repos)
	handlers := Controllers.NewHandler(services)
	s := &http.Server{
		Addr:    ":80",
		Handler: handlers.InitRoutes(),
	}
	s.ListenAndServe()
}
