package automobileAdService

import (
	"context"
	"errors"
	"fmt"
	"matar/clients"
	"matar/common/enum"
	"matar/common/responses"
	"matar/schemas/automobileAdSchema"
	"matar/services/mediaService"
	"matar/services/userService"
	"matar/utils/helper"
	"mime/multipart"
	"strings"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func GetAutomobileAdGeneralById(ctx context.Context, id string) (*automobileAdSchema.AutomobileAdGeneral, error) {
	var automobileAd automobileAdSchema.AutomobileAdGeneral
	var automobileAdCollection *mongo.Collection = clients.GetMongoCollection(clients.GetConnectedMongoClient(), automobileAdSchema.AutomobileAdCollectionName)
	objId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}
	err = automobileAdCollection.FindOne(ctx, bson.D{{Key: "_id", Value: objId}, {Key: "active", Value: true}}).Decode(&automobileAd)
	if err != nil {
		return nil, err
	}
	return &automobileAd, nil
}

func GetAutomobileAdGeneralByUserId(ctx context.Context, id string) (*automobileAdSchema.AutomobileAdGeneral, error) {
	var automobileAd automobileAdSchema.AutomobileAdGeneral
	var automobileAdCollection *mongo.Collection = clients.GetMongoCollection(clients.GetConnectedMongoClient(), automobileAdSchema.AutomobileAdCollectionName)
	userClaims := ctx.Value(userService.UserClaims{})
	userId := userClaims.(*userService.UserClaims).Id.Hex()

	user, err := userService.GetUserById(ctx, userId)
	if err != nil || user == nil {
		return nil, err
	}
	objId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}
	err = automobileAdCollection.FindOne(ctx, bson.D{{Key: "_id", Value: objId}, {Key: "user_id", Value: userId}}).Decode(&automobileAd)
	if err != nil {
		return nil, err
	}
	return &automobileAd, nil
}

func GetAutomobileAdsGeneralByUserId(ctx context.Context) ([]automobileAdSchema.AutomobileAdGeneral, error) {
	var automobileAds []automobileAdSchema.AutomobileAdGeneral
	var automobileAdCollection *mongo.Collection = clients.GetMongoCollection(clients.GetConnectedMongoClient(), automobileAdSchema.AutomobileAdCollectionName)
	userClaims := ctx.Value(userService.UserClaims{})
	userId := userClaims.(*userService.UserClaims).Id.Hex()

	user, err := userService.GetUserById(ctx, userId)
	if err != nil || user == nil {
		return nil, err
	}
	results, err := automobileAdCollection.Find(ctx, bson.M{"user_id": userId})
	if err != nil {
		return nil, err
	}
	for results.Next(ctx) {
		var automobileAd automobileAdSchema.AutomobileAdGeneral
		if err = results.Decode(&automobileAd); err != nil {
			return nil, err
		}

		automobileAds = append(automobileAds, automobileAd)
	}
	return automobileAds, nil
}

func GetCountAutomobileAdsGeneralByUserId(ctx context.Context, userId string) (int64, error) {
	var automobileAdCollection *mongo.Collection = clients.GetMongoCollection(clients.GetConnectedMongoClient(), automobileAdSchema.AutomobileAdCollectionName)
	count, err := automobileAdCollection.CountDocuments(ctx, bson.M{"user_id": userId})
	if err != nil {
		fmt.Println(err)
		return 0, err
	}

	return count, nil
}

