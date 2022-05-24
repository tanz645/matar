package v1

import (
	"matar/routes/v1/admin"
	"matar/routes/v1/general"

	"github.com/gin-gonic/gin"
)

func Load(router *gin.Engine) {
	v1RouterGroup := router.Group("/v1")
	admin.Load(v1RouterGroup)
	general.Load(v1RouterGroup)
}
