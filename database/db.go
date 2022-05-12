package database

import (
	"log"

	"github.com/douglira/alura-golang-gin-api-rest/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	DB  *gorm.DB
	err error
)

func ConnectDatabase() {
	connection := "host=localhost user=douglas password=docker dbname=alura_go_rest_gin port=5432 sslmode=disable"
	DB, err = gorm.Open(postgres.Open(connection))
	if nil != err {
		log.Panic("Error at database connection:", err)
	}
	DB.AutoMigrate(&models.Student{})
}
