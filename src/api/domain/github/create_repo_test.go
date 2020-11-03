package github

import (
	"encoding/json"
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCreateRepoRequest(t *testing.T) {
	request := CreateRepoRequest{
		Name: "Go Tutorial",
		Description: "a golang tutorial how to create a repo",
		Homepage: "https://github.com",
		Private: true,
		HasIssues: true,
		HasProjects: true,
		HasWiki: true,
	}

	// Marshal takes an input interface in form of struct and convert it into JSON
	bytes, err := json.Marshal(request)

	assert.Nil(t, err)
	assert.NotNil(t, bytes)

	//print the result
	s := string(bytes)
	fmt.Println(s)

	//unmarshal from JSON to STRUCT
	var target CreateRepoRequest
	err = json.Unmarshal(bytes, &target)

	assert.Nil(t, err)
	assert.EqualValues(t, target.Name, request.Name)
	assert.EqualValues(t, target.HasIssues, request.HasIssues)
	fmt.Println(target)
}
