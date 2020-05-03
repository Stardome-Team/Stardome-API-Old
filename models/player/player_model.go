package playermodel

import (
	errormodel "github.com/Blac-Panda/Stardome-API/models/error"
)

// Player :
type Player struct {
	ID             int    `json:"id"`
	PlayerID       string `json:"playerId"`
	UserName       string `json:"userName"`
	EmailAddress   string `json:"emailAddress"`
	DisplayName    string `json:"displayName"`
	AvatarURL      string `json:"avatarUrl"`
	AvatarBlurHash string `json:"avatarBlurHash"`
}

// PlayerAuthentication :
type PlayerAuthentication struct {
	UserName string `json:"userName"`
	Password string `json:"password"`
}

// PlayerRegistration :
type PlayerRegistration struct {
	UserName        string `json:"userName,omitempty" binding:"required,min=3,max=25"`
	Password        string `json:"password" binding:"required"`
	ConfirmPassword string `json:"confirmPassword" binding:"required"`
}

// GetErrors :
func (m *PlayerRegistration) GetErrors() []errormodel.FieldError {
	var errors []errormodel.FieldError

	if m.UserName == "" {

	}

	return errors
}
