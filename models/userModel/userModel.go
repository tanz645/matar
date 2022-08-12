package userModel

import (
	"context"
	"matar/clients"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
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
	Id                  primitive.ObjectID `json:"id,omitempty" bson:"_id"`
	Phone               string             `json:"phone" validate:"required,e164"`
	Password            string             `json:"password" validate:"required,max=40,min=6"`
	Type                string             `json:"type" validate:"required,oneof=individual company"`
	Country             string             `json:"country" validate:"required,oneof=Morocco"`
	Email               string             `json:"email,omitempty"`
	PhoneNumberVerified bool               `json:"phone_number_verified" bson:"phone_number_verified"`
	EmailVerified       bool               `json:"email_number_verified" bson:"email_number_verified"`
	Active              bool               `json:"active"`
	Company             *company           `json:"company,omitempty"`
	AdIds               []string           `json:"ad_ids" bson:"ad_ids"`
	MaxAd               int16              `json:"max_ad" bson:"max_ad"`
	CreatedAt           time.Time          `json:"created_at" bson:"created_at"`
	UpdatedAt           time.Time          `json:"updated_at" bson:"updated_at"`
}

type UserLogin struct {
	Phone    string `json:"phone" validate:"required,e164"`
	Password string `json:"password" validate:"required,max=40,min=6"`
}

func CreateUserIndexes(ctx context.Context, client *mongo.Client) {
	col := clients.GetMongoCollection(client, UserCollectionName)
	col.Indexes().CreateOne(ctx, mongo.IndexModel{
		Keys: bson.M{"phone": 1},
	})
}
