package errors

import "net/http"

//create a ERROR interface
type ApiError interface {
	Status() int
	Message() string
	Error() string
}

//apiError struct
type apiError struct {
	AStatus int `json:"status"`
	AMessage string `json:"message"`
	AnError string `json:"error,omitempty"`
}

//status
func (s *apiError) Status() int {
	return s.AStatus
}

//message
func (s *apiError) Message() string {
	return s.AMessage
}

//error
func (s *apiError) Error() string {
	return s.AnError
}

//New API Error
func NewApiError(statusCode int, message string) ApiError {
	return &apiError{AStatus: statusCode, AMessage: message}
}

//New internal server error
func NewInternalServerError(message string) ApiError {
	return &apiError{AStatus: http.StatusInternalServerError, AMessage: message}
}

//New Bad request error
func NewBadRequestError(message string) ApiError {
	return &apiError{AStatus: http.StatusBadRequest, AMessage: message}
}

//New Not Found Error
func NewNotFoundError(message string) ApiError {
	return &apiError{AStatus: http.StatusNotFound, AMessage: message}
}