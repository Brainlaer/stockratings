package services

import (
	"example/hello/repositories"
	"example/hello/utils"
)

type StockService struct {
	Repo *repositories.StockRatingRepository
}

func NewStockService(repo *repositories.StockRatingRepository)*StockService{
	return &StockService{Repo: repo}
}
func (c *StockService) GetAll()utils.Response{
	response:= utils.Response{}
	stocks, err := c.Repo.GetAll()

	if err != nil{
		response.Body.Status="500"
		response.Body.Error.Code="DATABASE_ERROR"
		response.Body.Error.Details=err.Error()
		return response
	}
		response.Body.Status="200"
		response.Body.Data=stocks
		return response
	
}