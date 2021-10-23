package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Health(c *gin.Context) {
	c.String(http.StatusOK, "good")
}

func Ping(c *gin.Context) {
	c.String(http.StatusOK, "pong")
}
