package controllers

import (
	"context"
	"matar/common/responses"
	"matar/models"
	"matar/services/userService"

	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var validate = validator.New()

func CreateUser() gin.HandlerFunc {

	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		var user models.User
		defer cancel()
		if err := c.BindJSON(&user); err != nil {
			c.JSON(http.StatusBadRequest, responses.FailedResponse{Status: http.StatusBadRequest, Error: true, Message: "User can not be created", Data: err.Error()})
			return
		}
		if validationErr := validate.Struct(&user); validationErr != nil {
			c.JSON(http.StatusUnprocessableEntity, responses.FailedResponse{Status: http.StatusUnprocessableEntity, Error: true, Message: "User can not be created", Data: validationErr.Error()})
			return
		}
		result, err := userService.CreateUser(ctx, user)
		if err != nil {
			c.JSON(http.StatusInternalServerError, responses.FailedResponse{Status: http.StatusInternalServerError, Error: true, Message: "User can not be created", Data: err.Error()})
			return
		}
		data := map[string]string{
			"id": result.InsertedID.(primitive.ObjectID).Hex(),
		}
		c.JSON(http.StatusCreated, responses.SuccessResponse{Status: http.StatusCreated, Success: true, Message: "User created", Data: data})
	}
}
