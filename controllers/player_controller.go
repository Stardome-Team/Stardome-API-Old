package controllers

import (
	"net/http"

	"github.com/Blac-Panda/Stardome-API/models"
	"github.com/Blac-Panda/Stardome-API/services"
	"github.com/Blac-Panda/Stardome-API/utils"

	"github.com/gin-gonic/gin"
)

// PlayerController :
type PlayerController interface {
	ListPlayers(c *gin.Context)
	GetPlayer(c *gin.Context)
	CreatePlayer(c *gin.Context)
	UpdatePlayer(c *gin.Context)
	ModifyPlayer(c *gin.Context)
	DeletePlayer(c *gin.Context)
}

// handler :
type handler struct {
	service services.PlayerService
}

// InitPlayerController :
func InitPlayerController(s services.PlayerService) PlayerController {
	return &handler{
		service: s,
	}
}

// ListPlayers :
func (h *handler) ListPlayers(c *gin.Context) {
	c.JSON(http.StatusOK, map[string]interface{}{
		"Hello": "World",
	})
}

// GetPlayer :
func (h *handler) GetPlayer(c *gin.Context) {
	c.JSON(http.StatusOK, map[string]interface{}{
		"Hello": "World",
	})
}

// CreatePlayer :
func (h *handler) CreatePlayer(c *gin.Context) {
	var regModel models.PlayerRegistration

	if err := c.ShouldBindJSON(&regModel); err != nil {
		c.Error(err).SetType(gin.ErrorTypeBind)
		c.Status(http.StatusBadRequest)
		return
	}

	if regModel.Password != regModel.ConfirmPassword {
		c.Error(utils.ErrorPasswordMismatch).SetType(gin.ErrorTypePublic).SetMeta(utils.ErrorPasswordMismatch)
		c.Status(http.StatusBadRequest)
		return
	}

	player, err := h.service.CreatePlayer(&regModel)

	if err != nil {
		c.Error(err).SetType(gin.ErrorTypePublic).SetMeta(utils.ReasonEntityCreationFailed)
		c.Status(http.StatusOK)
		return
	}

	c.JSON(http.StatusCreated, player)
}

// UpdatePlayer :
func (h *handler) UpdatePlayer(c *gin.Context) {
	c.JSON(http.StatusOK, map[string]interface{}{
		"Hello": "World",
	})
}

// ModifyPlayer :
func (h *handler) ModifyPlayer(c *gin.Context) {
	c.JSON(http.StatusOK, map[string]interface{}{
		"Hello": "World",
	})
}

// DeletePlayer :
func (h *handler) DeletePlayer(c *gin.Context) {
	c.JSON(http.StatusOK, map[string]interface{}{
		"Hello": "World",
	})
}
