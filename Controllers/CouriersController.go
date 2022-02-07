package Controllers

import (
	"encoding/json"
	"fmt"
	"github.com/Baraulia/COURIER_SERVICE/Models"
	"github.com/Baraulia/COURIER_SERVICE/db"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

var Couriers []db.SmallInfo

func GreetingPage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello! This is a courier delivery page")
}

func GetCouriers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	Couriers = Models.GetCouriers(Couriers)
	json.NewEncoder(w).Encode(Couriers)
}

func GetOneCourier(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	Couriers = Models.GetCouriers(Couriers)
	for _, courier := range Couriers {
		if strconv.Itoa(int(courier.IdCourier)) == params["id_courier"] {
			json.NewEncoder(w).Encode(courier)
			return
		}
	}
}
