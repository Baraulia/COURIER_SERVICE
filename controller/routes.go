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

func (h *Handler) InitRoutesGin() *gin.Engine {
	gin.SetMode(gin.ReleaseMode)
	router := gin.Default()
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	router.Use(
		middleware.CorsMiddleware,
	)

	couriers := router.Group("/couriers")
	//couriers.Use(h.userIdentity)
	{
		couriers.GET("/", h.GetCouriers)
		couriers.POST("/photo", h.SaveCourierPhoto)
		couriers.GET("/service", h.GetCouriersOfCourierService)
	}

	courier := router.Group("/courier")
	//courier.Use(h.userIdentity)
	{
		courier.GET("/:id", h.GetCourier)
		courier.POST("/", h.SaveCourier)
		courier.PUT("/:id", h.UpdateCourier)
	}

	orders := router.Group("/orders")
	//orders.Use(h.userIdentity)
	{
		orders.GET("/completed", h.GetCourierCompletedOrders)
		orders.GET("/", h.GetAllOrdersOfCourierService)
		orders.GET("/bymonth", h.GetCourierCompletedOrdersByMonth)
		orders.GET("/:id", h.GetOrders)
		orders.PUT("/:id", h.UpdateOrder)
		orders.GET("/service/completed", h.GetCompletedOrdersOfCourierService)
		orders.GET("/manager", h.GetOrdersOfCourierServiceForManager)

	}

	order := router.Group("/order")
	//order.Use(h.userIdentity)
	{
		order.GET("/:id", h.GetOrder)
		order.PUT("/status_change/:id", h.ChangeOrderStatus)
		order.GET("/detailed/:id", h.GetDetailedOrderById)
	}

	deliveryService := router.Group("/deliveryservice")
	//deliveryService.Use(h.userIdentity)
	{
		deliveryService.POST("/", h.CreateDeliveryService)
		deliveryService.GET("/:id", h.GetDeliveryServiceById)
		deliveryService.GET("/", h.GetAllDeliveryServices)
		deliveryService.PUT("/:id", h.UpdateDeliveryService)
		deliveryService.POST("/logo", h.SaveLogoController)
	}
	return router
}
