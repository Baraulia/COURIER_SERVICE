package Controllers

import (
	"encoding/json"
	"fmt"
	"github.com/Baraulia/COURIER_SERVICE/Models"
	"github.com/Baraulia/COURIER_SERVICE/db"
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
	var Courier db.SmallInfo
	id := r.URL.Query().Get("id")
	l, _ := strconv.Atoi(id)
	Courier = Models.GetOneCourier(l)
	json.NewEncoder(w).Encode(Courier)
}
