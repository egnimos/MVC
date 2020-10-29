package app

import "github.com/egnimos/mvc/controllers"

func mapUrls() {
	router.GET("/users/:user_id", controllers.GetUsers)
}