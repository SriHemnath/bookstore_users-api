package app

import (
	"github.com/gin-gonic/gin"
)

var (
	router = gin.Default() //private variable only available inside app package
)

func StartApplication() {
	mapUrls()
	router.Run(":8080") //blocking call
}
