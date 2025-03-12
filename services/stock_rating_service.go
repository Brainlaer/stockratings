package services

import (
	"example/hello/models"
	"example/hello/repositories"
	"example/hello/utils"
	"fmt"
	"reflect"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

type StockService struct {
	Repo *repositories.StockRatingRepository
}

func NewStockService(repo *repositories.StockRatingRepository) *StockService {
	return &StockService{Repo: repo}
}

func (c *StockService) GetAll(ctx *gin.Context) utils.Response {
	response := utils.Response{}
	filters,sortBy,order,limit,offset:=GetFilters(ctx)
	stocks, totalRecords, err := c.Repo.GetAll(filters,sortBy,order, limit, offset)

	if err != nil {
		response.Status = "500"
		response.Error.Code = "DATABASE_ERROR"
		response.Error.Details = err.Error()
		return response
	}
	response.Status = "200"
	response.Data = stocks
	response.Meta = map[string]interface{}{
		"limit":  limit,
		"offset": offset,
		"totalRecords": totalRecords,
	}
	return response

}

func (c *StockService) GetOne(ctx *gin.Context) utils.Response {
	response := utils.Response{}
	stock, err := c.Repo.GetOne(ctx.Param("id"))

	if err != nil {
		response.Status = "500"
		response.Error.Code = "DATABASE_ERROR"
		response.Error.Details = err.Error()
		return response
	}
	response.Status = "200"
	response.Data = stock
	return response
}

func (c *StockService) Create(ctx *gin.Context) *utils.Response {
	var response utils.Response

	stockRating, decodeResponse := DecodeJson(ctx)
	if decodeResponse != nil {
		return decodeResponse
	}
	stockTime, decodeResponse:=ParseTextToTime(stockRating.Time)
	if decodeResponse != nil {
		return decodeResponse
	}

	err := c.Repo.Create(stockRating,stockTime)
	if err != nil {
		response = utils.Response{
			Status: "500",
			Error: utils.ResponseError{
				Code:    "DATABASE_ERROR",
				Details: err.Error(),
			},
		}
		return &response
	}
	response = utils.Response{
		Status: "202",
		Data:   "Creado satisfactoriamente",
	}

	return &response
}

func (c *StockService) Update(ctx *gin.Context) *utils.Response {
	var response *utils.Response
	newStockRating := &models.StockRatingUpdateRequest{}

	newStockRating.Body, response = DecodeJson(ctx)
	if response != nil {
		return response
	}

	stockRating, err := c.Repo.GetOne(ctx.Param("id"))
	if err != nil {
		response = &utils.Response{
				Status: "500",
				Error: utils.ResponseError{
					Code:    "DATABASE_ERROR",
					Details: err.Error(),
				},
		}
		return response
	}

	UpdateValues(stockRating, newStockRating.Body)

	err = c.Repo.Update(stockRating)
	if err != nil {
		return &utils.Response{
				Status: "500",
				Error: utils.ResponseError{
					Code:    "UPDATE_ERROR",
					Details: err.Error(),
				},
		}
	}

	return &utils.Response{

			Status: "200",
			Data:   stockRating,

	}

}


func (c *StockService) Delete(ctx *gin.Context) *utils.Response{
	var response utils.Response
	err:=c.	Repo.Delete(ctx.Param("id"))
	if err != nil {
		response = utils.Response{
			Status: "500",
			Error: utils.ResponseError{
				Code:    "DATABASE_ERROR",
				Details: err.Error(),
			},
		}
		return &response 
	}
	response = utils.Response{
		Status: "200",
		Data:   "eliminado satisfactoriamente",
	}

	return &response
}

func DecodeJson(ctx *gin.Context) (*models.StockRatingCreate, *utils.Response) {
	var response utils.Response
	var stockRating models.StockRatingCreate

	if err := ctx.ShouldBind(&stockRating); err != nil {
		response = utils.Response{
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

func ParseTextToTime(text string)(*time.Time,*utils.Response){
	var response utils.Response
	fmt.Print(text)
	value:=text

	dateTime,err:=time.Parse(time.RFC3339, value)
	if err != nil {
		response = utils.Response{
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


func UpdateValues(stockRating *models.StockRatingGet , newStockRating *models.StockRatingCreate) {
	mockStockRating:=models.NewStockRatingCreate(newStockRating,stockRating.ID)
	v1 := reflect.ValueOf(stockRating).Elem()
	V2 := reflect.ValueOf(mockStockRating).Elem()

	for i :=0; i <v1.NumField(); i++{
		field1:= v1.Field(i)
		field2:= V2.Field(i)

		if field1.CanSet() && !reflect.DeepEqual(field1.Interface(), field2.Interface()){
			field1.Set(field2)
		}
	}
}

func GetFilters(ctx *gin.Context)(map[string]string, []string, []string, int, int){
	filters := map[string]string{}
	for _, key := range []string{"ticker", "company", "action", "brokerage", "rating_from", "rating_to"} {
		if value := ctx.Query(key); value != "" {
			filters[key] = value
		}
	}

	sortBy := strings.Split(ctx.DefaultQuery("sortBy", "time"), ",") 
	order := strings.Split(ctx.DefaultQuery("order", "desc"), ",")   
	limit, err := strconv.Atoi(ctx.DefaultQuery("limit", "10")) 
	if err != nil || limit <= 0 {
		limit = 10
	}

	offset, err := strconv.Atoi(ctx.DefaultQuery("offset", "0"))
	if err != nil || offset < 0{
		offset =0
	}
	return filters,sortBy,order, limit, offset
}