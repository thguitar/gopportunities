package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/thguitar/gopportunities.git/schemas"
)

// @BasePath /api/v1

// @Sumary List opening
// @Description List all job opening
// @Tag Openings
// @Accept json
// @Produce json
// @Success 200 {object} ListOpeningsResponse
// @Failure 400 {object} ErrorResponse
// @Router /openings [get]
func ListOpeningsHandler(ctx *gin.Context) {
	openings := []schemas.Opening{}

	if err := db.Find(&openings).Error; err != nil {
		sendError(ctx, http.StatusInternalServerError, "error listing openings")
		return
	}

	sendSuccess(ctx, "list-openings", openings)
}
