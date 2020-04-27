package main

import (
	router "github.com/Blac-Panda/Stardome-API/routers"
)

func main() {

	app := router.Routers()

	app.Run(":1010")

}
