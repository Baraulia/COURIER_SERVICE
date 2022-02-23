package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/swaggo/files"
	"github.com/swaggo/gin-swagger"
	_ "stlab.itechart-group.com/go/food_delivery/courier_service/docs"
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
	gin.SetMode(gin.ReleaseMode)
	router := gin.New()

	router.Use(
		middleware.CorsMiddleware,
	)
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	order := router.Group("/deliveryservice")
	{
		order.POST("/", h.CreateDeliveryService)
	}
	return router
}
