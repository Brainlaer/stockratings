package controllers

import (
	"example/hello/services"
	"example/hello/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

type StockRatingController struct {
	Serv *services.StockService
}

func NewStockRatingController(serv *services.StockService)*StockRatingController{
	return &StockRatingController{Serv: serv}
}
	
// swagger:route GET /stock getAllStock
//
//GetStock returns all stock
//
//responses:
//
//	200: Response
func (c *StockRatingController) GetAll(ctx *gin.Context){
	var response utils.Response=c.Serv.GetAll()
	ctx.IndentedJSON(http.StatusOK, response)
}

// swagger:route POST /stock getAllStock
//
//CreateStock returns message
//
//responses:
//
//	202: Response
func (c *StockRatingController) Post(ctx *gin.Context){
	 
	ctx.IndentedJSON(http.StatusAccepted,"Created")
}

// swagger:route PUT /stock updateStock
//
//updateStock returns message
//
//responses:
//
//	200: Response
func (c *StockRatingController) Put(ctx *gin.Context){
	 
	ctx.IndentedJSON(http.StatusOK,"Updated")
}

// swagger:route DELETE /stock delete
//
//Delete returns message
//
//responses:
//
//	200: Response
func (c *StockRatingController) Delete(ctx *gin.Context){
	ctx.IndentedJSON(http.StatusOK,"deleted")

}
