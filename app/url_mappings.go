package app

import "github.com/SriHemnath/bookstore_users-api/controllers"

//contains all mappings of the microservice
func mapUrls() {
	router.GET("/ping", controllers.Ping)
	router.GET("/health", controllers.Health)
	router.POST("/user", controllers.CreateUser)
	router.GET("/user/:user_id", controllers.GetUser)
	router.PUT("/user/update", controllers.UpdateUser)   //all fields should be present
	router.PATCH("/user/update", controllers.UpdateUser) //will update incoming field
	router.DELETE("/user/delete", controllers.DeleteUser)
	router.GET("/user/search", controllers.Search)
}
