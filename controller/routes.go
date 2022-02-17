package controller

import (
	"github.com/Baraulia/COURIER_SERVICE/middleware"
	"github.com/Baraulia/COURIER_SERVICE/model"
	"github.com/gin-gonic/gin"
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
	//r.HandleFunc("/orders/completed", h.GetCourierCompletedOrders).Methods("GET")
	//r.HandleFunc("/orders", h.GetAllOrdersOfCourierService).Methods("GET")
	//r.HandleFunc("/orders/bymonth",h.GetCourierCompletedOrdersByMonth).Methods("GET")
	//r.HandleFunc("/orders",h.AssigningOrderToCourier).Methods("PUT")
	return r
}

func (h *Handler) InitRoutesGin() *gin.Engine {
	router := gin.Default()

	router.Use(
		middleware.CorsMiddleware,
	)

	order := router.Group("/orders")
	{
		order.GET("/completed", h.GetCourierCompletedOrders)
		order.GET("/", h.GetAllOrdersOfCourierService)
		order.GET("/bymonth", h.GetCourierCompletedOrdersByMonth)
		order.PUT("/", h.UpdateOrder)
	}

	return router
}
