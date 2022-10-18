package userService

import (
	"github.com/dgrijalva/jwt-go"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type JwtClaims struct {
	Id primitive.ObjectID `json:"id"`
	jwt.StandardClaims
}

type UserClaims struct {
	Id primitive.ObjectID `json:"id"`
}

const (
	USER_TYPE_INDIVIDUAL string = "individual"
	USER_TYPE_COMPANY           = "company"
)
