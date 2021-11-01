package main

import (
	"github.com/SriHemnath/bookstore_users-api/app"
	"github.com/SriHemnath/bookstore_users-api/utils/logger"
)

func main() {
	logger.Info("Starting Appication")
	app.StartApplication()
}
