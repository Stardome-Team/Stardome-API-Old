package playercontroller

import (
	"net/http"

	playermodel "github.com/Blac-Panda/Stardome-API/models/player"

	"github.com/gin-gonic/gin"
)

// ListPlayers :
func ListPlayers(c *gin.Context) {
	c.JSON(http.StatusOK, map[string]interface{}{
		"Hello": "World",
	})
}

// GetPlayer :
func GetPlayer(c *gin.Context) {
	c.JSON(http.StatusOK, map[string]interface{}{
		"Hello": "World",
	})
}

// CreatePlayer :
func CreatePlayer(c *gin.Context) {
	var regModel playermodel.PlayerRegistration

	if err := c.ShouldBindJSON(&regModel); err != nil {
		c.Error(err)
		c.Status(http.StatusBadRequest)
		return
	}

	c.JSON(http.StatusCreated, regModel)
}

// UpdatePlayer :
func UpdatePlayer(c *gin.Context) {
	c.JSON(http.StatusOK, map[string]interface{}{
		"Hello": "World",
	})
}

// ModifyPlayer :
func ModifyPlayer(c *gin.Context) {
	c.JSON(http.StatusOK, map[string]interface{}{
		"Hello": "World",
	})
}

// DeletePlayer :
func DeletePlayer(c *gin.Context) {
	c.JSON(http.StatusOK, map[string]interface{}{
		"Hello": "World",
	})
}
