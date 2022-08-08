package automobileAdService

import (
	"context"
	"errors"
	"fmt"
	"matar/clients"
	"matar/models/automobileAdModel"
	"matar/services/userService"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func GeAutomobileAdById(ctx context.Context, id string) (*automobileAdModel.AutomobileAd, error) {
	var automobileAd automobileAdModel.AutomobileAd
	var automobileAdCollection *mongo.Collection = clients.GetMongoCollection(clients.GetConnectedMongoClient(), automobileAdModel.AutomobileAdCollectionName)
	err := automobileAdCollection.FindOne(ctx, bson.M{"_id": id}).Decode(&automobileAd)
	if err != nil {
		fmt.Println(err)
		return nil, errors.New("Can not get Ad")
	}
	return &automobileAd, nil
}

func CreateAutomobileAd(ctx context.Context, automobileAd automobileAdModel.AutomobileAd) (*mongo.InsertOneResult, error) {
	fmt.Println()
	userClaims := ctx.Value(userService.UserClaims{})

	userId := userClaims.(*userService.UserClaims).Id.Hex()
	fmt.Println(userId)
	var automobileAdCollection *mongo.Collection = clients.GetMongoCollection(clients.GetConnectedMongoClient(), automobileAdModel.AutomobileAdCollectionName)
	user, err := userService.GetUserById(ctx, userId)
	if err != nil || user == nil {
		return nil, errors.New("User not found")
	}
	newAutomobileAd := automobileAdModel.AutomobileAd{
		Id:               primitive.NewObjectID(),
		Title:            automobileAd.Title,
		UserId:           automobileAd.UserId,
		Brand:            automobileAd.Brand,
		BodyType:         automobileAd.BodyType,
		Address:          automobileAd.Address,
		Model:            automobileAd.Model,
		Milage:           automobileAd.Milage,
		Price:            automobileAd.Price,
		Images:           automobileAd.Images,
		ContactNo:        automobileAd.ContactNo,
		FuelType:         automobileAd.FuelType,
		Color:            automobileAd.Color,
		Transmission:     automobileAd.Transmission,
		WheelDrive:       automobileAd.WheelDrive,
		UsageCondition:   automobileAd.UsageCondition,
		RegistrationCard: automobileAd.RegistrationCard,
		SellerComments:   automobileAd.SellerComments,
		SeatCapacity:     automobileAd.SeatCapacity,
		Specification:    automobileAd.Specification,
		Active:           true,
		CreatedAt:        time.Now(),
		UpdatedAt:        time.Now(),
	}

	return automobileAdCollection.InsertOne(ctx, newAutomobileAd)
}
