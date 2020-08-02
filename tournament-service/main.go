package main

import router "github.com/Blac-Panda/Stardome-API/player-service/routers"

func main() {
	app := router.Routers()

	app.Run(":1011")
}
