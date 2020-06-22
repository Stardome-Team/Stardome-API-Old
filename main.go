package main

import (
	"flag"
	"fmt"

	"github.com/spf13/viper"

	"github.com/Blac-Panda/Stardome-API/routers"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

func init() {

	isProduction := flag.Bool("prod", false, "Flag sets environment variable to production")
	flag.Parse()

	if *isProduction {
		viper.SetConfigFile(".dev.env")
		viper.ReadInConfig()
	} else {
		viper.SetConfigFile(".prod.env")
		viper.ReadInConfig()
	}

	fmt.Printf("\n\n\n Secret -> %v \v\v\v", viper.GetString("JWT_SECRET_KEY"))

}

func main() {

	app := routers.Routers()

	app.Run(":1010")

}
