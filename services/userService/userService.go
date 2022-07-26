package userService

import (
	"context"
	"matar/clients"
	"matar/models"

	"go.mongodb.org/mongo-driver/mongo"
)

func CreateUser(ctx context.Context, user models.User) (*mongo.InsertOneResult, error) {
	var userCollection *mongo.Collection = clients.GetMongoCollection(clients.GetConnectedMongoClient(), models.UserCollectionName)

	newUser := models.User{
		Name:  user.Name,
		Phone: user.Phone,
		Email: user.Email,
	}

	return userCollection.InsertOne(ctx, newUser)
}
