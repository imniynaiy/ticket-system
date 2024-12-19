package errors

import "net/http"

var (
	ErrInvalidRequest = &AppError{
		Code:       "ERR001",
		Message:    "Invalid request format",
		HTTPStatus: http.StatusBadRequest,
	}
	ErrInternalServerError = &AppError{
		Code:       "ERR999",
		Message:    "Internal server error",
		HTTPStatus: http.StatusInternalServerError,
	}
	ErrNotFound = &AppError{
		Code:       "ERR404",
		Message:    "Not found",
		HTTPStatus: http.StatusNotFound,
	}
)
