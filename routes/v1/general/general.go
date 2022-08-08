package general

import (
	"github.com/gin-gonic/gin"
)

func Load(routerGroup *gin.RouterGroup) {
	generalRoutes := routerGroup.Group("/general")
	UserRoute(generalRoutes)
	AutomobileAdRoute(generalRoutes)
}
