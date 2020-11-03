package restclient

import (
	"bytes"
	"encoding/json"
	"net/http"
)

func Post(url string, body interface{}, headers http.Header) (*http.Response, error) {
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