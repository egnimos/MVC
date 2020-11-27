package repositories

import (
	"net/http"

	"github.com/egnimos/mvc/src/api/domain/repositories"
	"github.com/egnimos/mvc/src/api/services"
	"github.com/egnimos/mvc/src/api/utils/errors"
	"github.com/gin-gonic/gin"
)

func CreateRepo(ctx *gin.Context) {
	//
	var request repositories.CreateRepoRequest
	if err := ctx.ShouldBindJSON(&request); err != nil {
		apiErr := errors.NewBadRequestError("Invalid JSON body")
		ctx.JSON(apiErr.Status(), apiErr)
		return
	}

	//getting the response from the github
	result, err := services.RepositoryService.CreateRepo(request)
	if err != nil {
		// apiErr := errors.NewBadRequestError(err.Message())
		ctx.JSON(err.Status(), err)
		return
	}
	ctx.JSON(http.StatusCreated, result)
}

//create multiple repos
func CreateRepos(ctx *gin.Context) {
	//
	var request []repositories.CreateRepoRequest
	if err := ctx.ShouldBindJSON(&request); err != nil {
		apiErr := errors.NewBadRequestError("Invalid JSON body")
		ctx.JSON(apiErr.Status(), apiErr)
		return
	}

	//getting the response from the github
	result, err := services.RepositoryService.CreateRepos(request)
	if err != nil {
		// apiErr := errors.NewBadRequestError(err.Message())
		ctx.JSON(err.Status(), err)
		return
	}
	ctx.JSON(result.StatusCode, result)
}