package models

import (
	"context"
	"matar/clients"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var UserCollectionName = "users"

type location struct {
	Type        string     `json:"type"`
	Coordinates [2]float32 `json:"coordinates"`
}
type address struct {
	Country  string   `json:"country"`
	State    string   `json:"State"`
	City     string   `json:"City"`
	Location location `json:"location"`
}

type company struct {
	Name           string  `json:"name"`
	RegistrationNo string  `json:"registration_no" bson:"registration_no"`
	Address        address `json:"address"`
}

type User struct {
	Phone               string    `json:"phone" validate:"required"`
	Password            string    `json:"password" validate:"required"`
	Type                string    `json:"type" validate:"required,oneof=individual company"`
	Country             string    `json:"country" validate:"required,oneof=Morocco"`
	Email               string    `json:"email,omitempty"`
	PhoneNumberVerified bool      `json:"phone_number_verified" bson:"phone_number_verified"`
	EmailVerified       bool      `json:"email_number_verified" bson:"email_number_verified"`
	Active              bool      `json:"active"`
	Company             *company  `json:"company,omitempty"`
	CreatedAt           time.Time `json:"created_at" bson:"created_at"`
	UpdatedAt           time.Time `json:"updated_at" bson:"updated_at"`
}

func CreateUserIndexes(client *mongo.Client) {
	col := clients.GetMongoCollection(client, UserCollectionName)
	col.Indexes().CreateOne(context.Background(), mongo.IndexModel{
		Keys:    bson.M{"phone": 1},
		Options: options.Index().SetUnique(true),
	})
}
