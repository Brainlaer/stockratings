package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// swagger:route GET /stock getAllStock
//
//GetStock returns all stock
//
//responses:
//
//	200: successResponse
func GetAll(c *gin.Context){
	c.IndentedJSON(http.StatusOK,"STOCK")
}

// swagger:route POST /stock getAllStock
//
//CreateStock returns message
//
//responses:
//
//	202: acceptedResponse
func Post(c *gin.Context){
	 
	c.IndentedJSON(http.StatusAccepted,"Created")
}

// swagger:route PUT /stock updateStock
//
//updateStock returns message
//
//responses:
//
//	200: successResponse
func Put(c *gin.Context){
	 
	c.IndentedJSON(http.StatusOK,"Updated")
}

// swagger:route DELETE /stock delete
//
//Delete returns message
//
//responses:
//
//	200: successResponse
func Delete(c *gin.Context){
	c.IndentedJSON(http.StatusOK,"deleted")

}
