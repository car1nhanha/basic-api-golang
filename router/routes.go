package router

import (
	"github.com/carinhanha/gin/handler"
	"github.com/gin-gonic/gin"

	docs "github.com/carinhanha/gin/docs"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	// "github.com/swaggo/swag/example/basic/docs"
)

func InitializeRoutes(router *gin.Engine) {
	// initialize handler
	handler.InitializeHandler()

	basePath := "/api/v1"
	docs.SwaggerInfo.BasePath = basePath

	v1 := router.Group(basePath)
	{
		v1.GET("/opening", handler.ShowOpeningHandler)

		v1.POST("/opening", handler.CreateOpeningHandler)

		v1.DELETE("/opening", handler.DeleteOpeningHandler)

		v1.PUT("/opening", handler.UpdateOpeningHandler)

		v1.GET("/openings", handler.ListOpeningHandler)
	}
	// initialize swagger
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
}
