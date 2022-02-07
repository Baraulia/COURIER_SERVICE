package Controllers

import (
	"encoding/json"
	"github.com/Baraulia/COURIER_SERVICE/Models"
	"github.com/Baraulia/COURIER_SERVICE/db"
	"net/http"
)

var Couriers []db.Courier

func GetCouriers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	Couriers = Models.GetCouriers(Couriers)
	json.NewEncoder(w).Encode(Couriers)
}
