package restclient

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
)

var(
	enableMock = false
	mocks = make(map[string]*Mock)
)

type Mock struct {
	Url string
	HttpMethod string
	Response *http.Response
	Err error
}

func AddMockUp(mock Mock) {
	mocks[mock.Url] = &mock
}

func FlushMockUp() {
	mocks = make(map[string]*Mock)
}

//start mocking server for testing
func StartMocking() {
	enableMock = true
}

//stop mocking server
func StopMocking() {
	enableMock = false
}

func Post(url string, body interface{}, headers http.Header) (*http.Response, error) {
	//when the mocking is true
	if enableMock {
		mock := mocks[url]
		if mock == nil {
			return nil, errors.New("no mockup found for give request")
		}
		return mock.Response, mock.Err
	}

	//when the mocking is not there
	jsonBytes, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}

	//Generating a Request
	request, err := http.NewRequest(http.MethodPost, url, bytes.NewReader(jsonBytes))
	request.Header = headers

	//send the request to particular url GITHUB API
	client := http.Client{}
	return client.Do(request)
}