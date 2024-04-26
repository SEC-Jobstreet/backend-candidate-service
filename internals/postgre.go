package internals

import (
	"fmt"
	"github.com/SEC-Jobstreet/backend-candidate-service/utils"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var _db *gorm.DB

func InitGorm(config utils.Config) {
	db, err := gorm.Open(postgres.Open(config.DBSource), &gorm.Config{QueryFields: true})
	if err != nil {
		panic("failed to connect to database")
	}

	_db = db
	fmt.Println("Connected successfully!")
}

// GetDb get database connection
func GetDb() *gorm.DB {
	return _db
}
