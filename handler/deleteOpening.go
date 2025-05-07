package handler

import (
	"fmt"
	"net/http"

	"github.com/carinhanha/gin/schemas"
	"github.com/gin-gonic/gin"
)

// BasePath /api/v1
// @Summary Delete Opening
// @Description Delete a new job opening
// @Tags openings
// @Accept json
// @Produce json
// @Param id query string true "opening idetification"
// @Success 200 {object} DeleteOpeningResponse
// @failure 400 {object} ErrorResponse
// @failure 404 {object} ErrorResponse
// @Router /opening [delete]
func DeleteOpeningHandler(c *gin.Context) {
	id := c.Query("id")
	if id == "" {
		sendError(c, http.StatusBadRequest, errParamRequired("id", "queryParameter").Error())
		return
	}

	opening := schemas.Opening{}

	if err := db.First(&opening, id).Error; err != nil {
		sendError(c, http.StatusNotFound, fmt.Sprintf("opening with id: %s not found", id))
		return
	}

	if err := db.Delete(&opening).Error; err != nil {
		sendError(c, http.StatusInternalServerError, fmt.Sprintf("error deleting opening with id: %s", id))
		return
	}

	sendSuccess(c, "delete-opening", opening)
}
