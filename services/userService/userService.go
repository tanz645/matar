package userService

import (
	"context"
	"errors"
	"fmt"
	"matar/clients"
	"matar/common/enum"
	"matar/configs"
	"matar/schemas/userSchema"
	"matar/utils"
	"time"

	"github.com/dgrijalva/jwt-go"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func GetUserByPhone(ctx context.Context, phone string) (*userSchema.User, error) {
	var user userSchema.User
	var userCollection *mongo.Collection = clients.GetMongoCollection(clients.GetConnectedMongoClient(), userSchema.UserCollectionName)
	err := userCollection.FindOne(ctx, bson.M{"phone": phone}).Decode(&user)
	if err != nil {
		fmt.Println(err)
		return nil, errors.New("Can not get user")
	}
	return &user, nil
}

func PushAdId(ctx context.Context, userId string, adId string) (*userSchema.User, error) {
	var user userSchema.User
	var userCollection *mongo.Collection = clients.GetMongoCollection(clients.GetConnectedMongoClient(), userSchema.UserCollectionName)
	objId, err := primitive.ObjectIDFromHex(userId)
	if err != nil {
		return nil, errors.New("Can not get user")
	}
	filter := bson.D{{Key: "_id", Value: objId}}
	update := bson.D{
		{
			Key:   "$addToSet",
			Value: bson.D{{Key: "ad_ids", Value: adId}},
		},
		{
			Key: "$set",
			Value: bson.D{
				{Key: "updated_at", Value: time.Now()},
				{Key: "updated_by", Value: enum.UPDATED_BY_SYSTEM},
			},
		},
	}
	err = userCollection.FindOneAndUpdate(
		ctx,
		filter,
		update,
	).Decode(&user)
	fmt.Println(&user)
	if err != nil {
		return nil, errors.New("Can not update user")
	}
	return &user, nil
}

func RemoveAdId(ctx context.Context, userId string, adId string) (*userSchema.User, error) {
	var user userSchema.User
	var userCollection *mongo.Collection = clients.GetMongoCollection(clients.GetConnectedMongoClient(), userSchema.UserCollectionName)
	objId, err := primitive.ObjectIDFromHex(userId)
	if err != nil {
		return nil, errors.New("Can not get user")
	}
	filter := bson.D{{Key: "_id", Value: objId}}
	update := bson.D{
		{
			Key:   "$pull",
			Value: bson.D{{Key: "ad_ids", Value: adId}},
		},
		{
			Key: "$set",
			Value: bson.D{
				{Key: "updated_at", Value: time.Now()},
				{Key: "updated_by", Value: enum.UPDATED_BY_SYSTEM},
			},
		},
	}
	err = userCollection.FindOneAndUpdate(
		ctx,
		filter,
		update,
	).Decode(&user)
	fmt.Println(&user)
	if err != nil {
		return nil, errors.New("Can not update user")
	}
	return &user, nil
}

func GetUserById(ctx context.Context, id string) (*userSchema.User, error) {
	var user userSchema.User
	var userCollection *mongo.Collection = clients.GetMongoCollection(clients.GetConnectedMongoClient(), userSchema.UserCollectionName)
	objId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, errors.New("Can not get user")
	}
	err = userCollection.FindOne(ctx, bson.M{"_id": objId}).Decode(&user)
	if err != nil {
		fmt.Println(err)
		return nil, errors.New("Can not get user")
	}
	return &user, nil
}

func CreateUser(ctx context.Context, user userSchema.User) (*mongo.InsertOneResult, error) {
	var userCollection *mongo.Collection = clients.GetMongoCollection(clients.GetConnectedMongoClient(), userSchema.UserCollectionName)
	userByPhone, _ := GetUserByPhone(ctx, user.Phone)
	if userByPhone != nil && userByPhone.PhoneNumberVerified == true {
		return nil, errors.New("Phone number already verified, please login using it")
	}
	hashed, err := utils.HashPassword(user.Password)
	if err != nil {
		return nil, errors.New("Error in password hashing")
	}
	var maxAd int16 = 1
	if user.Type == USER_TYPE_COMPANY {
		maxAd = 30
	}
	if user.Type == USER_TYPE_INDIVIDUAL {
		maxAd = 3
	}
	newUser := userSchema.User{
		Id:                  primitive.NewObjectID(),
		Phone:               user.Phone,
		Password:            hashed,
		Type:                user.Type,
		Country:             user.Country,
		AdIds:               []string{},
		MaxAd:               maxAd,
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

func LoginUser(ctx context.Context, userLogin userSchema.UserLogin) (*string, error) {
	var jwtKey = []byte(configs.Common.Service.Secret)
	userByPhone, _ := GetUserByPhone(ctx, userLogin.Phone)
	if userByPhone == nil {
		return nil, errors.New("Username or password not matched")
	}
	verified := utils.CheckPasswordHash(userLogin.Password, userByPhone.Password)
	if !verified {
		return nil, errors.New("Username or password not matched")
	}

	expirationTime := time.Now().Add(170 * time.Hour)
	claims := &JwtClaims{
		Id: userByPhone.Id,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		return nil, errors.New("Can not login")
	}
	return &tokenString, nil
}

func VerifyToken(token string) (*UserClaims, error) {
	fmt.Print(token[0])
	var jwtKey = []byte(configs.Common.Service.Secret)
	claims := &JwtClaims{}
	tkn, err := jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})

	userClaims := &UserClaims{
		Id: claims.Id,
	}
	if err != nil {
		return nil, err
	}
	if !tkn.Valid {
		return nil, nil
	}
	return userClaims, nil
}
