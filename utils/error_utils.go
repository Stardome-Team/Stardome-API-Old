package utils

import "errors"

var (
	// ErrorInternalError :
	ErrorInternalError = errors.New("whoops something went wrong")
	// ErrorPasswordMismatch :
	ErrorPasswordMismatch = errors.New("Password don't match")
	// ErrorNotYetImplemented :
	ErrorNotYetImplemented = errors.New("Not yet implemented")
	// ErrorPlayerAlreadyExist :
	ErrorPlayerAlreadyExist = errors.New("Player already exist")
	// ErrorPlayerCreationFailed :
	ErrorPlayerCreationFailed = errors.New("Unable to create new player")
	// ErrorDatabaseOperationFailed :
	ErrorDatabaseOperationFailed = errors.New("Database operation failed")
)

var (
	// ReasonInternalServer :
	ReasonInternalServer = "InternalServerError"
	// ReasonPasswordMismatch :
	ReasonPasswordMismatch = "PasswordMismatch"
	// ReasonEntityCreationFailed :
	ReasonEntityCreationFailed = "EntityCreationFailed"
	// ReasonFieldValidationError :
	ReasonFieldValidationError = "FieldValidationError"
)
