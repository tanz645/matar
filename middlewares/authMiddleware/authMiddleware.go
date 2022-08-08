package authMiddleware

import (
	"matar/common/responses"
	"matar/services/userService"
	"net/http"

	"github.com/gin-gonic/gin"
)

func VerifyUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.Request.Header["Authorization"]
		claims, _ := userService.VerifyToken(token[0])
		if claims == nil {
			c.JSON(http.StatusUnauthorized, responses.FailedResponse{Status: http.StatusUnauthorized, Error: true, Message: "Not Allowed", Data: nil})
			c.Abort()
			return
		}
		c.Set("user", claims)
		c.Next()
	}
}
