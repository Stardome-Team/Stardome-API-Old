package router

import (
	errormiddleware "github.com/Blac-Panda/Stardome-API/middlewares/error"
	"github.com/gin-gonic/gin"

	authenticationcontroller "github.com/Blac-Panda/Stardome-API/controllers/authentication"
	playercontroller "github.com/Blac-Panda/Stardome-API/controllers/player"
	tournamentcontroller "github.com/Blac-Panda/Stardome-API/controllers/tournament"
)

// Routers This  function defines the available routes
// in Stardome API
func Routers() *gin.Engine {
	g := gin.Default()

	g.Use(errormiddleware.ErrorHandlerMiddleware())

	api := g.Group("/api")
	{
		api.POST("/auth/token", authenticationcontroller.AuthenticatePlayer)

		api.GET("/players", playercontroller.ListPlayers)
		api.POST("/players", playercontroller.CreatePlayer)

		api.GET("/players/:id", playercontroller.GetPlayer)
		api.PUT("/players/:id", playercontroller.UpdatePlayer)
		api.PATCH("/players/:id", playercontroller.ModifyPlayer)
		api.DELETE("/players/:id", playercontroller.DeletePlayer)

		api.GET("/tournaments", tournamentcontroller.ListTournaments)
		api.POST("/tournaments", tournamentcontroller.CreateTournament)

		api.GET("/tournaments/:id", tournamentcontroller.GetTournament)
		api.PUT("/tournaments/:id", tournamentcontroller.UpdateTournament)
		api.PATCH("/tournaments/:id", tournamentcontroller.ModifyTournament)
		api.DELETE("/tournaments/:id", tournamentcontroller.DeleteTournament)
	}

	return g
}
