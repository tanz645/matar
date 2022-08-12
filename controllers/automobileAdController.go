package controllers

import (
	"context"
	"fmt"
	"matar/common/responses"
	"matar/models/automobileAdModel"
	"matar/services/automobileAdService"
	"matar/services/userService"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func CreateAutomobileAd() gin.HandlerFunc {

	return func(c *gin.Context) {
		ctx := context.Background()
		ctx = context.WithValue(ctx, userService.UserClaims{}, c.Value("user"))
		ctx, cancel := context.WithTimeout(ctx, 10*time.Second)

		var automobileAd automobileAdModel.AutomobileAd
		defer cancel()
		if err := c.BindJSON(&automobileAd); err != nil {
			c.JSON(http.StatusBadRequest, responses.FailedResponse{Status: http.StatusBadRequest, Error: true, Message: "Ad can not be created", Data: err.Error()})
			return
		}
		if validationErr := validate.Struct(&automobileAd); validationErr != nil {
			c.JSON(http.StatusUnprocessableEntity, responses.FailedResponse{Status: http.StatusUnprocessableEntity, Error: true, Message: "Ad can not be created", Data: validationErr.Error()})
			return
		}
		result, err := automobileAdService.CreateAutomobileAd(ctx, automobileAd)
		if err != nil || result == "" {
			fmt.Println(err)
			c.JSON(http.StatusInternalServerError, responses.FailedResponse{Status: http.StatusInternalServerError, Error: true, Message: "Ad can not be created", Data: nil})
			return
		}
		data := map[string]string{
			"id": result,
		}
		c.JSON(http.StatusCreated, responses.SuccessResponse{Status: http.StatusCreated, Success: true, Message: "Ad created", Data: data})
	}
}

func GetAutomobileAdById() gin.HandlerFunc {

	return func(c *gin.Context) {
		ctx := context.Background()
		ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
		id := c.Param("id")

		defer cancel()
		result, err := automobileAdService.GeAutomobileAdGeneralById(ctx, id)
		if err != nil {
			fmt.Println(err)
			c.JSON(http.StatusNotFound, responses.FailedResponse{Status: http.StatusNotFound, Error: true, Message: "Can not get Ad", Data: nil})
			return
		}
		data := result
		c.JSON(http.StatusOK, responses.SuccessResponse{Status: http.StatusOK, Success: true, Message: "", Data: data})
	}
}

func UpdateAutomobileAdById() gin.HandlerFunc {

	return func(c *gin.Context) {
		ctx := context.Background()
		ctx = context.WithValue(ctx, userService.UserClaims{}, c.Value("user"))
		ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
		id := c.Param("id")

		var automobileAd automobileAdModel.AutomobileAd
		defer cancel()
		if err := c.BindJSON(&automobileAd); err != nil {
			c.JSON(http.StatusBadRequest, responses.FailedResponse{Status: http.StatusBadRequest, Error: true, Message: "Ad can not be updated", Data: err.Error()})
			return
		}
		if validationErr := validate.Struct(&automobileAd); validationErr != nil {
			c.JSON(http.StatusUnprocessableEntity, responses.FailedResponse{Status: http.StatusUnprocessableEntity, Error: true, Message: "Ad can not be updated", Data: validationErr.Error()})
			return
		}
		result, err := automobileAdService.UpdateAutomobileAdById(ctx, automobileAd, id)
		if err != nil || result == "" {
			fmt.Println(err)
			c.JSON(http.StatusInternalServerError, responses.FailedResponse{Status: http.StatusInternalServerError, Error: true, Message: "Ad can not be updated", Data: nil})
			return
		}
		data := map[string]string{
			"id": result,
		}
		c.JSON(http.StatusOK, responses.SuccessResponse{Status: http.StatusOK, Success: true, Message: "Ad updated", Data: data})
	}
}

func DeleteAutomobileAdById() gin.HandlerFunc {

	return func(c *gin.Context) {
		ctx := context.Background()
		ctx = context.WithValue(ctx, userService.UserClaims{}, c.Value("user"))
		ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
		id := c.Param("id")

		defer cancel()
		err := automobileAdService.RemoveAutomobileAdGenera(ctx, id)
		if err != nil {
			fmt.Println(err)
			c.JSON(http.StatusInternalServerError, responses.FailedResponse{Status: http.StatusInternalServerError, Error: true, Message: "Can not remove Ad", Data: nil})
			return
		}

		c.JSON(http.StatusOK, responses.SuccessResponse{Status: http.StatusOK, Success: true, Message: "Ad removed", Data: nil})
	}
}
