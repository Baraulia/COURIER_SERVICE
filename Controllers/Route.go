package Controllers

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
	r := gin.Default()
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	r.Use(
		middleware.CorsMiddleware,
	)

	couriers := r.Group("/couriers")
	{
		couriers.GET("/", h.GetCouriers)

	}

	courier := r.Group("/courier")
	{
		courier.GET("/:id", h.GetCourier)
	}

	orders := r.Group("/orders")
	{
		orders.GET("/:id", h.GetOrders)

	}

	order := r.Group("/order")
	{
		order.GET("/:id", h.GetOrder)
		order.PUT("/status_change/:id", h.ChangeOrderStatus)
	}

	return r
}
