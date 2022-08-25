package automobileAdController

import (
	"context"
	"fmt"
	"matar/common/responses"
	"matar/controllers"
	"matar/schemas/automobileAdSchema"
	"matar/services/automobileAdService"
	"matar/services/userService"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func UpdateAdActiveStatus() gin.HandlerFunc {

	return func(c *gin.Context) {
		ctx := context.Background()
		ctx = context.WithValue(ctx, userService.UserClaims{}, c.Value("user"))
		ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
		id := c.Param("id")

		var updateAutomobileAdActiveStatus automobileAdSchema.UpdateAutomobileAdActiveStatus
		defer cancel()
		if err := c.BindJSON(&updateAutomobileAdActiveStatus); err != nil {
			c.JSON(http.StatusBadRequest, responses.FailedResponse{Status: http.StatusBadRequest, Error: true, Message: "Ad can not be updated", Data: err.Error()})
			return
		}
		if validationErr := controllers.Validate.Struct(&updateAutomobileAdActiveStatus); validationErr != nil {
			c.JSON(http.StatusUnprocessableEntity, responses.FailedResponse{Status: http.StatusUnprocessableEntity, Error: true, Message: "Ad can not be updated", Data: validationErr.Error()})
			return
		}
		result, err := automobileAdService.UpdateAdActiveStatus(ctx, updateAutomobileAdActiveStatus, id)
		if err != nil || result == "" {
			fmt.Println(err)
			c.JSON(http.StatusInternalServerError, responses.FailedResponse{Status: http.StatusInternalServerError, Error: true, Message: "Ad can not be updated", Data: nil})
			return
		}
		data := map[string]string{
			"id": result,
		}
		c.JSON(http.StatusOK, responses.SuccessResponse{Status: http.StatusOK, Success: true, Message: "Ad updated", Data: data})
	}
}
