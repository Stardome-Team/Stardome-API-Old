package routers

import (
	"github.com/Blac-Panda/Stardome-API/player-service/configurations"
	"github.com/Blac-Panda/Stardome-API/player-service/controllers"
	"github.com/Blac-Panda/Stardome-API/player-service/middlewares"
	"github.com/Blac-Panda/Stardome-API/player-service/repositories"
	"github.com/Blac-Panda/Stardome-API/player-service/services"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

var (
	authenticationHandler controllers.AuthenticationController
	playerHandler         controllers.PlayerController
)

func init() {

	playerRepository := repositories.NewPlayerRepository(func() *gorm.DB {
		db, err := configurations.GetDB()

		if err != nil {
			return nil
		}

		return db
	})

	playerService := services.NewPlayerService(playerRepository)
	authenticationService := services.NewAuthenticationService(playerRepository)

	playerHandler = controllers.InitPlayerController(playerService)
	authenticationHandler = controllers.InitAuthenticationController(authenticationService)
}

// Routers This  function defines the available routes
// in Stardome API
func Routers() *gin.Engine {
	g := gin.Default()

	g.Use(middlewares.ErrorHandlerMiddleware())

	api := g.Group("/api")
	{
		api.POST("/auth/token", authenticationHandler.AuthenticatePlayer)
		api.POST("/players", playerHandler.CreatePlayer)

		auth := api.Group("", middlewares.AuthHandlerMiddleware())
		{
			auth.GET("/players", playerHandler.ListPlayers)

			auth.GET("/players/:id", playerHandler.GetPlayer)
			auth.PUT("/players/:id", playerHandler.UpdatePlayer)
			auth.PATCH("/players/:id", playerHandler.ModifyPlayer)
			auth.DELETE("/players/:id", playerHandler.DeletePlayer)
		}
	}

	return g
}
