package errors

import "net/http"

func (e *AppError) Error() string {
	return e.Message
}

// Authentication error codes
var (
	// Login errors (AUTH0xx)
	ErrInvalidCredentials = &AppError{
		Code:       "AUTH001",
		Message:    "Invalid email or password",
		HTTPStatus: http.StatusBadRequest,
	}

	// Registration errors (AUTH1xx)
	ErrEmailAlreadyExists = &AppError{
		Code:       "AUTH101",
		Message:    "Email is already registered",
		HTTPStatus: http.StatusBadRequest,
	}
	ErrInvalidEmail = &AppError{
		Code:       "AUTH102",
		Message:    "Invalid email format",
		HTTPStatus: http.StatusBadRequest,
	}
	ErrWeakPassword = &AppError{
		Code:       "AUTH103",
		Message:    "Password does not meet security requirements",
		HTTPStatus: http.StatusBadRequest,
	}
	ErrInvalidUsername = &AppError{
		Code:       "AUTH104",
		Message:    "Username must be between 3 and 30 characters",
		HTTPStatus: http.StatusBadRequest,
	}

	// Token errors (AUTH2xx)
	ErrInvalidToken = &AppError{
		Code:       "AUTH201",
		Message:    "Invalid or expired token",
		HTTPStatus: http.StatusBadRequest,
	}
	ErrForbidden = &AppError{
		Code:       "AUTH202",
		Message:    "Forbidden",
		HTTPStatus: http.StatusForbidden,
	}
)
