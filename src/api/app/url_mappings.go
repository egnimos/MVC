package app

import (
	"github.com/egnimos/mvc/src/api/controller/polo"
	"github.com/egnimos/mvc/src/api/controller/repositories"
)

func mapUrls() {
	router.GET("/lan", polo.IndexFunction)
	router.POST("/createRepo", repositories.CreateRepo)
	router.POST("/createRepos", repositories.CreateRepos)
}
