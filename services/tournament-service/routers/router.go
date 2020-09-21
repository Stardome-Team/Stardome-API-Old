package router

import (
	"github.com/Stardome-Team/Stardome-API/services/tournament-service/controllers"
	"github.com/gin-gonic/gin"
)

// Routers : initialize routes for tournamenet service
func Routers() *gin.Engine {
	g := gin.Default()

	api := g.Group("/api")
	{

		api.GET("/tournaments", controllers.ListTournaments)
		api.POST("/tournaments", controllers.CreateTournament)

		api.GET("/tournaments/:id", controllers.GetTournament)
		api.PUT("/tournaments/:id", controllers.UpdateTournament)
		api.PATCH("/tournaments/:id", controllers.ModifyTournament)
		api.DELETE("/tournaments/:id", controllers.DeleteTournament)
	}
	return g
}
