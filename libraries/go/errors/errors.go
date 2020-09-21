package errors

// Error represents the error object
type Error struct {
	Error *ErrorResponse `json:"error,omitempty"`
}

// ErrorResponse represents the response of the error
type ErrorResponse struct {
	StatusCode int           `json:"statusCode"`
	Message    string        `json:"message"`
	Errors     []ErrorObject `json:"errors,omitempty"`
}

// ErrorObject gives extensive info to the error and other errors involved
type ErrorObject struct {
	Domain       string `json:"domain,omitempty"`
	Reason       string `json:"reason,omitempty"`
	Message      string `json:"message,omitempty"`
	ExtendedHelp string `json:"extendedHelp,omitempty"`
	SendReport   string `json:"sendReport,omitempty"`
}

// InternalError a response for any internal error within the system
var InternalError ErrorObject = ErrorObject{
	Reason:  "InternalServerError",
	Message: "whoops something went wrong",
}

// ValidationError a response for validation errors
var ValidationError ErrorObject = ErrorObject{
	Reason:  "FieldValidationError",
	Message: "Your request is in a bad format",
}
