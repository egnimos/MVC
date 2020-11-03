package github

//[{"key":"Authorization","value":"token fea1bd2cafef906c63fd6a5429141bfbe3d84442","description":"","type":"text","enabled":true}]

type GitErrorResponse struct {
	StatusCode int `json:"status_code"`
	Message string `json:"message"`
	DocumentationUrl string `json:"documentation_url"`
}