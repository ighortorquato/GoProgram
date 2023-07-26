package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ighortorquato/GoProgram.git/schemas"
)

// @BasePath /api/v1

// @Summary List opening
// @Description List a new job opening
// @Tags Openings
// @Accept json
// @Produce json
// @Success 200 {object} ListOpeningsResponse
// @Failure 500 {objetc} ErrorResponse
// @Router /openings [get]
func ListOpenHandler(ctx *gin.Context) {
	openings := []schemas.Opening{}

	if err := db.Find(&openings).Error; err != nil {
		sendError(ctx, http.StatusInternalServerError, "error listing openings")
		return
	}
	sendSuccess(ctx, "list-openings", openings)
}
