package controller

import (
	"fmt"
	"github.com/Baraulia/COURIER_SERVICE/dao"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
)

type delete struct {
	Status bool `json:"deleted"`
}

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
	Couriers, err := h.services.GetCouriers()
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err})
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
// @Success 200 {object} dao.Courier
// @Failure 400 {string} string
// @Failure 500 {string} err
// @Router /courier/{id} [get]
func (h *Handler) GetCourier(ctx *gin.Context) {
	idQuery := ctx.Param("id")
	id, err := strconv.Atoi(idQuery)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err})
		return
	}
	Courier, err := h.services.GetCourier(id)
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
	var Courier *dao.Courier
	if err := ctx.ShouldBindJSON(&Courier); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "Invalid request"})
		return
	}
	Courier, err := h.services.SaveCourier(Courier)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": err})
		return
	}
	ctx.JSON(http.StatusCreated, Courier)
}

// ChangeCourierStatus by courier ID godoc
// @Summary changeCourierStatus
// @Description put courier status by courier ID
// @Tags Courier
// @Accept  json
// @Produce  json
// @Param id path int true "ID"
// @Param input body bool true "deleted"
// @Success 200 {object} int
// @Failure 400 {string} string
// @Failure 500 {string} err
// @Router /courier/{id} [put]
func (h *Handler) UpdateCourier(ctx *gin.Context) {
	idQuery := ctx.Param("id")
	id, err := strconv.Atoi(idQuery)

	var txt delete
	var status bool

	if err := ctx.ShouldBindJSON(&txt); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "Invalid request"})
		return
	}

	status = txt.Status
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"Error with query parameter": err})
		return
	}
	courierId, err := h.services.CourierApp.UpdateCourier(uint16(id), status)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"No such courier": err})
		return
	}
	ctx.JSON(http.StatusOK, map[string]interface{}{
		"Courier id": courierId})
}

// @Summary SaveCourierPhoto
// @Description set photo to DO Spaces and it's way to DB
// @Tags Couriers
// @Accept  image/jpeg
// @Produce   json
// @Param id query int true "id courier"
// @Param logo  formData  file  true  "logo image"
// @Success 204
// @Failure 400 {string} string
// @Router /couriers/photo [post]
func (h *Handler) SaveCourierPhoto(ctx *gin.Context) {
	id, er := strconv.Atoi(ctx.Query("id"))
	if er != nil || id <= 0 {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "expect an integer greater than 0"})
		return
	}
	if ctx.Request.Body == nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "empty"})
	}
	defer ctx.Request.Body.Close()
	cover, err := ioutil.ReadAll(ctx.Request.Body)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "Invalid request"})
	}
	if err := h.services.SaveCourierPhoto(cover, id); err != nil {
		log.Println(err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": fmt.Sprintf("Error: %s", err)})
		return
	}
	ctx.Status(http.StatusNoContent)
}

type listCouriers struct {
	Data []dao.Courier `json:"data"`
}

// @Summary GetCouriersOfCourierService
// @Description get list of all couriers by courier service id
// @Tags Couriers
// @Produce json
// @Param page query int true "page"
// @Param limit query int true "limit"
// @Param iddeliveryservice query int true "iddeliveryservice"
// @Success 200 {object} listCouriers
// @Failure 400 {string} string
// @Failure 500 {string} string
// @Router /couriers/service [get]
func (h *Handler) GetCouriersOfCourierService(ctx *gin.Context) {
	page, er := strconv.Atoi(ctx.Query("page"))
	if er != nil || page == 0 {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "page query param is wrong. Expected an integer greater than 0"})
		return
	}
	limit, er1 := strconv.Atoi(ctx.Query("limit"))
	if er1 != nil || limit == 0 {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "limit query param is wrong. Expected an integer greater than 0"})
		return
	}
	idService, er := strconv.Atoi(ctx.Query("iddeliveryservice"))
	if er != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "expect an integer greater than 0"})
		return
	}

	Couriers, err := h.services.GetCouriersOfCourierService(limit, page, idService)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": fmt.Sprintf("Error: %s", err)})
		return
	}
	ctx.JSON(http.StatusOK, listCouriers{Data: Couriers})

}
