package services

import (
	"fmt"
	"github.com/egnimos/mvc/src/api/config"
	"github.com/egnimos/mvc/src/api/domain/github"
	"github.com/egnimos/mvc/src/api/domain/repositories"
	"github.com/egnimos/mvc/src/api/providers/github_provider"
	"github.com/egnimos/mvc/src/api/utils/errors"
	"net/http"
	"sync"

	//"strings"
)

type repoService struct {}

type repoServiceInterface interface {
	CreateRepo(input repositories.CreateRepoRequest) (*repositories.CreateRepoResponse, errors.ApiError)
	CreateRepos(input []repositories.CreateRepoRequest) (*repositories.CreateReposResponse, errors.ApiError)
}

var (
	RepositoryService repoServiceInterface
)

func init() {
	RepositoryService = &repoService{}
}

func (s *repoService) CreateRepo(input repositories.CreateRepoRequest) (*repositories.CreateRepoResponse, errors.ApiError) {
	//validate the repo name (Repo Name should NOT BE EMPTY)
	if err := input.Validate(); err != nil {
		return nil, err
	}

	//assigining the input value to the **CREATEREPOREQUEST** of github package
	request := github.CreateRepoRequest{
		Name: input.Name,
		Description: input.Description,
		Private: false,
	}

	//getting the response by the github_provider package
	response, err := github_provider.CreateRepo(config.GetGithubAccessToken(), request)
	if err != nil {
		return nil, errors.NewApiError(err.StatusCode, err.Message)
	}

	result := repositories.CreateRepoResponse{
		Id: response.Id,
		Owner: response.Owner.Login,
		Name: response.Name,
	}
	return &result, nil
}

func (s *repoService) CreateRepos(requests []repositories.CreateRepoRequest) (*repositories.CreateReposResponse, errors.ApiError) {
	reposResponseChannel := make(chan repositories.CreateReposResponse)
	reposResultChannel := make(chan repositories.CreateRepositoriesResult)
	//create a slice
	//var results []*repositories.CreateRepositoriesResult
	var wg sync.WaitGroup
	go s.handleRepoResult(&wg, reposResultChannel, reposResponseChannel)

	for _, request := range requests {
		wg.Add(1)
		go s.createRepoConcurrent(request, reposResultChannel)
	}

	wg.Wait()
	fmt.Println("waiting....")
	fmt.Println("Closed.....")
	close(reposResultChannel)
	response := <-reposResponseChannel
	fmt.Println("result channel.....")
	defer close(reposResponseChannel)

	successCreations := 0
	for _, responseResult := range response.Results {
		if responseResult.Response != nil {
			successCreations++
		}
	}

	//check the condition of the success creation
	if successCreations == 0 {
		response.StatusCode = response.Results[0].Error.Status()
	} else if successCreations == len(requests) {
		response.StatusCode = http.StatusCreated
	} else {
		response.StatusCode = http.StatusPartialContent
	}

	return &response, nil
}

func (s *repoService) handleRepoResult(wg *sync.WaitGroup, reposResultChannel chan repositories.CreateRepositoriesResult, reposResponseChannel chan repositories.CreateReposResponse) {
	var response repositories.CreateReposResponse

	for result := range reposResultChannel {
		repoResult := repositories.CreateRepositoriesResult{
			Response: result.Response,
			Error: result.Error,
		}
		response.Results = append(response.Results, repoResult)
		fmt.Println("result added to the response variable", response.Results)
		wg.Done()
	}
	reposResponseChannel <- response
}

func (s *repoService) createRepoConcurrent(repoRequest repositories.CreateRepoRequest, repoResultChannel chan repositories.CreateRepositoriesResult) {

	//send the request to the github api
	result, err := s.CreateRepo(repoRequest)
	if err != nil {
		fmt.Println("repo result error", err)
		repoResultChannel <- repositories.CreateRepositoriesResult{Error: err}
		return
	}

	fmt.Println("repo result without error", result)
	repoResultChannel <- repositories.CreateRepositoriesResult{Response: result}
}