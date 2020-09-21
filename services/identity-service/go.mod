module github.com/Blac-Panda/Stardome-API/services/identity-service

go 1.14

replace github.com/Blac-Panda/Stardome-API/libraries/go/errors => ../../libraries/go/errors

replace github.com/Blac-Panda/Stardome-API/libraries/go/responses => ../../libraries/go/responses

require (
	github.com/Blac-Panda/Stardome-API/libraries/go/errors v0.0.0
	github.com/Blac-Panda/Stardome-API/libraries/go/responses v0.0.0
	github.com/gin-gonic/gin v1.6.3
	github.com/go-playground/validator/v10 v10.3.0
	github.com/google/uuid v1.1.2
	github.com/jackc/pgproto3/v2 v2.0.4 // indirect
	github.com/kelseyhightower/envconfig v1.4.0
	github.com/lib/pq v1.8.0
	github.com/stoewer/go-strcase v1.2.0
	go.uber.org/zap v1.16.0
	golang.org/x/crypto v0.0.0-20200820211705-5c72a883971a // indirect
	golang.org/x/sys v0.0.0-20200905004654-be1d3432aa8f // indirect
	golang.org/x/tools v0.0.0-20200904185747-39188db58858 // indirect
	gopkg.in/yaml.v2 v2.3.0
	gorm.io/driver/postgres v1.0.0
	gorm.io/gorm v1.20.0
)
