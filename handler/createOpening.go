package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateOpenHandler(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"msg": "Post Opening",
	})
}
