package main

import (
	"github.com/Blac-Panda/Stardome-API/routers"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

func main() {

	app := routers.Routers()

	app.Run(":1010")

}
