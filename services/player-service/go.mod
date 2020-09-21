module github.com/Blac-Panda/Stardome-API/services/player-service

go 1.14

replace github.com/Blac-Panda/Stardome-API/libraries/go/jwt => ../../libraries/go/jwt

require (
	github.com/Blac-Panda/Stardome-API/libraries/go/jwt v0.0.0
	github.com/DATA-DOG/go-sqlmock v1.4.1
	github.com/alexedwards/argon2id v0.0.0-20200802152012-2464efd3196b
	github.com/fatih/structs v1.1.0
	github.com/gin-gonic/gin v1.6.3
	github.com/go-playground/validator/v10 v10.3.0
	github.com/jinzhu/gorm v1.9.15
	github.com/rs/xid v1.2.1
	github.com/spf13/viper v1.7.1
	github.com/stoewer/go-strcase v1.2.0
)
