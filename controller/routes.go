package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/swaggo/files"
	"github.com/swaggo/gin-swagger"
	_ "stlab.itechart-group.com/go/food_delivery/courier_service/docs"
	"stlab.itechart-group.com/go/food_delivery/courier_service/middleware"
	"stlab.itechart-group.com/go/food_delivery/courier_service/service"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}

func (h *Handler) InitRoutesGin() *gin.Engine {
	gin.SetMode(gin.ReleaseMode)
	router := gin.Default()

	router.Use(
		middleware.CorsMiddleware,
	)
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	order := router.Group("/orders")
	{
		order.GET("/completed", h.GetCourierCompletedOrders)
		order.GET("/", h.GetAllOrdersOfCourierService)
		order.GET("/bymonth", h.GetCourierCompletedOrdersByMonth)
	}

	return router
}
