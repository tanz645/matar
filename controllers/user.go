package controllers

import (
	"context"
	"matar/clients"
	"matar/common/responses"
	"matar/models"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.mongodb.org/mongo-driver/mongo"
)

var validate = validator.New()

func CreateUser() gin.HandlerFunc {
	var userCollection *mongo.Collection = clients.GetMongoCollection(clients.GetConnectedMongoClient(), models.UserCollectionName)

	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		var user models.User
		defer cancel()

		//validate the request body
		if err := c.BindJSON(&user); err != nil {
			c.JSON(http.StatusBadRequest, responses.FailedResponse{Status: http.StatusBadRequest, Error: true, Message: "User can not be created", Data: err.Error()})
			return
		}

		//use the validator library to validate required fields
		if validationErr := validate.Struct(&user); validationErr != nil {
			c.JSON(http.StatusBadRequest, responses.FailedResponse{Status: http.StatusBadRequest, Error: true, Message: "User can not be created", Data: validationErr.Error()})
			return
		}

		newUser := models.User{
			// Id:    primitive.NewObjectID(),
			Name:  user.Name,
			Phone: user.Phone,
			Email: user.Email,
		}

		result, err := userCollection.InsertOne(ctx, newUser)
		if err != nil {
			c.JSON(http.StatusInternalServerError, responses.FailedResponse{Status: http.StatusInternalServerError, Error: true, Message: "User can not be created", Data: err.Error()})
			return
		}

		c.JSON(http.StatusCreated, responses.SuccessResponse{Status: http.StatusCreated, Success: true, Message: "User created", Data: result})
	}
}
