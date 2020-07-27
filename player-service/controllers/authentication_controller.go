package controllers

import (
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
func InitAuthenticationController(as services.AuthenticationService) AuthenticationController {
	return &handler{
		authService: as,
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

	token, err := h.authService.AuthenticatePlayer(c, &authModel)

	if err != nil {
		c.Error(err.Error).SetType(err.Type).SetMeta(err.Metadata)
		c.Status(err.StatusCode)
		return
	}

	c.JSON(
		http.StatusCreated,
		token,
	)
}
