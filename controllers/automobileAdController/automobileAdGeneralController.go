package automobileAdController

import (
	"context"
	"fmt"
	"matar/common/responses"
	"matar/controllers"
	"matar/schemas/automobileAdSchema"
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

		var automobileAd automobileAdSchema.AutomobileAd
		defer cancel()
		if err := c.BindJSON(&automobileAd); err != nil {
			c.JSON(http.StatusBadRequest, responses.FailedResponse{Status: http.StatusBadRequest, Error: true, Message: "Ad can not be created", Data: err.Error()})
			return
		}
		if validationErr := controllers.Validate.Struct(&automobileAd); validationErr != nil {
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
		result, err := automobileAdService.GetAutomobileAdGeneralById(ctx, id)
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

		var automobileAd automobileAdSchema.AutomobileAd
		defer cancel()
		if err := c.BindJSON(&automobileAd); err != nil {
			c.JSON(http.StatusBadRequest, responses.FailedResponse{Status: http.StatusBadRequest, Error: true, Message: "Ad can not be updated", Data: err.Error()})
			return
		}
		if validationErr := controllers.Validate.Struct(&automobileAd); validationErr != nil {
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

func SearchAutomobileAd() gin.HandlerFunc {

	return func(c *gin.Context) {
		ctx := context.Background()
		ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
		defer cancel()
		var searchAutomobileAdGeneral automobileAdSchema.SearchAutomobileAdGeneral
		if err := c.ShouldBindQuery(&searchAutomobileAdGeneral); err != nil {
			c.JSON(http.StatusBadRequest, responses.FailedResponse{Status: http.StatusBadRequest, Error: true, Message: "Ad can not be created", Data: err.Error()})
			return
		}
		if validationErr := controllers.Validate.Struct(&searchAutomobileAdGeneral); validationErr != nil {
			c.JSON(http.StatusUnprocessableEntity, responses.FailedResponse{Status: http.StatusUnprocessableEntity, Error: true, Message: "Ad can not be created", Data: validationErr.Error()})
			return
		}
		result, err := automobileAdService.SearchAutomobileAdGeneral(ctx, searchAutomobileAdGeneral)
		if err != nil {
			fmt.Println(err)
			c.JSON(http.StatusInternalServerError, responses.FailedResponse{Status: http.StatusInternalServerError, Error: true, Message: "Can not get Ads", Data: nil})
			return
		}
		c.JSON(http.StatusOK, responses.SuccessResponse{Status: http.StatusOK, Success: true, Message: "", Data: result})
	}
}

func GetAutomobileAdByUserId() gin.HandlerFunc {

	return func(c *gin.Context) {
		ctx := context.Background()
		ctx = context.WithValue(ctx, userService.UserClaims{}, c.Value("user"))
		ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
		id := c.Param("id")

		defer cancel()
		result, err := automobileAdService.GetAutomobileAdGeneralByUserId(ctx, id)
		if err != nil || result == nil {
			fmt.Println(err)
			c.JSON(http.StatusNotFound, responses.FailedResponse{Status: http.StatusNotFound, Error: true, Message: "Can not get Ad", Data: nil})
			return
		}
		data := result
		c.JSON(http.StatusOK, responses.SuccessResponse{Status: http.StatusOK, Success: true, Message: "", Data: data})
	}
}

func GetAutomobileAdsByUserId() gin.HandlerFunc {

	return func(c *gin.Context) {
		ctx := context.Background()
		ctx = context.WithValue(ctx, userService.UserClaims{}, c.Value("user"))
		ctx, cancel := context.WithTimeout(ctx, 10*time.Second)

		defer cancel()
		result, err := automobileAdService.GetAutomobileAdsGeneralByUserId(ctx)
		if err != nil || result == nil {
			fmt.Println(err)
			c.JSON(http.StatusNotFound, responses.FailedResponse{Status: http.StatusNotFound, Error: true, Message: "Can not get Ad", Data: nil})
			return
		}
		data := result
		c.JSON(http.StatusOK, responses.SuccessResponse{Status: http.StatusOK, Success: true, Message: "", Data: data})
	}
}

func UploadImages() gin.HandlerFunc {

	return func(c *gin.Context) {
		ctx := context.Background()
		ctx = context.WithValue(ctx, userService.UserClaims{}, c.Value("user"))
		ctx, cancel := context.WithTimeout(ctx, 100*time.Second)
		id := c.Param("id")
		form, _ := c.MultipartForm()
		files := form.File["file"]
		defer cancel()
		if len(files) == 0 || files == nil {
			c.JSON(http.StatusUnprocessableEntity, responses.FailedResponse{Status: http.StatusUnprocessableEntity, Error: true, Message: "No Image provided", Data: nil})
			return
		}
		if len(files) > 1 {
			c.JSON(http.StatusUnprocessableEntity, responses.FailedResponse{Status: http.StatusUnprocessableEntity, Error: true, Message: "Max 1 Image can be uploaded per request", Data: nil})
			return
		}
		result, err := automobileAdService.UploadAutomobileAdMedia(ctx, files, id)
		if err != nil || result == "" {
			fmt.Println(err)
			c.JSON(http.StatusInternalServerError, responses.FailedResponse{Status: http.StatusInternalServerError, Error: true, Message: "Image can not be uploaded", Data: nil})
			return
		}
		c.JSON(http.StatusOK, responses.SuccessResponse{Status: http.StatusOK, Success: true, Message: "Image uploaded", Data: nil})
	}
}

func DeleteImage() gin.HandlerFunc {

	return func(c *gin.Context) {
		ctx := context.Background()
		ctx = context.WithValue(ctx, userService.UserClaims{}, c.Value("user"))
		ctx, cancel := context.WithTimeout(ctx, 100*time.Second)
		id := c.Param("id")
		defer cancel()

		var deleteAutomobileAdImage automobileAdSchema.DeleteAutomobileAdImage
		defer cancel()
		if err := c.BindJSON(&deleteAutomobileAdImage); err != nil {
			c.JSON(http.StatusBadRequest, responses.FailedResponse{Status: http.StatusBadRequest, Error: true, Message: "Image can not be deleted", Data: err.Error()})
			return
		}
		if validationErr := controllers.Validate.Struct(&deleteAutomobileAdImage); validationErr != nil {
			c.JSON(http.StatusUnprocessableEntity, responses.FailedResponse{Status: http.StatusUnprocessableEntity, Error: true, Message: "Image can not be deleted", Data: validationErr.Error()})
			return
		}
		result, err := automobileAdService.DeleteAutomobileAdMedia(ctx, deleteAutomobileAdImage, id)
		if err != nil || result == "" {
			fmt.Println(err)
			c.JSON(http.StatusInternalServerError, responses.FailedResponse{Status: http.StatusInternalServerError, Error: true, Message: "Image can not be deleted", Data: nil})
			return
		}
		c.JSON(http.StatusOK, responses.SuccessResponse{Status: http.StatusOK, Success: true, Message: "Image deleted", Data: nil})
	}
}
