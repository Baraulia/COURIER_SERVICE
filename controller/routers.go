package controller

import (
"github.com/gin-gonic/gin"
"stlab.itechart-group.com/go/food_delivery/courier_service/middleware"
"stlab.itechart-group.com/go/food_delivery/courier_service/model"
)

type Handler struct {
	services *model.Service
}

func NewHandler(services *model.Service) *Handler {
	return &Handler{services: services}
}

func (h *Handler) InitRoutesGin() *gin.Engine {
	router := gin.Default()

	router.Use(
		middleware.CorsMiddleware,
	)

	order := router.Group("/orders")
	{
		order.PUT("/", h.UpdateOrder)
		order.GET("/completed", h.GetAllCompletedOrdersByService)
		order.GET("/",h.GetDetailedOrdersById)
	}

	return router
}

