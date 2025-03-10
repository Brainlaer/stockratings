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
	stockRatingRepo:= repositories.NewStockRatingRepository(db)
	stockRatingServ:= services.NewStockService(stockRatingRepo)
	stockRatingController:= controllers.NewStockRatingController(stockRatingServ)

	api := router.Group("api/v1")
	{
		api.GET("stock", stockRatingController.GetAll)
		api.GET("stock/:id", stockRatingController.GetOne)
		api.POST("stock", stockRatingController.Post)
		api.PUT("stock/:id", stockRatingController.Put)
		api.DELETE("stock/:id", stockRatingController.Delete)
	}
}

