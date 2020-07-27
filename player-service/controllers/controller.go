package controllers

import (
	"github.com/Blac-Panda/Stardome-API/player-service/services"
)

// handler :
type handler struct {
	playerService services.PlayerService
	authService   services.AuthenticationService
}
