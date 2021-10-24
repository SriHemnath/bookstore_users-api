package controllers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"

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

	restErr := errors.NewBadRequestError("invalid json body")
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
	userId, userErr := strconv.ParseInt(c.Param("user_id"), 10, 64)
	if userErr != nil {
		err := errors.NewBadRequestError("user id should be a number")
		c.JSON(err.Status, err)
	}
	user, getErr := services.GetUser(userId)
	if getErr != nil {
		c.JSON(getErr.Status, getErr)
		return
	}
	c.JSON(http.StatusOK, user)
}

func DeleteUser(c *gin.Context) {
	c.String(http.StatusNotImplemented, "Implement me!")
}
