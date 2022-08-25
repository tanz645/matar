package userController

import (
	"context"
	"matar/common/responses"
	"matar/controllers"
	"matar/schemas/userSchema"
	"matar/services/userService"
	"matar/utils/helper"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func CreateUser() gin.HandlerFunc {

	return func(c *gin.Context) {
		ctx := context.Background()
		ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
		var user userSchema.User
		defer cancel()
		if err := c.BindJSON(&user); err != nil {
			c.JSON(http.StatusBadRequest, responses.FailedResponse{Status: http.StatusBadRequest, Error: true, Message: "User can not be created", Data: err.Error()})
			return
		}
		if validationErr := controllers.Validate.Struct(&user); validationErr != nil {
			c.JSON(http.StatusUnprocessableEntity, responses.FailedResponse{Status: http.StatusUnprocessableEntity, Error: true, Message: "User can not be created", Data: validationErr.Error()})
			return
		}
		result, err := userService.CreateUser(ctx, user)
		if err != nil {
			if helper.IsDup(err) {
				c.JSON(http.StatusBadRequest, responses.FailedResponse{Status: http.StatusBadRequest, Error: true, Message: "Phone number already exists", Data: nil})
				return
			}
			c.JSON(http.StatusInternalServerError, responses.FailedResponse{Status: http.StatusInternalServerError, Error: true, Message: "User can not be created", Data: nil})
			return
		}
		data := map[string]string{
			"id": result.InsertedID.(primitive.ObjectID).Hex(),
		}
		c.JSON(http.StatusCreated, responses.SuccessResponse{Status: http.StatusCreated, Success: true, Message: "User created", Data: data})
	}
}

func LoginUser() gin.HandlerFunc {

	return func(c *gin.Context) {
		ctx := context.Background()
		ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
		var userLogin userSchema.UserLogin
		defer cancel()
		if err := c.BindJSON(&userLogin); err != nil {
			c.JSON(http.StatusBadRequest, responses.FailedResponse{Status: http.StatusBadRequest, Error: true, Message: "User can not be created", Data: err.Error()})
			return
		}
		if validationErr := controllers.Validate.Struct(&userLogin); validationErr != nil {
			c.JSON(http.StatusUnprocessableEntity, responses.FailedResponse{Status: http.StatusUnprocessableEntity, Error: true, Message: "User can not be created", Data: validationErr.Error()})
			return
		}
		result, err := userService.LoginUser(ctx, userLogin)
		if err != nil {
			c.JSON(http.StatusBadRequest, responses.FailedResponse{Status: http.StatusBadRequest, Error: true, Message: "User can not be created", Data: err.Error()})
			return
		}
		data := map[string]*string{
			"token": result,
		}
		c.JSON(http.StatusCreated, responses.SuccessResponse{Status: http.StatusCreated, Success: true, Message: "Login success", Data: data})
	}
}
