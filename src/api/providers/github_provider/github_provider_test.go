package github_provider

import (
	"errors"
	"github.com/egnimos/mvc/src/api/clients/restclient"
	"github.com/egnimos/mvc/src/api/domain/github"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
	"testing"
)

//run the test main function
func TestMain(m *testing.M) {
	restclient.StartMocking()
	os.Exit(m.Run())
}

//run test constant for authorization
func TestConstants(t *testing.T) {
	assert.EqualValues(t, "Authorization", headerAuthorization)
	assert.EqualValues(t, "token %s", headerAuthorizationFormat)
	assert.EqualValues(t, "https://api.github.com/user/repos", urlCreateRepo)
}

//testing the get the Authorization function
func TestGetAuthorizationHeader(t *testing.T) {
	//execution
	header := getAuthorizationHeader("123")
	assert.EqualValues(t, "token 123", header)
}

//testing the CREATE REPO function
func TestRestClientWhenError(t *testing.T) {
	//we dont want to connect the external api while mocking and testing the restClient
	//initialization
	restclient.FlushMockUp()
	restclient.AddMockUp(restclient.Mock{
		Url: "https://api.github.com/user/repos",
		HttpMethod: http.MethodPost,
		Err: errors.New("invalid restClient response"),
	})
	//execution
	response, err := CreateRepo("", github.CreateRepoRequest{})
	assert.Nil(t, response)
	assert.NotNil(t, err)
	assert.EqualValues(t, http.StatusBadRequest, err.StatusCode)
	assert.EqualValues(t, "invalid restClient response", err.Message)
}

//testing the read response
func TestCreateRepoInvalidResponse(t *testing.T) {
	restclient.FlushMockUp()
	//invalid closer
	invalidCloser, _ := os.Open("-asf3")
	restclient.AddMockUp(restclient.Mock{
		Url: "https://api.github.com/user/repos",
		HttpMethod: http.MethodPost,
		Response: &http.Response{
			StatusCode: http.StatusCreated,
			Body: invalidCloser,
		},
	})

	//sending the request to Create Repo
	response, err := CreateRepo("", github.CreateRepoRequest{})
	assert.Nil(t, response)
	assert.NotNil(t, err)
	assert.EqualValues(t, http.StatusBadRequest, err.StatusCode)
	assert.EqualValues(t, "invalid response body", err.Message)
}

//testing when the error occurs in github api
func TestGithubApiErrorUnmarshaling(t *testing.T) {
	restclient.FlushMockUp()

	restclient.AddMockUp(restclient.Mock{
		Url: "https://api.github.com/user/repos",
		HttpMethod: http.MethodPost,
		Response: &http.Response{
			StatusCode: http.StatusInternalServerError,
			Body: ioutil.NopCloser(strings.NewReader(`{"message": 1}`)),
		},
	})

	response, err := CreateRepo("", github.CreateRepoRequest{})
	assert.Nil(t, response)
	assert.NotNil(t, err)
	assert.EqualValues(t, http.StatusBadRequest, err.StatusCode)
	assert.EqualValues(t, "invalid json response body", err.Message)
}

//testing the JSON github response
func TestGithubErrorResponse(t *testing.T) {
	restclient.FlushMockUp()

	restclient.AddMockUp(restclient.Mock{
		Url: "https://api.github.com/user/repos",
		HttpMethod: http.MethodPost,
		Response: &http.Response{
			StatusCode: http.StatusUnauthorized,
			Body: ioutil.NopCloser(
				strings.NewReader(`{"message": "Requires authentication","documentation_url": "https://developer.github.com/v3/repos/#create"}`)),
		},
	})

	response, err := CreateRepo("", github.CreateRepoRequest{})
	assert.Nil(t, response)
	assert.NotNil(t, err)
	assert.EqualValues(t, http.StatusUnauthorized, err.StatusCode)
	assert.EqualValues(t, "Requires authentication", err.Message)
}

//testing the error while inmarshaling the result
func TestCreateRepoResultUnmarshalingError(t *testing.T) {
	restclient.FlushMockUp()

	restclient.AddMockUp(restclient.Mock{
		Url: "https://api.github.com/user/repos",
		HttpMethod: http.MethodPost,
		Response: &http.Response{
			StatusCode: http.StatusCreated,
			Body: ioutil.NopCloser(strings.NewReader(`{"id": "123"}`)),
		},
	})

	response, err := CreateRepo("", github.CreateRepoRequest{})
	assert.Nil(t, response)
	assert.NotNil(t, err)
	assert.EqualValues(t, http.StatusInternalServerError, err.StatusCode)
	assert.EqualValues(t, "error when trying to unmarshal the github response", err.Message)
}


//testing the final result
func TestCreateRepoNoerror(t *testing.T) {
	restclient.FlushMockUp()

	restclient.AddMockUp(restclient.Mock{
		Url: "https://api.github.com/user/repos",
		HttpMethod: http.MethodPost,
		Response: &http.Response{
			StatusCode: http.StatusCreated,
			Body: ioutil.NopCloser(strings.NewReader(`{"id": 123, "name": "golang-test-repo"}`)),
		},
	})

	response, err := CreateRepo("", github.CreateRepoRequest{})
	assert.NotNil(t, response)
	assert.Nil(t, err)
	assert.EqualValues(t, 123, response.Id)
	assert.EqualValues(t, "golang-test-repo", response.Name)
}