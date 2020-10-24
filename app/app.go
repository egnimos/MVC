package app

import (
	"fmt"
	"github.com/egnimos/mvc/controllers"
	"net/http"
)

func StartApp()  {
	http.HandleFunc("/users", controllers.GetUsers)

	fmt.Println("Server is Started...")
	http.ListenAndServe(":8080", nil)
}