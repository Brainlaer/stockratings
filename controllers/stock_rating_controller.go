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
	
// swagger:route GET /stock stock GetStock
//
//GetStock returns all stock
//
//responses:
//
//	200: Response
func (c *StockRatingController) GetAll(ctx *gin.Context){
	var response utils.Response=c.Serv.GetAll(ctx)
	ctx.IndentedJSON(http.StatusOK, response)
}

// swagger:route GET /stock/{id} stock id
//
//GetOneStock returns one from the stock
//
//responses:
//
//	200: Response
func (c *StockRatingController) GetOne(ctx *gin.Context){
	var response utils.Response=c.Serv.GetOne(ctx)
	ctx.IndentedJSON(http.StatusOK, response)
}

// swagger:route POST /stock stock createStock
//
//CreateStock returns message
//
//responses:
//
//	200: Response
func (c *StockRatingController) Post(ctx *gin.Context){
	response:=c.Serv.Create(ctx)
	ctx.IndentedJSON(http.StatusOK,response)
}

// swagger:route PUT /stock/{id} stock updateStock
//
//updateStock returns message
//
//responses:
//
//	200: Response
func (c *StockRatingController) Put(ctx *gin.Context){
	response:=c.Serv.Update(ctx)
	ctx.IndentedJSON(http.StatusOK,response)
}

// swagger:route DELETE /stock/{id} stock id
//
//Delete returns message
//
//responses:
//
//	200: Response
func (c *StockRatingController) Delete(ctx *gin.Context){
	response:=c.Serv.Delete(ctx)
	ctx.IndentedJSON(http.StatusOK,response)

}
