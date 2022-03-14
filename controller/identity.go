package controller

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strings"
)


func (h *Handler) userIdentity(ctx *gin.Context) {
	header := ctx.GetHeader("Authorization")
	if header == "" {
		log.Println("User:empty auth header")
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "empty auth header"})
		return
	}
	headerParts := strings.Split(header, " ")
	if len(headerParts) != 2 || headerParts[0] != "Bearer" {
		log.Println("User:invalid auth header")
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "invalid auth header"})
		return
	}
	if len(headerParts[1]) == 0 {
		log.Println("User:token is empty")
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "token is empty"})
		return
	}
	ok, err := h.services.CourierApp.CheckRights(headerParts[1], "Courier")
	if err != nil || !ok {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized,gin.H{"message": err.Error()})
		return
	}
	ctx.Next()
}

