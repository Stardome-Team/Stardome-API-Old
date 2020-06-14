package main

import (
	"github.com/spf13/viper"

	"github.com/Blac-Panda/Stardome-API/routers"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

func main() {

	viper.SetConfigFile(".env")
	viper.ReadInConfig()

	app := routers.Routers()

	app.Run(":1010")

}
