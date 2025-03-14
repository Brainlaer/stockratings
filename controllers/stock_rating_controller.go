package controllers

import (
	"example/hello/services"
	"example/hello/utils"
	"net/http"
	"github.com/gin-gonic/gin"
)

type StockController struct {
	Serv *services.StockService
}

func NewStockController(serv *services.StockService)*StockController{
	return &StockController{Serv: serv}
}

// swagger:route GET /stock stock getStocks
//GetStocks returns all stock
//responses:
//	200: ResponseSuccess
//  400: ResponseError
//  500: ResponseError
func (c *StockController) GetAll(ctx *gin.Context){
	var response utils.Response=*c.Serv.GetAll(ctx)
	ctx.IndentedJSON(http.StatusOK, response)
}

// swagger:route GET /stock/{id} stock getStock
//GetOneStock returns one from the stock
//responses:
//	200: ResponseSuccess
//  400: ResponseError
//  500: ResponseError
func (c *StockController) GetOne(ctx *gin.Context){
	var response utils.Response=*c.Serv.GetOne(ctx)
	ctx.IndentedJSON(http.StatusOK, response)
}

// swagger:route POST /stock stock createStock
//CreateStock returns a success message
//responses:
//	202: ResponseSuccess
//  400: ResponseError
//  500: ResponseError
func (c *StockController) Create(ctx *gin.Context){
	var response utils.Response=*c.Serv.Create(ctx)
	ctx.IndentedJSON(http.StatusOK,response)
}

// swagger:route PUT /stock/{id} stock updateStock
//updateStock returns a success message
//responses:
//	200: ResponseSuccess
//  400: ResponseError
//  500: ResponseError
func (c *StockController) Update(ctx *gin.Context){
	var response utils.Response=*c.Serv.Update(ctx)
	ctx.IndentedJSON(http.StatusOK,response)
}

// swagger:route DELETE /stock/{id} stock deleteStock
//Delete returns message
//responses:
//	200: ResponseSuccess
//  400: ResponseError
//  500: ResponseError
func (c *StockController) Delete(ctx *gin.Context){
	var response utils.Response=*c.Serv.Delete(ctx)
	ctx.IndentedJSON(http.StatusOK,response)

}
