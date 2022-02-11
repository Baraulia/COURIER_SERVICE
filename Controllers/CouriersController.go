package Controllers

import (
	"encoding/json"
	"github.com/Baraulia/COURIER_SERVICE/db"
	"github.com/Baraulia/COURIER_SERVICE/other"
	"net/http"
	"strconv"
)

var Couriers []db.SmallInfo

func (h *Handler) GetCouriers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	Couriers, err := h.services.GetCouriers()
	if err != nil {
		other.RespondWithJSON(w, 400, err.Error())
		return
	}
	json.NewEncoder(w).Encode(Couriers)
}

func (h *Handler) GetOneCourier(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var Courier []db.SmallInfo
	id := r.URL.Query().Get("id")
	l, err := strconv.Atoi(id)
	if err != nil {
		other.RespondWithJSON(w, 400, err.Error())
		return
	}

	Courier, err = h.services.GetOneCourier(l)
	if err != nil {
		other.RespondWithJSON(w, 400, err.Error())
		return
	}
	json.NewEncoder(w).Encode(Courier)
}
