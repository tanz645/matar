package userService

import (
	"context"
	"errors"
	"matar/clients"
	"matar/models"
	"matar/utils"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
)

func CreateUser(ctx context.Context, user models.User) (*mongo.InsertOneResult, error) {
	var userCollection *mongo.Collection = clients.GetMongoCollection(clients.GetConnectedMongoClient(), models.UserCollectionName)

	hashed, err := utils.HashPassword(user.Password)
	if err != nil {
		return nil, errors.New("Error in password hashing")
	}
	newUser := models.User{
		Phone:               user.Phone,
		Password:            hashed,
		Type:                user.Type,
		Country:             user.Country,
		Email:               user.Email,
		PhoneNumberVerified: false,
		EmailVerified:       false,
		Active:              true,
		Company:             nil,
		CreatedAt:           time.Now(),
		UpdatedAt:           time.Now(),
	}

	return userCollection.InsertOne(ctx, newUser)
}
