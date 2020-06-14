package configurations

import (
	"github.com/spf13/viper"
)

// GetTokenSecretKey :
func GetTokenSecretKey() string {

	return viper.GetString("JWT_SECRET_KEY")
}
