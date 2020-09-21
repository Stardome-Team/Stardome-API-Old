package main

import (
	"flag"

	"github.com/spf13/viper"

	"github.com/Stardome-Team/Stardome-API/services/player-service/routers"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

func init() {

	isProduction := flag.Bool("prod", false, "Flag sets environment variable to production")
	flag.Parse()

	if *isProduction {
		viper.SetConfigFile("./.prod.env")
		viper.ReadInConfig()
	} else {
		viper.SetConfigFile("./.dev.env")
		viper.ReadInConfig()
	}
}

func main() {

	app := routers.Routers()

	app.Run(":1010")

}
