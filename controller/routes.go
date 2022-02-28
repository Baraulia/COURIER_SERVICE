package controller

import (
	_ "github.com/Baraulia/COURIER_SERVICE/docs"
	"github.com/Baraulia/COURIER_SERVICE/middleware"
	"github.com/Baraulia/COURIER_SERVICE/service"
	"github.com/gin-gonic/gin"
	"github.com/swaggo/files"
	"github.com/swaggo/gin-swagger"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}

func (h *Handler) InitRoutes() *gin.Engine {
	gin.SetMode(gin.ReleaseMode)
	router := gin.Default()
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	router.Use(
		middleware.CorsMiddleware,
	)

	couriers := router.Group("/couriers")
	{
		couriers.GET("/", h.GetCouriers)

	}

	courier := router.Group("/courier")
	{
		courier.GET("/:id", h.GetCourier)
		courier.POST("/", h.SaveCourier)
		//courier.DELETE("/:id", h.DeleteCourier)
	}

	orders := router.Group("/orders")
	{
		orders.GET("/completed", h.GetCourierCompletedOrders)
		orders.GET("/", h.GetAllOrdersOfCourierService)
		orders.GET("/byMonth", h.GetCourierCompletedOrdersByMonth)
		orders.GET("/:id", h.GetOrders)

	}

	order := router.Group("/order")
	{
		order.GET("/:id", h.GetOrder)
		order.PUT("/statusChange/:id", h.ChangeOrderStatus)
	}

	return router
}