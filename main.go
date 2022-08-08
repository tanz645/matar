package main

import (
	"matar/clients"
	"matar/configs"
	"matar/models/automobileAdModel"
	"matar/models/userModel"
	"matar/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	client := clients.ConnectToMongoDB()
	userModel.CreateUserIndexes(client)
	automobileAdModel.CreateAutomobileAdIndexes(client)
	router.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"data": "ok",
		})
	})
	routes.Load(router)

	router.Run("localhost:" + configs.Common.Service.Port)
}
