package handler

import (
	"net/http"

	"github.com/carinhanha/gin/schemas"
	"github.com/gin-gonic/gin"
)

// BasePath /api/v1
// @Summary List Opening
// @Description List a new job opening
// @Tags openings
// @Accept json
// @Produce json
// @Success 200 {object} ListOpeningResponse
// @failure 500 {object} ErrorResponse
// @Router /openings [get]
func ListOpeningHandler(c *gin.Context) {
	openings := []schemas.Opening{}

	if err := db.Find(&openings).Error; err != nil {
		sendError(c, http.StatusInternalServerError, "error listing openings")
		return
	}

	sendSuccess(c, "list-openings", openings)
}
