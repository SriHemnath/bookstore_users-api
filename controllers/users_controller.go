package controllers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/SriHemnath/bookstore_users-api/domain/users"
	"github.com/SriHemnath/bookstore_users-api/services"
	"github.com/SriHemnath/bookstore_users-api/utils/errors"
	"github.com/gin-gonic/gin"
)

func CreateUser(c *gin.Context) {
	var user users.User
	byte, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		fmt.Println(err)
		return
	}

	restErr := errors.RestError{
		Message: "invalid json body",
		Status:  http.StatusBadRequest,
		Error:   "bad_request",
	}
	if err = json.Unmarshal(byte, &user); err != nil {
		fmt.Println("Error during unmarshelling ", err)
		c.JSON(restErr.Status, restErr)
		return
	}
	//shouldBindJson will do the same unmarshelling and binding as above

	result, saveErr := services.CreateUser(user)
	if saveErr != nil {
		c.JSON(saveErr.Status, saveErr)
		return
	}
	c.JSON(http.StatusCreated, result)
}

func FindUser(c *gin.Context) {
	c.String(http.StatusNotImplemented, "Implement me!")
}

func GetUser(c *gin.Context) {
	c.String(http.StatusNotImplemented, "Implement me!")
}

func DeleteUser(c *gin.Context) {
	c.String(http.StatusNotImplemented, "Implement me!")
}
