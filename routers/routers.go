package router

import (
	controllers "github.com/Blac-Panda/Stardome-API/controllers/players"
	"github.com/gin-gonic/gin"
)

// Routers This  function defines the available routes
// in Stardome API
func Routers() *gin.Engine {
	g := gin.Default()

	api := g.Group("/api")
	{
		api.GET("/players", controllers.ListPlayers)
		api.POST("/players", controllers.CreatePlayer)

		api.GET("/players/:id", controllers.GetPlayer)
		api.PUT("/players/:id", controllers.UpdatePlayer)
		api.PATCH("/players/:id", controllers.ModifyPlayer)
		api.DELETE("/players/:id", controllers.DeletePlayer)
	}

	return g
}
