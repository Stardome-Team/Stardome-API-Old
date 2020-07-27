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
	// ErrorPlayerNotFound :
	ErrorPlayerNotFound = errors.New("Player not found")
	// ErrorEncryptionFailed :
	ErrorEncryptionFailed = errors.New("Encryption failed")
	// ErrorInvalidQuery :
	ErrorInvalidQuery = errors.New("Invalid query parameter")
	// ErrorRequestIDMismatch :
	ErrorRequestIDMismatch = errors.New("Request ID do not match")
	// ErrorPlayerUpdateFailed :
	ErrorPlayerUpdateFailed = errors.New("Unable to update player")
	// ErrorPlayerDeleteFailed :
	ErrorPlayerDeleteFailed = errors.New("Unable to delete player")
	// ErrorAuthenticationFailed :
	ErrorAuthenticationFailed = errors.New("Username or Password not valid")
	// ErrorAuthorizationTokenNotFound :
	ErrorAuthorizationTokenNotFound = errors.New("Player authorization not found")
	// ErrorAuthorizationVerificationFailed :
	ErrorAuthorizationVerificationFailed = errors.New("Authorization failed")
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
	// ReasonEntityNotFound :
	ReasonEntityNotFound = "EntityNotFound"
	// ReasonIDMismatch :
	ReasonIDMismatch = "IDMismatch"
	// ReasonEntityDeletionFailed :
	ReasonEntityDeletionFailed = "EntityDeletionFailed"
	// ReasonAuthorizationFailed :
	ReasonAuthorizationFailed = "AuthorizationFailed"
)
