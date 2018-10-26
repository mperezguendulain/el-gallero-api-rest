package errors

// APIError represents an error that can be sent in an error response.
type APIError struct {

	// ErrorCode is the code uniquely identifying an error
	ErrorCode int `json:"error_code"`

	// Message is the error message that may be displayed to end users
	Message string `json:"message"`
}
