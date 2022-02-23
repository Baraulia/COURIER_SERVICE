package Controllers

import (
	"github.com/Baraulia/COURIER_SERVICE/db"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func (h *Handler) GetCouriers(ctx *gin.Context) {
	Couriers, err := h.services.GetCouriers()
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err})
		return
	}
	ctx.JSON(http.StatusOK, Couriers)
}

func (h *Handler) GetCourier(ctx *gin.Context) {
	var Courier db.SmallInfo
	idQuery := ctx.Query("id")
	id, err := strconv.Atoi(idQuery)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err})
		return
	}
	Courier, err = h.services.GetCourier(id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"No such courier": err})
		return
	}
	ctx.JSON(http.StatusOK, Courier)
}

func (h *Handler) SaveCourier(ctx *gin.Context) {
	var Courier *db.Courier
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
