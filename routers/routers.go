package router

import (
	"github.com/gin-gonic/gin"

	"github.com/Blac-Panda/Stardome-API/controllers/players"
	"github.com/Blac-Panda/Stardome-API/controllers/tournaments"
)

// Routers This  function defines the available routes
// in Stardome API
func Routers() *gin.Engine {
	g := gin.Default()

	api := g.Group("/api")
	{
		api.GET("/players", players.ListPlayers)
		api.POST("/players", players.CreatePlayer)

		api.GET("/players/:id", players.GetPlayer)
		api.PUT("/players/:id", players.UpdatePlayer)
		api.PATCH("/players/:id", players.ModifyPlayer)
		api.DELETE("/players/:id", players.DeletePlayer)

		api.GET("/tournaments", tournaments.ListTournaments)
		api.POST("/tournaments", tournaments.CreateTournament)

		api.GET("/tournaments/:id", tournaments.GetTournament)
		api.PUT("/tournaments/:id", tournaments.UpdateTournament)
		api.PATCH("/tournaments/:id", tournaments.ModifyTournament)
		api.DELETE("/tournaments/:id", tournaments.DeleteTournament)
	}

	return g
}
