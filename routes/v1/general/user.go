package general

import (
	"matar/controllers"

	"github.com/gin-gonic/gin"
)

func UserRoute(routerGroup *gin.RouterGroup) {
	users := routerGroup.Group("/users")
	users.POST("/", controllers.CreateUser())
	users.POST("/login", controllers.LoginUser())
}
