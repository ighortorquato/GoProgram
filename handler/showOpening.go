package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func ShowOpenHandler(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"msg": "GET Opening",
	})
}