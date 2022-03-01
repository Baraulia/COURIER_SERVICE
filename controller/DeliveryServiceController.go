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

type listDeliveryServices struct {
	Data []dao.DeliveryService `json:"data"`
}

// @Summary GetAllDeliveryServices
// @Description get list of all delivery service
// @Tags DeliveryService
// @Produce json
// @Success 200 {object} listDeliveryServices
// @Failure 400 {string} string
// @Failure 500 {string} string
// @Router /deliveryservice [get]
func (h *Handler) GetAllDeliveryServices(ctx *gin.Context) {
	services, err := h.services.DeliveryServiceApp.GetAllDeliveryServices()
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": fmt.Sprintf("Error: %s", err)})
		return
	}
	ctx.JSON(http.StatusOK, listDeliveryServices{Data: services})
}

// @Summary UpdateDeliveryService
// @Tags DeliveryService
// @Description update delivery service information
// @ID UpdateDeliveryService
// @Accept  json
// @Produce json
// @Param id path int true "order_id"
// @Param input body dao.DeliveryService true "delivery service"
// @Success 204
// @Failure 400 {string} string
// @Router /deliveryservice/{id} [put]
func (h *Handler) UpdateDeliveryService(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		log.Println(err)
		ctx.JSON(http.StatusBadRequest, gin.H{"message": fmt.Sprintf("Error: %s", err)})
		return
	}
	var service dao.DeliveryService
	if err := ctx.ShouldBindJSON(&service); err != nil {
		log.Println(err)
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "Invalid request"})
		return
	}
	service.Id = id
	if err := h.services.UpdateDeliveryService(service); err != nil {
		log.Println(err)
		ctx.JSON(http.StatusBadRequest, gin.H{"message": fmt.Sprintf("Error: %s", err)})
		return
	}
	ctx.Status(http.StatusNoContent)
}
