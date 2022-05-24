package admin

import (
	"github.com/gin-gonic/gin"
)

func Load(routerGroup *gin.RouterGroup) {
	adminRoutes := routerGroup.Group("/admin")
	UserRoute(adminRoutes)
}
