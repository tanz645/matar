package controllers

import (
	"context"
	"matar/common/responses"
	"matar/models/automobileAdModel"
	"matar/services/automobileAdService"
	"matar/services/userService"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func CreateAutomobileAd() gin.HandlerFunc {

	return func(c *gin.Context) {
		ctx := context.Background()
		ctx = context.WithValue(ctx, userService.UserClaims{}, c.Value("user"))
		ctx, cancel := context.WithTimeout(ctx, 10*time.Second)

		var automobileAd automobileAdModel.AutomobileAd
		defer cancel()
		if err := c.BindJSON(&automobileAd); err != nil {
			c.JSON(http.StatusBadRequest, responses.FailedResponse{Status: http.StatusBadRequest, Error: true, Message: "Ad can not be created", Data: err.Error()})
			return
		}
		// if validationErr := validate.Struct(&automobileAd); validationErr != nil {
		// 	c.JSON(http.StatusUnprocessableEntity, responses.FailedResponse{Status: http.StatusUnprocessableEntity, Error: true, Message: "Ad can not be created", Data: validationErr.Error()})
		// 	return
		// }
		result, err := automobileAdService.CreateAutomobileAd(ctx, automobileAd)
		if err != nil {
			c.JSON(http.StatusInternalServerError, responses.FailedResponse{Status: http.StatusInternalServerError, Error: true, Message: "Ad can not be created", Data: nil})
			return
		}
		data := map[string]string{
			"id": result.InsertedID.(primitive.ObjectID).Hex(),
		}
		c.JSON(http.StatusCreated, responses.SuccessResponse{Status: http.StatusCreated, Success: true, Message: "Ad created", Data: data})
	}
}
