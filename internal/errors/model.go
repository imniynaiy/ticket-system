package errors

// AppError represents authentication domain specific errors
type AppError struct {
	Code       string `json:"code"`
	Message    string `json:"message"`
	HTTPStatus int    `json:"-"` // HTTP status code, not exposed in JSON
	Err        error  `json:"-"` // Internal error details, not exposed in JSON
}
