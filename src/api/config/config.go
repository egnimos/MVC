package config

import "os"

const (
	apiGithubAccessToken = "SECRET_GITHUB_ACCESS_TOKEN"
)

//  1316715faa4111dec5b1509bf3d7b1940647a998

var (
	githubAccessToken = os.Getenv(apiGithubAccessToken)
)

func GetGithubAccessToken() string {
	return githubAccessToken
}
