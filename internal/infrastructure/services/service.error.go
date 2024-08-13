package services

import (
	"fmt"
	"net/http"
)

type ServiceError struct {
	Message   string
	ErrorCode int
}

func (e *ServiceError) Error() string {
	return e.Message
}

func (e *ServiceError) HttpStatus() int {
	errorCodeMap := map[int]int{
		Conflict:            http.StatusConflict,
		DocumentNotFound:    http.StatusNotFound,
		EndpointNotFound:    http.StatusNotFound,
		InternalServerError: http.StatusInternalServerError,
		NotModified:         http.StatusNotModified,
	}

	httpCode, ok := errorCodeMap[e.ErrorCode]
	if !ok {
		return http.StatusInternalServerError
	}

	return httpCode
}

func (e *ServiceError) StringErrorCode() string {
	return fmt.Sprintf("%X", e.ErrorCode)
}

func NewServiceError(errorCode int, format string, v ...any) *ServiceError {
	return &ServiceError{
		Message:   fmt.Sprintf(format, v...),
		ErrorCode: errorCode,
	}
}

const (
	DocumentNotFound    = 0x190001
	EndpointNotFound    = 0x190002
	InternalServerError = 0x1F4001
	RepositoryError     = 0x1F4002
	Conflict            = 0x1F4009
	NotModified         = 0x1F3004
)
