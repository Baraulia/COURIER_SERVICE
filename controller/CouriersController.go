package controller

import (
	"github.com/Baraulia/COURIER_SERVICE/model"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"net/http"
	"strconv"
)

// GetCouriers godoc
// @Summary getCouriers
// @Description get all couriers
// @Tags Courier
// @Accept  json
// @Produce  json
// @Success 200 {object} model.SmallInfo
// @Failure 400 {string} model.ErrorResponse
// @Failure 500 {string} model.ErrorResponse
// @Router /couriers [get]
func (h *Handler) GetCouriers(ctx *gin.Context) {
	Couriers, err := h.services.CourierApp.GetCouriers()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, model.ErrorResponse{Message: err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, Couriers)
}

// GetCourier by ID godoc
// @Summary getCourier
// @Description get courier by ID
// @Tags Courier
// @Accept  json
// @Produce  json
// @Param id path int true "Courier ID"
// @Success 200 {object} model.SmallInfo
// @Failure 400 {string} model.ErrorResponse
// @Failure 500 {string} model.ErrorResponse
// @Router /courier/{id} [get]
func (h *Handler) GetCourier(ctx *gin.Context) {
	var Courier *model.SmallInfo
	if err := ctx.ShouldBindJSON(&Courier); err != nil {
		logrus.Warnf("Handler getCourier (binding JSON):%s", err)
		ctx.JSON(http.StatusBadRequest, model.ErrorResponse{Message: "invalid request"})
		return
	}
	idQuery := ctx.Param("id")
	id, err := strconv.Atoi(idQuery)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, model.ErrorResponse{Message: "invalid request"})
		return
	}
	Courier, err = h.services.CourierApp.GetCourier(uint16(id))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError,model.ErrorResponse{Message: err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, Courier)
}

// SaveCourier  godoc
// @Summary postCourier
// @Description post new courier
// @Tags Courier
// @Accept  json
// @Produce  json
// @Param input body model.Courier true "Courier"
// @Success 200 {object} model.Courier
// @Failure 400 {string} model.ErrorResponse
// @Failure 500 {string} model.ErrorResponse
// @Router /courier [post]
func (h *Handler) SaveCourier(ctx *gin.Context) {
	var input *model.Courier
	if err := ctx.ShouldBindJSON(&input); err != nil {
		logrus.Warnf("Handler createCourier (binding JSON):%s", err)
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

// DeleteCourier godoc
// @Summary deleteCourierByID
// @Description delete courier by ID
// @Tags Courier
// @Accept  json
// @Produce  json
// @Param id path int true "Courier ID"
// @Success 200  {string} map[string]interface{}
// @Failure 400 {object} model.ErrorResponse
// @Failure 500 {object} model.ErrorResponse
// @Router /courier/{id} [delete]
func (h *Handler) DeleteCourier(ctx *gin.Context) {
	paramID := ctx.Param("id")
	varID, err := strconv.Atoi(paramID)
	if err != nil {
		logrus.Warnf("Handler deleteCourierByID (reading param):%s", err)
		ctx.JSON(http.StatusBadRequest, model.ErrorResponse{Message: "Invalid id"})
		return
	}
	id, err := h.services.CourierApp.DeleteCourier(uint16(varID))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, model.ErrorResponse{Message: err.Error()})
		return
	} else {
		ctx.JSON(http.StatusOK, map[string]interface{}{
			"Courier profile has successfully deleted": id,
		})
	}
}
