package handler

import (
	"net/http"

	"github.com/carinhanha/gin/schemas"
	"github.com/gin-gonic/gin"
)

// BasePath /api/v1
// @Summary Show Opening
// @Description Show a job opening
// @Tags openings
// @Accept json
// @Produce json
// @Param id query string true "opening idetification"
// @Success 200 {object} ShowOpeningResponse
// @failure 400 {object} ErrorResponse
// @failure 404 {object} ErrorResponse
// @Router /opening [get]
func ShowOpeningHandler(c *gin.Context) {
	id := c.Query("id")
	if id == "" {
		sendError(c, http.StatusBadRequest, errParamRequired("id", "queryParameter").Error())
		return
	}

	opening := schemas.Opening{}

	if err := db.First(&opening, id).Error; err != nil {
		sendError(c, http.StatusNotFound, "opening not found")
	}

	sendSuccess(c, "show-opening", opening)
}
