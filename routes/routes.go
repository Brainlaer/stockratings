package routes

import (
	"example/hello/controllers"
	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine) {
	api := router.Group("api/v1")
	{
		api.GET("stock", controllers.GetAll)
		api.POST("stock", controllers.Post)
		api.PUT("stock", controllers.Put)
		api.DELETE("stock", controllers.Delete)
	}

}

