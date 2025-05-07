package handler

import (
	"net/http"

	"github.com/carinhanha/gin/schemas"
	"github.com/gin-gonic/gin"
)

// BasePath /api/v1
// @Summary Create Opening
// @Description Create a new job opening
// @Tags openings
// @Accept json
// @Produce json
// @Param request body CreateOpeningRequest true "Request body"
// @Success 200 {object} CreateOpeningResponse
// @failure 400 {object} ErrorResponse
// @failure 500 {object} ErrorResponse
// @Router /opening [post]
func CreateOpeningHandler(c *gin.Context) {
	request := CreateOpeningRequest{}

	c.BindJSON(&request)

	// logger.Infof("request received: %+v", request)
	if err := request.Validate(); err != nil {
		logger.Errf("validation error: %v", err.Error())
		sendError(c, http.StatusBadRequest, err.Error())
		return
	}

	opening := schemas.Opening{
		Role:     request.Role,
		Company:  request.Company,
		Location: request.Location,
		Remote:   *request.Remote,
		Link:     request.Link,
		Salary:   request.Salary,
	}

	if err := db.Create(&opening).Error; err != nil {
		logger.Errf("error creating opening: %v", err.Error())
		sendError(c, http.StatusInternalServerError, "error create openong on database")
		return
	}

	sendSuccess(c, "create-opening,", opening)
}
