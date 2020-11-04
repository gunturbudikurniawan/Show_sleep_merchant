package controllers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gunturbudikurniawan/Show_sleep_merchant/api/auth"
	"github.com/gunturbudikurniawan/Show_sleep_merchant/api/models"
	"github.com/gunturbudikurniawan/Show_sleep_merchant/api/security"
	"github.com/gunturbudikurniawan/Show_sleep_merchant/api/utils/formaterror"
	"golang.org/x/crypto/bcrypt"
)

func (server *Server) LoginMerchant(c *gin.Context) {

	//clear previous error if any
	errList = map[string]string{}

	body, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"status":      http.StatusUnprocessableEntity,
			"first error": "Unable to get request",
		})
		return
	}
	merchant := models.Subscribers1{}
	err = json.Unmarshal(body, &merchant)
	log.Println(err)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"status": http.StatusUnprocessableEntity,
			"error":  "Cannot unmarshal body",
		})
		return
	}
	merchant.Prepare()
	errorMessages := merchant.Validate("login")
	if len(errorMessages) > 0 {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"status": http.StatusUnprocessableEntity,
			"error":  errorMessages,
		})
		return
	}
	merchantData, err := server.SignInMerchant(merchant.Email, merchant.SecretPassword)
	if err != nil {
		formattedError := formaterror.FormatError(err.Error())
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"status": http.StatusUnprocessableEntity,
			"error":  formattedError,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status":   http.StatusOK,
		"response": merchantData,
	})
}

func (server *Server) SignInMerchant(email, password string) (map[string]interface{}, error) {

	var err error

	merchantData := make(map[string]interface{})

	merchant := models.Subscribers1{}

	err = server.DB.Debug().Model(models.Subscribers1{}).Where("email = ?", email).Take(&merchant).Error
	if err != nil {
		fmt.Println("this is the error getting the user: ", err)
		return nil, err
	}
	err = security.VerifyPassword(merchant.SecretPassword, password)
	if err != nil && err == bcrypt.ErrMismatchedHashAndPassword {
		fmt.Println("this is the error hashing the password: ", err)
		return nil, err
	}
	token, err := auth.CreateToken(merchant.ID)
	if err != nil {
		fmt.Println("this is the error creating the token: ", err)
		return nil, err
	}
	merchantData["token"] = token
	merchantData["id"] = merchant.ID
	merchantData["email"] = merchant.Email

	return merchantData, nil
}
