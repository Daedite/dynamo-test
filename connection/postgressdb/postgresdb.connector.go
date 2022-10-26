package postgressdb

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Connect() (gormDB *gorm.DB, err error) {
	return gorm.Open(postgres.Open(getConnectionString()))
}

func getConnectionString() string {
	return fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=%s TimeZone=%s", "localhost", "root", "root", "jwo_connector", 5432, "disable", "Europe/London")
}
