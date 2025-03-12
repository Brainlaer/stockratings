package main

import (
	"example/hello/routes"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	
  r := gin.Default()
  r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Referer", "Content-Type", "Accept"},
		AllowCredentials: true,
	}))
  routes.SetupRoutes(r)
  r.Run(":8081") 
}