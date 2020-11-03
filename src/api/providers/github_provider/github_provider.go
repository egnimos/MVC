package github_provider

import (
	"encoding/json"
	"fmt"
	"github.com/egnimos/mvc/src/api/clients/restclient"
	"github.com/egnimos/mvc/src/api/domain/github"
	"io/ioutil"
	"log"
	"net/http"
)

const (
	headerAuthorization = "Authorization"
	headerAuthorizationFormat = "token %s"

	//url
	urlCreateRepo = "https://api.github.com/user/repos"
)

func getAuthorizationHeader(accessToken string) string {
	return fmt.Sprintf(headerAuthorizationFormat, accessToken)
}

func CreateRepo(accessToken string, request github.CreateRepoRequest) (*github.CreateRepoResponse, *github.GitErrorResponse) {
	headers := http.Header{}
	headers.Set(headerAuthorization, getAuthorizationHeader(accessToken))

	//sending the parameters to restClient api and get the response from the method
	response, err := restclient.Post(urlCreateRepo, request, headers)
	if err != nil {
		log.Fatalln("error when trying to create a new repo in github")
		return nil, &github.GitErrorResponse{StatusCode: http.StatusBadRequest, Message: err.Error()}
	}

	//read the response
	bytes, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, &github.GitErrorResponse{StatusCode: http.StatusBadRequest, Message: "invalid response body"}
	}
	//close the body
	defer response.Body.Close()

	//if the github sends the error in response
	if response.StatusCode > 299 {
		var errResponse github.GitErrorResponse
		//covert JSON into struct
		if err := json.Unmarshal(bytes, errResponse); err != nil {
			return nil, &github.GitErrorResponse{StatusCode: http.StatusInternalServerError, Message: "invalid json response body"}
		}
		errResponse.StatusCode = response.StatusCode
		return nil, &errResponse
	}

	//if there is no error
	var result github.CreateRepoResponse
	//convert the JSON into STRUCT
	if err := json.Unmarshal(bytes, result); err != nil {
		return nil, &github.GitErrorResponse{StatusCode: http.StatusInternalServerError, Message: "error when trying to unmarshal the github response"}
	}
	return &result, nil
}