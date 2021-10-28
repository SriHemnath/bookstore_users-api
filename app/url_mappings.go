package app

import "github.com/SriHemnath/bookstore_users-api/controllers"

func mapUrls() {
	router.GET("/ping", controllers.Ping)
	router.GET("/health", controllers.Health)
	router.POST("/user", controllers.CreateUser)
	router.GET("/user/:user_id", controllers.GetUser)
}
