package handler

import (
	"net/http"

	"github.com/carinhanha/gin/schemas"
	"github.com/gin-gonic/gin"
)

// BasePath /api/v1
// @Summary Update Opening
// @Description Update a job opening
// @Tags openings
// @Accept json
// @Produce json
// @Param id query string true "opening idetification"
// @Param opening body UpdateOpeningRequest true "Opening data to update"
// @Success 200 {object} UpdateOpeningResponse
// @failure 400 {object} ErrorResponse
// @failure 404 {object} ErrorResponse
// @Router /opening [put]
func UpdateOpeningHandler(c *gin.Context) {
	request := UpdateOpeningRequest{}

	c.BindJSON(&request)

	if err := request.Validate(); err != nil {
		logger.Errf("validation error %v", err.Error())
		sendError(c, http.StatusBadRequest, err.Error())
		return
	}

	id := c.Query("id")
	if id == "" {
		sendError(c, http.StatusBadRequest, errParamRequired("id", "queryParameter").Error())
		return
	}

	opening := schemas.Opening{}

	if err := db.First(&opening, id).Error; err != nil {
		sendError(c, http.StatusNotFound, "opening not found")
	}

	// update opening
	if request.Role != "" {
		opening.Role = request.Role
	}
	if request.Company != "" {
		opening.Company = request.Company
	}
	if request.Location != "" {
		opening.Location = request.Location
	}
	if request.Remote != nil {
		opening.Remote = *request.Remote
	}
	if request.Link != "" {
		opening.Link = request.Link
	}
	if request.Salary >= 0 {
		opening.Salary = request.Salary
	}

	if err := db.Save(&opening).Error; err != nil {
		logger.Errf("error updating opening %v", err.Error())
		sendError(c, http.StatusInternalServerError, "error updating opening")
		return
	}

	sendSuccess(c, "update-opening", opening)
}
