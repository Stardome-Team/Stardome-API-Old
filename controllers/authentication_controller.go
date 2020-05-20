package controllers

import (
	"fmt"
	"net/http"

	"github.com/Blac-Panda/Stardome-API/models"
	"github.com/Blac-Panda/Stardome-API/services"
	"github.com/gin-gonic/gin"
)

// AuthenticationController :
type AuthenticationController interface {
	AuthenticatePlayer(c *gin.Context)
}

// InitAuthenticationController :
func InitAuthenticationController(playerService services.PlayerService) AuthenticationController {
	return &handler{
		playerService: playerService,
	}
}

// AuthenticatePlayer :
func (h *handler) AuthenticatePlayer(c *gin.Context) {
	var authModel models.PlayerAuthentication

	if err := c.ShouldBindJSON(&authModel); err != nil {
		c.Error(err).SetType(gin.ErrorTypeBind)
		c.Status(http.StatusBadRequest)
		return
	}

	player, parsingErr := h.playerService.GetPlayerByUserName(authModel.UserName)

	if parsingErr != nil {
		c.Error(parsingErr.Error).SetType(parsingErr.Type).SetMeta(parsingErr.Metadata)
		c.Status(parsingErr.StatusCode)
		return
	}

	// match, err := utils.CompareHashWithPassword(authModel.Password, *player.PassHash)

	// if err != nil

	fmt.Print(player)
}
