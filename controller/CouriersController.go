package controller

import (
	"github.com/Baraulia/COURIER_SERVICE/dao"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
)

// getCouriers godoc
// @Summary getCouriers
// @Security ApiKeyAuth
// @Description get all couriers
// @Tags Couriers
// @Accept  json
// @Produce  json
// @Success 200 {object} dao.SmallInfo
// @Failure 400 {string} string
// @Failure 500 {string} string
// @Router /couriers [get]
func (h *Handler) GetCouriers(ctx *gin.Context) {
	necessaryRole := "Courier manager"
	if err := h.services.AllProjectApp.CheckRoleRights(nil, necessaryRole, ctx.GetString("perms"), ctx.GetString("role")); err != nil {
		log.Print("Handler GetCouriers:not enough rights")
		ctx.JSON(http.StatusUnauthorized, gin.H{"message": "not enough rights"})
		return
	}
	Couriers, err := h.services.AllProjectApp.GetCouriers()
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err})
		return
	}
	ctx.JSON(http.StatusOK, Couriers)
}

// getCourier by ID godoc
// @Summary getCourier
// @Security ApiKeyAuth
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
	necessaryRole := "Courier manager"
	if err := h.services.AllProjectApp.CheckRoleRights(nil, necessaryRole, ctx.GetString("perms"), ctx.GetString("role")); err != nil {
		log.Print("Handler GetCourier:not enough rights")
		ctx.JSON(http.StatusUnauthorized, gin.H{"message": "not enough rights"})
		return
	}
	var Courier dao.SmallInfo
	idQuery := ctx.Param("id")
	id, err := strconv.Atoi(idQuery)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err})
		return
	}
	Courier, err = h.services.AllProjectApp.GetCourier(id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"No such courier": err})
		return
	}
	ctx.JSON(http.StatusOK, Courier)
}

// postCourier  godoc
// @Summary postCourier
// @Security ApiKeyAuth
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
	necessaryRole := "Courier manager"
	if err := h.services.AllProjectApp.CheckRoleRights(nil, necessaryRole, ctx.GetString("perms"), ctx.GetString("role")); err != nil {
		log.Print("Handler SaveCourier:not enough rights")
		ctx.JSON(http.StatusUnauthorized, gin.H{"message": "not enough rights"})
		return
	}
	var Courier *dao.Courier
	if err := ctx.ShouldBindJSON(&Courier); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "Invalid request"})
		return
	}
	Courier, err := h.services.AllProjectApp.SaveCourier(Courier)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": err})
		return
	}
	ctx.JSON(http.StatusCreated, Courier)
}

// ChangeCourierStatus by courier ID godoc
// @Summary changeCourierStatus
// @Security ApiKeyAuth
// @Description put courier status by courier ID
// @Tags Courier
// @Accept  json
// @Produce  json
// @Param id path int true "ID"
// @Success 200 {object} int
// @Failure 400 {string} string
// @Failure 500 {string} err
// @Router /courier/{id} [put]
func (h *Handler) UpdateCourier(ctx *gin.Context) {
	necessaryRole := "Courier manager"
	if err := h.services.AllProjectApp.CheckRoleRights(nil, necessaryRole, ctx.GetString("perms"), ctx.GetString("role")); err != nil {
		log.Print("Handler UpdateCourier:not enough rights")
		ctx.JSON(http.StatusUnauthorized, gin.H{"message": "not enough rights"})
		return
	}
	idQuery := ctx.Param("id")
	id, err := strconv.Atoi(idQuery)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"Error with query parameter": err})
		return
	}
	courierId, err := h.services.AllProjectApp.UpdateCourier(uint16(id))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"No such courier": err})
		return
	}
	ctx.JSON(http.StatusOK, map[string]interface{}{
		"Courier id": courierId})
}
