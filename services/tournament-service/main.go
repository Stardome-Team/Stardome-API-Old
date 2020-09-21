package main

import router "github.com/Stardome-Team/Stardome-API/services/tournament-service/routers"

func main() {
	app := router.Routers()

	app.Run(":1011")
}
