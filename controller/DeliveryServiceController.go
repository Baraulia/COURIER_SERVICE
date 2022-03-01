package controller

import (
	"errors"
	"fmt"
	"github.com/Baraulia/COURIER_SERVICE/dao"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
)

// @Summary  CreateDeliveryService
// @Tags DeliveryService
// @Description create a Delivery Service
// @ID CreateDeliveryService
// @Accept  json
// @Produce json
// @Param input body dao.DeliveryService true "Delivery Service"
// @Success 200 {object} dao.DeliveryService
// @Failure 400 {object} map[string]interface{}
// @Failure 500 {string} string
// @Router /deliveryservice [post]
func (h *Handler) CreateDeliveryService(ctx *gin.Context) {
	var service dao.DeliveryService
	if err := ctx.ShouldBindJSON(&service); err != nil {
		log.Println(err)
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "Invalid request"})
		return
	}
	if service.Email == "" || service.Name == "" {
		log.Println(errors.New("empty fields"))
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "empty fields"})
		return
	}
	idService, err := h.services.CreateDeliveryService(service)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": fmt.Sprintf("Error: %s", err)})
		return
	}
	ctx.JSON(http.StatusOK, map[string]interface{}{"id": idService})
}

// @Summary GetDeliveryServiceById
// @Description get delivery service by id
// @Tags DeliveryService
// @Produce json
// @Param id path int true "id"
// @Success 200 {object} dao.DeliveryService
// @Failure 400 {string} string
// @Failure 500 {string} string
// @Router /deliveryservice/{id} [get]
func (h *Handler) GetDeliveryServiceById(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "expect an integer greater than 0"})
		return
	}

	service, err := h.services.DeliveryServiceApp.GetDeliveryServiceById(id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": fmt.Sprintf("Error: %s", err)})
		return
	}
	ctx.JSON(http.StatusOK, service)
}
