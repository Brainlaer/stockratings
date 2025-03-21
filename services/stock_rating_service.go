package services

import (
	"example/hello/models"
	"example/hello/repositories"
	"example/hello/utils"
	"github.com/gin-gonic/gin"
	"reflect"
	"strconv"
	"strings"
	"time"
)

type StockService struct {
	Repo *repositories.StockRepository
}

func NewStockService(repo *repositories.StockRepository) *StockService {
	return &StockService{Repo: repo}
}

func (s *StockService) GetAll(ctx *gin.Context) *utils.Response {
	var response utils.Response
	filters, sortBy, order, limit, offset, strongQuery := GetFilters(ctx)
	stocks, totalRecords, err := s.Repo.GetAll(filters, sortBy, order, limit, offset, strongQuery)

	if err != nil {
		response = utils.Response{
			Status: "404",
			Error: utils.ResponseError{
				Code:    "NOT_FOUND",
				Details: err.Error(),
			},
		}
		return &response
	}
	response = utils.Response{
		Status: "200",
		Success: utils.ResponseSuccess{
			Data:   stocks,
			Meta: map[string]interface{}{
				"limit":        limit,
				"offset":       offset,
				"totalRecords": totalRecords,
			},
		},
	}
	return &response

}

func (s *StockService) GetOne(ctx *gin.Context) *utils.Response {
	response := utils.Response{}
	stock, err := s.Repo.GetOne(ctx.Param("id"))

	if err != nil {
		response = utils.Response{
			Status: "404",
			Error: utils.ResponseError{
				Code:    "NOT_FOUND",
				Details: err.Error(),
			},
		}
		return &response
	}
	response = utils.Response{
		Status: "200",
		Success: utils.ResponseSuccess{
			Data:   stock,
		},
	}
	return &response
}

func (s *StockService) Create(ctx *gin.Context) *utils.Response {
	var response utils.Response

	stock, decodeResponse := DecodeJson(ctx)
	if decodeResponse != nil {
		return decodeResponse
	}
	if stock.Time == nil {
		stock.Time = new(time.Time)
	}

	err := s.Repo.Create(stock)
	if err != nil {
		response = utils.Response{
			Status: "400",
			Error: utils.ResponseError{
				Code:    "BAD_REQUEST",
				Details: err.Error(),
			},
		}
		return &response
	}
	response = utils.Response{
		Status: "202",
		Success: utils.ResponseSuccess{
			Data:   "Creado satisfactoriamente",
		},
	}

	return &response
}

func (s *StockService) Update(ctx *gin.Context) *utils.Response {
	var response *utils.Response
	newStock := &models.StockRequestUpdate{}
	id := ctx.Param("id")

	newStock.Body, response = DecodeJson(ctx)
	if response != nil {
		return response
	}

	stock, err := s.Repo.GetOne(id)
	if err != nil {
		response = &utils.Response{
			Status: "400",
			Error: utils.ResponseError{
				Code:    "BAD_REQUEST",
				Details: err.Error(),
			},
		}
		return response
	}

	UpdateValues(stock, newStock)

	err = s.Repo.Update(id, &stock.Stock)
	if err != nil {
		return &utils.Response{
			Status: "400",
			Error: utils.ResponseError{
				Code:    "BAD_REQUEST",
				Details: err.Error(),
			},
		}
	}

	return &utils.Response{

		Status: "200",
		Success: utils.ResponseSuccess{
			Data:   "Actualizado Correctamente",
		},
	}

}

func (s *StockService) Delete(ctx *gin.Context) *utils.Response {
	var response utils.Response
	id := ctx.Param("id")
	_, err := s.Repo.GetOne(id)
	if err != nil {
		response = utils.Response{
			Status: "404",
			Error: utils.ResponseError{
				Code:    "NOT_FOUND",
				Details: err.Error(),
			},
		}
		return &response
	}
	err = s.Repo.Delete(id)
	if err != nil {
		response = utils.Response{
			Status: "400",
			Error: utils.ResponseError{
				Code:    "BAD_REQUEST",
				Details: err.Error(),
			},
		}
		return &response
	}
	response = utils.Response{
		Status: "200",
		Success: utils.ResponseSuccess{
			Data:   "eliminado satisfactoriamente",
		},
	}

	return &response
}

func DecodeJson(ctx *gin.Context) (*models.Stock, *utils.Response) {
	var response utils.Response
	var stockRating models.Stock

	if err := ctx.ShouldBind(&stockRating); err != nil {
		response = utils.Response{
			Status: "400",
			Error: utils.ResponseError{
				Code:    "BAD_REQUEST",
				Details: err.Error(),
			},
		}
		return nil, &response
	}
	return &stockRating, nil
}

func UpdateValues(stockRating *models.StockResponseGet, newStockRating *models.StockRequestUpdate) {
	v1 := reflect.ValueOf(&stockRating.Stock).Elem()
	V2 := reflect.ValueOf(newStockRating.Body).Elem()

	for i := 0; i < v1.NumField(); i++ {
		field1 := v1.Field(i)
		field2 := V2.Field(i)

		if field1.CanSet() && !reflect.DeepEqual(field1.Interface(), field2.Interface()) {
			field1.Set(field2)
		}
	}
}

func GetFilters(ctx *gin.Context) (map[string]interface{}, []string, []string, int, int, bool) {
	filters := map[string]interface{}{}
	for _, key := range []string{"ticker", "company", "brokerage"} {
		if value := ctx.Query(key); value != "" {
			filters[key] = value
		}
	}

	if ratingTo := ctx.Query("rating_to"); ratingTo != "" {
		filters["rating_to"] = strings.Split(ratingTo, ",") 
	}
	if ratingFrom := ctx.Query("rating_from"); ratingFrom != "" {
		filters["rating_from"] = strings.Split(ratingFrom, ",")
	}
	if action := ctx.Query("action"); action != "" {
		filters["action"] = strings.Split(action, ",")
	}
	
	strongQuery, err:= strconv.ParseBool(ctx.DefaultQuery("strongQuery", "or"))
	if err != nil {
		strongQuery=true
	}
	sortBy := strings.Split(ctx.DefaultQuery("sortBy", "time"), ",")
	order := strings.Split(ctx.DefaultQuery("order", "desc"), ",")
	limit, err := strconv.Atoi(ctx.DefaultQuery("limit", "10"))
	if err != nil || limit <= 0 {
		limit = 10
	}

	offset, err := strconv.Atoi(ctx.DefaultQuery("offset", "0"))
	if err != nil || offset < 0 {
		offset = 0
	}
	return filters, sortBy, order, limit, offset, strongQuery
}
