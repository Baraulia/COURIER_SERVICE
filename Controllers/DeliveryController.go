package Controllers

import (
	"encoding/json"
	"github.com/Baraulia/COURIER_SERVICE/db"
	"github.com/Baraulia/COURIER_SERVICE/other"
	"net/http"
	"strconv"
)

var Orders []db.Order

func (h *Handler) GetOrders(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	Orders, err := h.services.GetOrders()
	if err != nil {
		other.RespondWithJSON(w, 400, err.Error())
		return
	}
	json.NewEncoder(w).Encode(Orders)
}

func (h *Handler) GetOneOrder(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var Order []db.Order
	id := r.URL.Query().Get("id")
	l, err := strconv.Atoi(id)
	if err != nil {
		other.RespondWithJSON(w, 400, err.Error())
		return
	}
	Order, err = h.services.GetOneOrder(l)
	if err != nil {
		other.RespondWithJSON(w, 400, err.Error())
		return
	}
	json.NewEncoder(w).Encode(Order)
}

func (h *Handler) ChangeOrderStatus(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	id := r.URL.Query().Get("id")
	l, err := strconv.Atoi(id)
	if err != nil {
		other.RespondWithJSON(w, 400, err.Error())
		return
	}
	err = h.services.ChangeOrderStatus(l)
	if err != nil {
		other.RespondWithJSON(w, 400, err.Error())
		return
	}
	json.NewEncoder(w).Encode(map[string]interface{}{
		"Order id": id})
}
