package main

import (
	"flag"

	"github.com/spf13/viper"

	"github.com/Blac-Panda/Stardome-API/player-service/routers"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

func init() {

	isProduction := flag.Bool("prod", false, "Flag sets environment variable to production")
	flag.Parse()

	if *isProduction {
<<<<<<< HEAD:main.go
		viper.SetConfigFile(".prod.env")
		viper.ReadInConfig()
	} else {
		viper.SetConfigFile(".dev.env")
=======
		viper.SetConfigFile("./.prod.env")
		viper.ReadInConfig()
	} else {
		viper.SetConfigFile("./.dev.env")
>>>>>>> micro-split:player-service/main.go
		viper.ReadInConfig()
	}
}

func main() {

	app := routers.Routers()

	app.Run(":1010")

}
