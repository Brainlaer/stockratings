package services

import (
	"example/hello/models"
	"example/hello/repositories"
	"example/hello/utils"
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
)

type StockService struct {
	Repo *repositories.StockRatingRepository
}

func NewStockService(repo *repositories.StockRatingRepository) *StockService {
	return &StockService{Repo: repo}
}
func (c *StockService) GetAll() utils.Response {
	response := utils.Response{}
	stocks, err := c.Repo.GetAll()

	if err != nil {
		response.Body.Status = "500"
		response.Body.Error.Code = "DATABASE_ERROR"
		response.Body.Error.Details = err.Error()
		return response
	}
	response.Body.Status = "200"
	response.Body.Data = stocks
	return response

}

func (c *StockService) Create(ctx *gin.Context) *utils.Response {
	var response utils.Response

	stockRating, decodeResponse := decodeJson(ctx)
	if decodeResponse != nil {
		return decodeResponse
	}
	stockTime, decodeResponse:=parseTextToTime(stockRating.Time)
	if decodeResponse != nil {
		return decodeResponse
	}

	err := c.Repo.Create(stockRating,stockTime)
	if err != nil {
		response.Body = utils.ResponseBody{
			Status: "500",
			Error: utils.ResponseError{
				Code:    "DATABASE_ERROR",
				Details: err.Error(),
			},
		}
		return &response
	}
	response.Body = utils.ResponseBody{
		Status: "202",
		Data:   "Creado satisfactoriamente",
	}

	return &response
}

func (c *StockService) Delete(ctx *gin.Context) *utils.Response{
	var response utils.Response
	err:=c.	Repo.Delete(ctx.Param("id"))
	if err != nil {
		response.Body = utils.ResponseBody{
			Status: "500",
			Error: utils.ResponseError{
				Code:    "DATABASE_ERROR",
				Details: err.Error(),
			},
		}
		return &response 
	}
	response.Body = utils.ResponseBody{
		Status: "200",
		Data:   "eliminado satisfactoriamente",
	}

	return &response
}

func decodeJson(ctx *gin.Context) (*models.StockRatingCreate, *utils.Response) {
	var response utils.Response
	var stockRating models.StockRatingCreate

	if err := ctx.ShouldBind(&stockRating); err != nil {
		response.Body = utils.ResponseBody{
			Status: "400",
			Error: utils.ResponseError{
				Code:    "BAD_REQUEST",
				Details: err.Error(),
			},
		}
		return nil,&response
	}
	return &stockRating, nil
}

func parseTextToTime(text string)(*time.Time,*utils.Response){
	var response utils.Response
	fmt.Print(text)
	value:=text

	dateTime,err:=time.Parse(time.RFC3339, value)
	if err != nil {
		response.Body = utils.ResponseBody{
			Status: "400",
			Error: utils.ResponseError{
				Code:    "ERROR_WHEN_PARSING_TEXT",
				Details: err.Error(),
			},
		}
		return nil,&response
	}
	return &dateTime,nil

}
