package routes

import (
	v1 "matar/routes/v1"

	"github.com/gin-gonic/gin"
)

func Load(router *gin.Engine) {
	v1.Load(router)
}
