package admin

import (
	"github.com/gin-gonic/gin"
)

func UserRoute(routerGroup *gin.RouterGroup) {
	users := routerGroup.Group("/users")
	users.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"data": "some users for admin",
		})
	})
}
