package main

import (
	"github.com/Baraulia/COURIER_SERVICE/Controllers"
	"github.com/Baraulia/COURIER_SERVICE/db"
	_ "github.com/lib/pq"
	"log"
	"net/http"
)

func main() {
	db.ConnectDB()
	http.HandleFunc("/couriers", Controllers.GetCouriers)
	err := http.ListenAndServe(":80", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
