package main

import (
	"context"
	"matar/clients"
	"matar/configs"
	"matar/models/automobileAdModel"
	"matar/models/userModel"
	"matar/routes"
	"time"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client := clients.ConnectToMongoDB(ctx)
	userModel.CreateUserIndexes(ctx, client)
	automobileAdModel.CreateAutomobileAdIndexes(ctx, client)
	router.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"data": "ok",
		})
	})
	routes.Load(router)

	router.Run("localhost:" + configs.Common.Service.Port)
}
