package models

// RequestError describes error that occurred in application
type RequestError struct {
	// ErrorCode is the code of your error; not exported to json. Added only for demonstration purposes.
	ErrorCode int `json:"-"`
	// Message is the description of error; exported to json
	Message string `json:"message"`
}
