package Controllers

import (
	"github.com/Baraulia/COURIER_SERVICE/service"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}

func (h *Handler) InitRoutes() *gin.Engine {
	r := gin.Default()

	orders := r.Group("/orders")
	{
		orders.GET("/", h.GetOrders)

	}

	order := r.Group("/order")
	{
		order.GET("/", h.GetOrder)
	}
	return r
}
