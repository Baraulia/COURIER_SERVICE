package Controllers

import (
	"github.com/Baraulia/COURIER_SERVICE/service"
	"github.com/gorilla/mux"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}

func (h *Handler) InitRoutes() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/orders", h.GetOrders).Methods("GET")
	r.HandleFunc("/order", h.GetOneOrder).Methods("GET")
	return r
}
