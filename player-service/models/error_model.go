package models

import "github.com/gin-gonic/gin"

// Error :
type Error struct {
	Error *ErrorObject `json:"error,omitempty"`
}

// ErrorObject :
type ErrorObject struct {
	Code    int            `json:"code"`
	Message string         `json:"message"`
	Errors  []ErrorsObject `json:"errors,omitempty"`
}

// ErrorsObject :
type ErrorsObject struct {
	Domain       string `json:"domain,omitempty"`
	Reason       string `json:"reason,omitempty"`
	Message      string `json:"message,omitempty"`
	ExtendedHelp string `json:"extendedHelp,omitempty"`
	SendReport   string `json:"sendReport,omitempty"`
}

// ErrorParsing :
type ErrorParsing struct {
	Type       gin.ErrorType
	Metadata   interface{}
	Error      error
	StatusCode int
}
