package api_error

import "net/http"

type ApiError struct {
	StatusCode int    `json:"statusCode"`
	Error      string `json:"error"`
	Message    string `json:"message"`
}

func New(statusCode int, err error) ApiError {
	return ApiError{
		StatusCode: statusCode,
		Error:      http.StatusText(statusCode),
		Message:    err.Error(),
	}
}