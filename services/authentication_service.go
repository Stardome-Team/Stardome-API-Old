package services

// AuthenticationService :
type AuthenticationService interface {
}

// NewAuthenticationService :
func NewAuthenticationService() AuthenticationService {
	return &service{}
}
