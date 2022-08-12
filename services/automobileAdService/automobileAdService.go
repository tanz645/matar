package automobileAdService

import (
	"context"
	"errors"
	"fmt"
	"matar/clients"
	"matar/models/automobileAdModel"
	"matar/services/userService"
	"matar/utils/helper"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func GeAutomobileAdGeneralById(ctx context.Context, id string) (*automobileAdModel.AutomobileAdGeneral, error) {
	var automobileAd automobileAdModel.AutomobileAdGeneral
	var automobileAdCollection *mongo.Collection = clients.GetMongoCollection(clients.GetConnectedMongoClient(), automobileAdModel.AutomobileAdCollectionName)
	objId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}
	err = automobileAdCollection.FindOne(ctx, bson.M{"_id": objId}).Decode(&automobileAd)
	if err != nil {
		return nil, err
	}
	return &automobileAd, nil
}

func GetAutomobileAdsGeneralByUserId(ctx context.Context, userId string) ([]automobileAdModel.AutomobileAdGeneral, error) {
	var automobileAds []automobileAdModel.AutomobileAdGeneral
	var automobileAdCollection *mongo.Collection = clients.GetMongoCollection(clients.GetConnectedMongoClient(), automobileAdModel.AutomobileAdCollectionName)
	results, err := automobileAdCollection.Find(ctx, bson.M{"user_id": userId})
	if err != nil {
		return nil, err
	}
	for results.Next(ctx) {
		var automobileAd automobileAdModel.AutomobileAdGeneral
		if err = results.Decode(&automobileAd); err != nil {
			return nil, err
		}

		automobileAds = append(automobileAds, automobileAd)
	}
	return automobileAds, nil
}

func GetCountAutomobileAdsGeneralByUserId(ctx context.Context, userId string) (int64, error) {
	var automobileAdCollection *mongo.Collection = clients.GetMongoCollection(clients.GetConnectedMongoClient(), automobileAdModel.AutomobileAdCollectionName)
	count, err := automobileAdCollection.CountDocuments(ctx, bson.M{"user_id": userId})
	if err != nil {
		fmt.Println(err)
		return 0, err
	}

	return count, nil
}

func DeleteAutomobileAdById(ctx context.Context, id string) error {
	var automobileAdCollection *mongo.Collection = clients.GetMongoCollection(clients.GetConnectedMongoClient(), automobileAdModel.AutomobileAdCollectionName)
	objId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}
	_, err = automobileAdCollection.DeleteOne(ctx, bson.M{"_id": objId})
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}

func CreateAutomobileAd(ctx context.Context, automobileAd automobileAdModel.AutomobileAd) (string, error) {
	var automobileAdCollection *mongo.Collection = clients.GetMongoCollection(clients.GetConnectedMongoClient(), automobileAdModel.AutomobileAdCollectionName)
	userClaims := ctx.Value(userService.UserClaims{})
	userId := userClaims.(*userService.UserClaims).Id.Hex()
	user, err := userService.GetUserById(ctx, userId)
	if err != nil || user == nil {
		return "", err
	}
	totalAutomobileAds, err := GetCountAutomobileAdsGeneralByUserId(ctx, userId)
	if err != nil || user == nil {
		return "", err
	}
	if totalAutomobileAds >= int64(user.MaxAd) {
		return "", errors.New("can not exceed max ad per account")
	}
	newAutomobileAd := automobileAdModel.AutomobileAd{
		Id:               primitive.NewObjectID(),
		Title:            automobileAd.Title,
		UserId:           userId,
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

	saveAdResult, err := automobileAdCollection.InsertOne(ctx, newAutomobileAd)
	if err != nil {
		return "", err
	}
	insertedId := saveAdResult.InsertedID.(primitive.ObjectID).Hex()
	user, err = userService.PushAdId(ctx, userId, insertedId)
	if err != nil {
		DeleteAutomobileAdById(ctx, insertedId)
		return "", err
	}
	return insertedId, nil
}

func UpdateAutomobileAdById(ctx context.Context, automobileAd automobileAdModel.AutomobileAd, adId string) (string, error) {
	var automobileAdCollection *mongo.Collection = clients.GetMongoCollection(clients.GetConnectedMongoClient(), automobileAdModel.AutomobileAdCollectionName)
	userClaims := ctx.Value(userService.UserClaims{})
	userId := userClaims.(*userService.UserClaims).Id.Hex()
	user, err := userService.GetUserById(ctx, userId)
	if err != nil || user == nil {
		return "", err
	}
	if !helper.Contains(user.AdIds, adId) {
		return "", errors.New("Can not update ad")
	}

	update := bson.M{
		"title":             automobileAd.Title,
		"brand":             automobileAd.Brand,
		"body_type":         automobileAd.BodyType,
		"address":           automobileAd.Address,
		"model":             automobileAd.Model,
		"milage":            automobileAd.Milage,
		"price":             automobileAd.Price,
		"images":            automobileAd.Images,
		"contactNo":         automobileAd.ContactNo,
		"fuel_type":         automobileAd.FuelType,
		"color":             automobileAd.Color,
		"transmission":      automobileAd.Transmission,
		"wheel_drive":       automobileAd.WheelDrive,
		"usage_condition":   automobileAd.UsageCondition,
		"registration_card": automobileAd.RegistrationCard,
		"seller_comments":   automobileAd.SellerComments,
		"seat_capacity":     automobileAd.SeatCapacity,
		"specification":     automobileAd.Specification,
		"updated_at":        time.Now(),
	}
	objId, err := primitive.ObjectIDFromHex(adId)
	if err != nil {
		return "", err
	}
	saveAdResult, err := automobileAdCollection.UpdateOne(ctx, bson.M{"_id": objId}, bson.M{"$set": update})
	fmt.Println(saveAdResult.ModifiedCount)
	if err != nil {
		return "", err
	}
	return adId, nil
}

func RemoveAutomobileAdGenera(ctx context.Context, adId string) error {
	userClaims := ctx.Value(userService.UserClaims{})
	userId := userClaims.(*userService.UserClaims).Id.Hex()
	user, err := userService.GetUserById(ctx, userId)
	if err != nil || user == nil {
		return err
	}
	if !helper.Contains(user.AdIds, adId) {
		return errors.New("Can not delete ad")
	}
	user, err = userService.RemoveAdId(ctx, userId, adId)
	if err == nil {
		err = DeleteAutomobileAdById(ctx, adId)
		if err != nil {
			user, err = userService.PushAdId(ctx, userId, adId)
			if err != nil {
				return err
			}
			return nil
		}
	}
	return err
}
