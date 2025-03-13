package routes

import (
	"example/hello/controllers"
	"example/hello/db"
	"example/hello/repositories"
	"example/hello/services"
	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine) {
	db:=connectiondb.ConnectDB()
	stockRepo:= repositories.NewStockRepository(db)
	stockServ:= services.NewStockService(stockRepo)
	stockController:= controllers.NewStockController(stockServ)

	api := router.Group("api/v1")
	{
		api.GET("stock", stockController.GetAll)
		api.GET("stock/:id", stockController.GetOne)
		api.POST("stock", stockController.Create)
		api.PUT("stock/:id", stockController.Update)
		api.DELETE("stock/:id", stockController.Delete)
	}
}

