package database

import (
	"go-rest-gin/models"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	DB  *gorm.DB
	err error
)

func Connection() {
	conString := "host=localhost user=root password=root dbname=root port=5432 sslmode=disable"

	DB, err = gorm.Open(postgres.Open(conString))

	if err != nil {
		log.Panic("Erro to Connect")
	}

	DB.AutoMigrate(&models.Aluno{})
}
