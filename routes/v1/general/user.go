package general

import (
	"matar/controllers/userController"

	"github.com/gin-gonic/gin"
)

func UserRoute(routerGroup *gin.RouterGroup) {
	users := routerGroup.Group("/users")
	users.POST("/", userController.CreateUser())
	users.POST("/login", userController.LoginUser())
}
