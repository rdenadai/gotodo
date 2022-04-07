package database

import (
	"github.com/rdenadai/gotodo/app/models"
	_ "gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var (
	DB  *gorm.DB
	err error
)

func StartDatabaseConnection() {
	// dsn := "host=localhost user=rdenadai password=rdenadai dbname=gotodo port=5432 sslmode=disable"
	// DB, err = gorm.Open(postgres.Open(dsn))
	DB, err = gorm.Open(sqlite.Open("gotodo.db"))
	if err != nil {
		panic(err.Error())
	}
	DB.AutoMigrate(&models.Todo{})
}