func DeleteAutomobileAdById(ctx context.Context, id string) error {
	var automobileAdCollection *mongo.Collection = clients.GetMongoCollection(clients.GetConnectedMongoClient(), automobileAdSchema.AutomobileAdCollectionName)
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

func CreateAutomobileAd(ctx context.Context, automobileAd automobileAdSchema.AutomobileAd) (string, error) {
	var automobileAdCollection *mongo.Collection = clients.GetMongoCollection(clients.GetConnectedMongoClient(), automobileAdSchema.AutomobileAdCollectionName)
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
	newAutomobileAd := automobileAdSchema.AutomobileAd{
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

func UpdateAutomobileAdById(ctx context.Context, automobileAd automobileAdSchema.AutomobileAd, adId string) (string, error) {
	var automobileAdCollection *mongo.Collection = clients.GetMongoCollection(clients.GetConnectedMongoClient(), automobileAdSchema.AutomobileAdCollectionName)
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

func SearchAutomobileAdGeneral(ctx context.Context, query automobileAdSchema.SearchAutomobileAdGeneral) (*responses.ListingResponse, error) {
	var automobileAds []automobileAdSchema.AutomobileAdInListing = []automobileAdSchema.AutomobileAdInListing{}
	var listingResponse responses.ListingResponse
	listingResponse.TotalCount = 0
	listingResponse.List = automobileAds
	var automobileAdCollection *mongo.Collection = clients.GetMongoCollection(clients.GetConnectedMongoClient(), automobileAdSchema.AutomobileAdCollectionName)
	limit := int64(query.Limit)
	page := int64(query.Page)
	sort := bson.D{{Key: "milage", Value: 1}}
	skip := int64(page*limit - limit)
	match := bson.D{{Key: "active", Value: true}}
	projection := bson.D{
		{Key: "_id", Value: 1},
		{Key: "title", Value: 1},
		{Key: "brand", Value: 1},
		{Key: "body_type", Value: 1},
		{Key: "address", Value: 1},
		{Key: "model", Value: 1},
		{Key: "milage", Value: 1},
		{Key: "price", Value: 1},
		{Key: "images", Value: 1},
		{Key: "fuel_type", Value: 1},
		{Key: "color", Value: 1},
		{Key: "transmission", Value: 1},
		{Key: "wheel_drive", Value: 1},
		{Key: "created_at", Value: 1},
	}
	if len(query.CityName) > 0 {
		cities := strings.Split(query.CityName, ",")
		match = append(match, bson.E{Key: "address.city", Value: bson.D{{Key: "$in", Value: cities}}})
	}
	if len(query.BrandName) > 0 {
		brands := strings.Split(query.BrandName, ",")
		match = append(match, bson.E{Key: "brand.name", Value: bson.D{{Key: "$in", Value: brands}}})
	}
	if len(query.BodyType) > 0 {
		bodyTypes := strings.Split(query.BodyType, ",")
		match = append(match, bson.E{Key: "body_type", Value: bson.D{{Key: "$in", Value: bodyTypes}}})
	}
	if len(query.Transmission) > 0 {
		transmissions := strings.Split(query.Transmission, ",")
		match = append(match, bson.E{Key: "transmission", Value: bson.D{{Key: "$in", Value: transmissions}}})
	}
	if len(query.FuelType) > 0 {
		fuelTypes := strings.Split(query.FuelType, ",")
		match = append(match, bson.E{Key: "fuel_type", Value: bson.D{{Key: "$in", Value: fuelTypes}}})
	}
	if len(query.WheelDrive) > 0 {
		wheelDrives := strings.Split(query.WheelDrive, ",")
		match = append(match, bson.E{Key: "wheel_drive", Value: bson.D{{Key: "$in", Value: wheelDrives}}})
	}
	if len(query.SortBy) > 0 && query.SortOrder >= -1 {
		sort = bson.D{{Key: query.SortBy, Value: query.SortOrder}}
	}
	opts := options.Find().SetProjection(projection).SetSort(sort).SetSkip(skip).SetLimit(limit)
	results, err := automobileAdCollection.Find(ctx, match, opts)
	if err != nil {
		return nil, err
	}
	totalCount, err := automobileAdCollection.CountDocuments(ctx, match)
	if err != nil {
		return nil, err
	}
	for results.Next(ctx) {
		var automobileAd automobileAdSchema.AutomobileAdInListing
		if err = results.Decode(&automobileAd); err != nil {
			return nil, err
		}
		automobileAds = append(automobileAds, automobileAd)
	}
	listingResponse.TotalCount = uint64(totalCount)
	listingResponse.List = automobileAds
	return &listingResponse, nil
}

func UpdateAdActiveStatus(ctx context.Context, updateAutomobileAdActiveStatus automobileAdSchema.UpdateAutomobileAdActiveStatus, adId string) (string, error) {
	var automobileAdCollection *mongo.Collection = clients.GetMongoCollection(clients.GetConnectedMongoClient(), automobileAdSchema.AutomobileAdCollectionName)
	userClaims := ctx.Value(userService.UserClaims{})
	userId := userClaims.(*userService.UserClaims).Id.Hex()
	user, err := userService.GetUserById(ctx, userId)
	if err != nil || user == nil {
		return "", err
	}
	if user.Role != "admin" {
		return "", errors.New("Can not update ad")
	}

	update := bson.M{
		"active":     updateAutomobileAdActiveStatus.Active,
		"updated_at": time.Now(),
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

func UploadAutomobileAdMedia(ctx context.Context, files []*multipart.FileHeader, contentId string) (string, error) {
	var automobileAdCollection *mongo.Collection = clients.GetMongoCollection(clients.GetConnectedMongoClient(), automobileAdSchema.AutomobileAdCollectionName)
	userClaims := ctx.Value(userService.UserClaims{})
	userId := userClaims.(*userService.UserClaims).Id.Hex()
	user, err := userService.GetUserById(ctx, userId)
	if err != nil || user == nil {
		return "", err
	}
	content, err := GetAutomobileAdGeneralById(ctx, contentId)
	if err != nil || user == nil {
		return "", err
	}
	if len(content.Images) >= int(enum.MAX_FILE_PER_CONTENT) {
		return "", errors.New("Max File limit reached")
	}
	prefix := userId + "/" + contentId + "/"
	validFormats := []string{"image/png", "image/jpg", "image/jpeg"}
	metaData := map[string]*string{}
	for _, file := range files {
		contentType := file.Header.Get("Content-Type")
		if !helper.Contains(validFormats, contentType) {
			return "", errors.New("Format not supported")
		}
		url, uploadErr := mediaService.UploadFiles(file, "", "", prefix, metaData)
		if uploadErr != nil {
			return "", errors.New("Can not upload file")
		}
		update := bson.M{
			"$addToSet": bson.M{"images": url},
			"$set":      bson.M{"updated_at": time.Now()},
		}
		objId, err := primitive.ObjectIDFromHex(contentId)
		if err != nil {
			return "", err
		}
		saveAdResult, err := automobileAdCollection.UpdateOne(ctx, bson.M{"_id": objId}, update)
		if err != nil {
			return "", err
		}
		fmt.Println(saveAdResult.ModifiedCount)
	}
	return "Uploaded", nil
}

func DeleteAutomobileAdMedia(ctx context.Context, deleteAutomobileAdImage automobileAdSchema.DeleteAutomobileAdImage, contentId string) (string, error) {
	var automobileAdCollection *mongo.Collection = clients.GetMongoCollection(clients.GetConnectedMongoClient(), automobileAdSchema.AutomobileAdCollectionName)
	userClaims := ctx.Value(userService.UserClaims{})
	userId := userClaims.(*userService.UserClaims).Id.Hex()
	user, err := userService.GetUserById(ctx, userId)
	if err != nil || user == nil {
		return "", err
	}
	content, err := GetAutomobileAdGeneralById(ctx, contentId)
	if err != nil || user == nil {
		return "", err
	}
	if content == nil {
		return "", errors.New("No content found")
	}
	url := ""
	for _, v := range content.Images {
		if v == deleteAutomobileAdImage.Url {
			url = v
		}
	}
	split := strings.Split(deleteAutomobileAdImage.Url, "/")
	if url == "" {
		return "", errors.New("No image found")
	}
	joined := strings.Join(split[3:], "/")
	deleteErr := mediaService.DeleteFile(joined, "", "")
	if deleteErr != nil {
		return "", errors.New("Can not delete file")
	}
	update := bson.M{
		"$pull": bson.M{"images": url},
		"$set":  bson.M{"updated_at": time.Now()},
	}
	objId, err := primitive.ObjectIDFromHex(contentId)
	if err != nil {
		return "", err
	}
	saveAdResult, err := automobileAdCollection.UpdateOne(ctx, bson.M{"_id": objId}, update)
	if err != nil {
		return "", err
	}
	fmt.Println(saveAdResult.ModifiedCount)
	return "Removed", nil
}
