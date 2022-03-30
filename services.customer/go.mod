module wptest/services.customer

go 1.16

replace wptest/shared => ../shared

require (
	github.com/gin-gonic/gin v1.6.3
	github.com/google/uuid v1.2.0
	github.com/pact-foundation/pact-go v1.5.1
	github.com/satori/go.uuid v1.2.0
	wptest/shared v0.0.0-00010101000000-000000000000
	gorm.io/driver/postgres v1.0.8
	gorm.io/gorm v1.21.3
)
