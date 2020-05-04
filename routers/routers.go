package router

import (
	"github.com/Blac-Panda/Stardome-API/controller"
	"github.com/Blac-Panda/Stardome-API/middlewares"
	"github.com/gin-gonic/gin"
)

// Routers This  function defines the available routes
// in Stardome API
func Routers() *gin.Engine {
	g := gin.New()

	g.Use(middlewares.ErrorHandlerMiddleware())

	api := g.Group("/api")
	{
		api.POST("/auth/token", controller.AuthenticatePlayer)

		api.GET("/controllers", controller.ListPlayers)
		api.POST("/controllers", controller.CreatePlayer)

		api.GET("/controllers/:id", controller.GetPlayer)
		api.PUT("/controllers/:id", controller.UpdatePlayer)
		api.PATCH("/controllers/:id", controller.ModifyPlayer)
		api.DELETE("/controllers/:id", controller.DeletePlayer)

		api.GET("/tournaments", controller.ListTournaments)
		api.POST("/tournaments", controller.CreateTournament)

		api.GET("/tournaments/:id", controller.GetTournament)
		api.PUT("/tournaments/:id", controller.UpdateTournament)
		api.PATCH("/tournaments/:id", controller.ModifyTournament)
		api.DELETE("/tournaments/:id", controller.DeleteTournament)
	}

	return g
}
