package models

import (
	"context"
	"matar/clients"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var UserCollectionName = "users"

type User struct {
	Name  string `json:"name,omitempty" validate:"required"`
	Phone string `json:"phone,omitempty" validate:"required"`
	Email string `json:"email"`
}

func CreateUserIndexes(client *mongo.Client) {
	col := clients.GetMongoCollection(client, UserCollectionName)
	col.Indexes().CreateOne(context.Background(), mongo.IndexModel{
		Keys:    bson.M{"phone": 1},
		Options: options.Index().SetUnique(true),
	})
}
