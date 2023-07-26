package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ighortorquato/GoProgram.git/schemas"
)

func ListOpenHandler(ctx *gin.Context) {
	openings := []schemas.Opening{}

	if err := db.Find(&openings).Error; err != nil {
		sendError(ctx, http.StatusInternalServerError, "error listing openings")
		return
	}
	sendSuccess(ctx, "list-openings", openings)
}
