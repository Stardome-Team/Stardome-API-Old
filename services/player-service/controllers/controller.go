package controllers

import (
	"github.com/Stardome-Team/Stardome-API/services/player-service/services"
)

// handler :
type handler struct {
	playerService services.PlayerService
	authService   services.AuthenticationService
}
