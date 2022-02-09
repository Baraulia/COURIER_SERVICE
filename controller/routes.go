package controller

import (
	"github.com/Baraulia/COURIER_SERVICE/model"
	"github.com/gorilla/mux"
)

type Handler struct {
	services *model.Service
}

func NewHandler(services *model.Service) *Handler {
	return &Handler{services: services}
}

func (h *Handler) InitRoutes() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/orders/completed", h.GetCourierCompletedOrders).Methods("GET")
	r.HandleFunc("/orders", h.GetAllOrdersOfCourierService).Methods("GET")
	r.HandleFunc("/orders/bymonth",h.GetCourierCompletedOrdersByMonth).Methods("GET")
	return r
}

