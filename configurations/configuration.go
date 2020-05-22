package configurations

import (
	"github.com/spf13/viper"
)

var (
	// JWTTokenSecretKey :
	JWTTokenSecretKey string
)

func init() {

	viper.SetConfigFile(".env")

	err := viper.ReadInConfig()

	if err != nil {
		return
	}

	JWTTokenSecretKey = viper.GetString("JWT_SECRET_KEY")
}