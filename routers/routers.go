package router

import (
	"github.com/Blac-Panda/Stardome-API/controllers"
	"github.com/Blac-Panda/Stardome-API/middlewares"
	"github.com/Blac-Panda/Stardome-API/repositories/database"
	"github.com/Blac-Panda/Stardome-API/services"
	"github.com/gin-gonic/gin"
)

var (
	playerHandler controllers.PlayerController
)

func init() {
	playerRepository := database.NewPlayerRepository(nil)
	playerService := services.NewPlayerService(playerRepository)
	playerHandler = controllers.InitPlayerController(playerService)
}

// Routers This  function defines the available routes
// in Stardome API
func Routers() *gin.Engine {
	g := gin.New()

	g.Use(middlewares.ErrorHandlerMiddleware())

	api := g.Group("/api")
	{
		api.POST("/auth/token", controllers.AuthenticatePlayer)

		api.GET("/players", playerHandler.ListPlayers)
		api.POST("/players", playerHandler.CreatePlayer)

		api.GET("/players/:id", playerHandler.GetPlayer)
		api.PUT("/players/:id", playerHandler.UpdatePlayer)
		api.PATCH("/players/:id", playerHandler.ModifyPlayer)
		api.DELETE("/players/:id", playerHandler.DeletePlayer)

		api.GET("/tournaments", controllers.ListTournaments)
		api.POST("/tournaments", controllers.CreateTournament)

		api.GET("/tournaments/:id", controllers.GetTournament)
		api.PUT("/tournaments/:id", controllers.UpdateTournament)
		api.PATCH("/tournaments/:id", controllers.ModifyTournament)
		api.DELETE("/tournaments/:id", controllers.DeleteTournament)
	}

	return g
}
