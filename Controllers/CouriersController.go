package Controllers

import (
	"encoding/json"
	"fmt"
	"github.com/Baraulia/COURIER_SERVICE/Models"
	"github.com/Baraulia/COURIER_SERVICE/db"
	"github.com/Baraulia/COURIER_SERVICE/other"
	"net/http"
	"strconv"
)

var Couriers []db.SmallInfo

func GreetingPage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello! This is a courier delivery page")
}

func GetCouriers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	Couriers = Models.GetCouriers()
	json.NewEncoder(w).Encode(Couriers)
}

func GetOneCourier(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var Courier db.SmallInfo
	id := r.URL.Query().Get("id")
	l, err := strconv.Atoi(id)
	if err != nil {
		other.RespondWithError(w, 400, "bad request")
		return
	}

	Courier = Models.GetOneCourier(l)
	json.NewEncoder(w).Encode(Courier)
}
