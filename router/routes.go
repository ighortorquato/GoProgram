package router

import (
	"github.com/gin-gonic/gin"
	"github.com/ighortorquato/GoProgram.git/docs"
	"github.com/ighortorquato/GoProgram.git/handler"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func initializeRoutes(router *gin.Engine) {
	//Initialize Handler
	handler.InitializeHandler()
	basePath := "/api/v1"
	docs.SwaggerInfo.BasePath = basePath

	v1 := router.Group(basePath)
	{
		v1.GET("/opening", handler.ShowOpenHandler)
		v1.POST("/opening", handler.CreateOpeningHandler)
		v1.DELETE("/opening", handler.DeleteOpenHandler)
		v1.PUT("/opening", handler.UpdateOpenHandler)
		v1.GET("/openings", handler.ListOpenHandler)

	}
	//Initialize swagger
	router.GET("swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
}
