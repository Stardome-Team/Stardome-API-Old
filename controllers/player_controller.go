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

	qInt, err := utils.ParseQueryToInt(c.Query("index"), c.Query("size"))

	if err != nil {
		c.Error(err.Error).SetType(err.Type).SetMeta(err.Metadata)
		c.Status(err.StatusCode)
		return
	}

	players, err := h.service.ListPlayers(qInt[0], qInt[1])

	if err != nil {
		c.Error(err.Error).SetType(err.Type).SetMeta(err.Metadata)
		c.Status(err.StatusCode)
		return
	}

	c.JSON(http.StatusOK, models.Result{
		Data: models.Data{
			Data: &models.DataObject{
				StartIndex:       *players.StartIndex,
				ItemsPerPage:     *players.ItemsPerPage,
				TotalItems:       *players.TotalItems,
				CurrentItemCount: players.CurrentItemCount,
				Items:            players.Items,
			},
		},
	})
}

// GetPlayer :
func (h *handler) GetPlayer(c *gin.Context) {
	id := c.Param("id")

	player, err := h.service.GetPlayer(id)

	if err != nil {
		c.Error(err.Error).SetType(err.Type).SetMeta(err.Metadata)
		c.Status(err.StatusCode)
		return
	}

	c.JSON(http.StatusOK, models.Result{
		Data: models.Data{
			Data: &models.DataObject{
				Player: player,
			},
		},
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
		c.Error(utils.ErrorPasswordMismatch).SetType(gin.ErrorTypePublic).SetMeta(utils.ReasonPasswordMismatch)
		c.Status(http.StatusBadRequest)
		return
	}

	player, err := h.service.CreatePlayer(&regModel)

	if err != nil {
		c.Error(err.Error).SetType(err.Type).SetMeta(err.Metadata)
		c.Status(err.StatusCode)
		return
	}

	c.JSON(http.StatusCreated, models.Result{
		Data: models.Data{
			Data: &models.DataObject{
				Player: player,
			},
		},
	})
}

// UpdatePlayer :
func (h *handler) UpdatePlayer(c *gin.Context) {
	id := c.Param("id")

	var playerModel models.Player

	if err := c.ShouldBindJSON(&playerModel); err != nil {
		c.Error(err).SetType(gin.ErrorTypeBind)
		c.Status(http.StatusBadRequest)
		return
	}

	player, err := h.service.UpdatePlayer(id, &playerModel)

	if err != nil {
		c.Error(err.Error).SetType(err.Type).SetMeta(err.Metadata)
		c.Status(err.StatusCode)
		return
	}

	c.JSON(http.StatusOK, models.Result{
		Data: models.Data{
			Data: &models.DataObject{
				Player: player,
			},
		},
	})
}

// ModifyPlayer :
func (h *handler) ModifyPlayer(c *gin.Context) {
	id := c.Param("id")

	var playerModel map[string]interface{} = make(map[string]interface{})

	if err := c.ShouldBindJSON(&playerModel); err != nil {
		c.Error(err).SetType(gin.ErrorTypeBind)
		c.Status(http.StatusBadRequest)
		return
	}

	player, err := h.service.ModifyPlayer(id, playerModel)

	if err != nil {
		c.Error(err.Error).SetType(err.Type).SetMeta(err.Metadata)
		c.Status(err.StatusCode)
		return
	}

	c.JSON(http.StatusOK, models.Result{
		Data: models.Data{
			Data: &models.DataObject{
				Player: player,
			},
		},
	})
}

// DeletePlayer :
func (h *handler) DeletePlayer(c *gin.Context) {
	c.JSON(http.StatusOK, map[string]interface{}{
		"Hello": "World",
	})
}
