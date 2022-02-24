package controller

import (
	"github.com/Baraulia/COURIER_SERVICE/model"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"net/http"
	"strconv"
)

// getCouriers godoc
// @Summary getCouriers
// @Description get all couriers
// @Tags Couriers
// @Accept  json
// @Produce  json
// @Success 200 {object} dao.SmallInfo
// @Failure 400 {string} string
// @Failure 500 {string} string
// @Router /couriers [get]
func (h *Handler) GetCouriers(ctx *gin.Context) {
	Couriers, err := h.services.CourierApp.GetCouriers()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, model.ErrorResponse{Message: err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, Couriers)
}

// getCourier by ID godoc
// @Summary getCourier
// @Description get courier by ID
// @Tags Courier
// @Accept  json
// @Produce  json
// @Param id path int true "Courier ID"
// @Success 200 {object} dao.SmallInfo
// @Failure 400 {string} string
// @Failure 500 {string} err
// @Router /courier/{id} [get]
func (h *Handler) GetCourier(ctx *gin.Context) {
	var Courier model.SmallInfo
	if err := ctx.ShouldBindJSON(&Courier); err != nil {
		logrus.Warnf("Handler createUser (binding JSON):%s", err)
		ctx.JSON(http.StatusBadRequest, model.ErrorResponse{Message: "invalid request"})
		return
	}
	idQuery := ctx.Param("id")
	id, err := strconv.Atoi(idQuery)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err})
		return
	}
	Courier, err = h.services.CourierApp.GetCourier(uint16(id))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"No such courier": err})
		return
	}
	ctx.JSON(http.StatusOK, Courier)
}

// postCourier  godoc
// @Summary postCourier
// @Description post new courier
// @Tags Courier
// @Accept  json
// @Produce  json
// @Param input body dao.Courier true "Courier"
// @Success 200 {object} dao.Courier
// @Failure 400 {string} string
// @Failure 500 {string} err
// @Router /courier [post]
func (h *Handler) SaveCourier(ctx *gin.Context) {
	var input *model.Courier
	if err := ctx.ShouldBindJSON(&input); err != nil {
		logrus.Warnf("Handler createUser (binding JSON):%s", err)
		ctx.JSON(http.StatusBadRequest, model.ErrorResponse{Message: "invalid request"})
		return
	}
	permId, err := h.services.CourierApp.SaveCourier(input)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, model.ErrorResponse{Message: err.Error()})
		return
	}
	ctx.JSON(http.StatusCreated, map[string]uint16{
		"id": permId,
	})
}

