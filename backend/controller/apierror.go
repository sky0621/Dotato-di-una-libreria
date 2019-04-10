package controller

import "net/http"

type apiError struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}

func errorJSON(statusCode int, message string) apiError {
	return apiError{
		Status:  http.StatusText(statusCode),
		Message: message,
	}
}
