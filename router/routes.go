package router

import (
	"github.com/gin-gonic/gin"
	"github.com/ighortorquato/GoProgram.git/handler"
)

func initializeRoutes(router *gin.Engine) {
	v1 := router.Group("/api/v1")
	{
		v1.GET("/opening", handler.ShowOpenHandler)
		v1.POST("/opening", handler.CreateOpenHandler)
		v1.DELETE("/opening", handler.DeleteOpenHandler)
		v1.PUT("/opening", handler.UpdateOpenHandler)
		v1.GET("/openings", handler.ListOpenHandler)

	}
}
