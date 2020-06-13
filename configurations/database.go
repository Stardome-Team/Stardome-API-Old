package configurations

import (
	"fmt"

	"github.com/jinzhu/gorm"
	"github.com/spf13/viper"

	_ "github.com/jinzhu/gorm/dialects/postgres" // import for psq;
)

var (
	connection string
	host       string
	port       string
	database   string
	username   string
	password   string
)

func init() {
	viper.SetConfigFile("../.env")
}

// GetDB :
func GetDB() (*gorm.DB, error) {
	err := viper.ReadInConfig()

	if err != nil {
		// TODO: log error
		return nil, err
	}
	connection = viper.GetString("DB_CONNECTION")
	host = viper.GetString("DB_HOST")
	port = viper.GetString("DB_PORT")
	database = viper.GetString("DB_DATABASE")
	username = viper.GetString("DB_USERNAME")
	password = viper.GetString("DB_PASSWORD")

	var dbInfo string = fmt.Sprintf("host=%s port=%s dbname=%s user=%s password=%s sslmode=disable", host, port, database, username, password)

	db, err := gorm.Open(connection, dbInfo)

	if err != nil {
		// TODO: log error
		return nil, err
	}

	return db, nil
}
