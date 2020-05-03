package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// ListTournaments :
func ListTournaments(c *gin.Context) {
	c.JSON(http.StatusOK, map[string]interface{}{
		"Hello": "World",
	})
}

// GetTournament :
func GetTournament(c *gin.Context) {
	c.JSON(http.StatusOK, map[string]interface{}{
		"Hello": "World",
	})
}

// CreateTournament :
func CreateTournament(c *gin.Context) {
	c.JSON(http.StatusOK, map[string]interface{}{
		"Hello": "World",
	})
}

// UpdateTournament :
func UpdateTournament(c *gin.Context) {
	c.JSON(http.StatusOK, map[string]interface{}{
		"Hello": "World",
	})
}

// ModifyTournament :
func ModifyTournament(c *gin.Context) {
	c.JSON(http.StatusOK, map[string]interface{}{
		"Hello": "World",
	})
}

// DeleteTournament :
func DeleteTournament(c *gin.Context) {
	c.JSON(http.StatusOK, map[string]interface{}{
		"Hello": "World",
	})
}
