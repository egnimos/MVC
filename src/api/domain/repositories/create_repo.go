package repositories

import (
	"github.com/egnimos/mvc/src/api/utils/errors"
	"strings"
)

type CreateRepoRequest struct {
	Name string `json:"name"`
	Description string `json:"description"`
}

//validate the name space in the create REPO request NAME
func (r *CreateRepoRequest) Validate() errors.ApiError{
	r.Name = strings.TrimSpace(r.Name)
	if r.Name == "" {
		return errors.NewBadRequestError("Invalid repository name")
	}
	return nil
}

type CreateRepoResponse struct {
	Id int `json:"id"`
	Owner string `json:"owner"`
	Name string `json:"name"`
}

type CreateReposResponse struct {
	StatusCode int `json:"status"`
	Results []CreateRepositoriesResult `json:"results"`
}

type CreateRepositoriesResult struct {
	Response *CreateRepoResponse `json:"repo"`
	Error errors.ApiError `json:"error"`
}