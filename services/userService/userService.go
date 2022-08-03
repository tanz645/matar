package userService

import (
	"context"
	"errors"
	"fmt"
	"matar/clients"
	"matar/configs"
	"matar/models"
	"matar/utils"
	"time"

	"github.com/dgrijalva/jwt-go"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type Claims struct {
	Id primitive.ObjectID `json:"id"`
	jwt.StandardClaims
}

func GetUserByPhone(ctx context.Context, phone string) (*models.User, error) {
	var user models.User
	var userCollection *mongo.Collection = clients.GetMongoCollection(clients.GetConnectedMongoClient(), models.UserCollectionName)
	err := userCollection.FindOne(ctx, bson.M{"phone": phone}).Decode(&user)
	if err != nil {
		fmt.Println(err)
		return nil, errors.New("Can not get user")
	}
	return &user, nil
}

func CreateUser(ctx context.Context, user models.User) (*mongo.InsertOneResult, error) {
	var userCollection *mongo.Collection = clients.GetMongoCollection(clients.GetConnectedMongoClient(), models.UserCollectionName)
	userByPhone, _ := GetUserByPhone(ctx, user.Phone)
	if userByPhone != nil && userByPhone.PhoneNumberVerified == true {
		return nil, errors.New("Phone number already verified, please login using it")
	}
	hashed, err := utils.HashPassword(user.Password)
	if err != nil {
		return nil, errors.New("Error in password hashing")
	}
	newUser := models.User{
		Id:                  primitive.NewObjectID(),
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

func LoginUser(ctx context.Context, userLogin models.UserLogin) (*string, error) {
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
	claims := &Claims{
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
