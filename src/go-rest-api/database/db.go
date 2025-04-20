package database

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	DB  *gorm.DB
	err error
)

func ConectaDB() {

	connectionString := "host=localhost user=leonardo password=leonardo dbname=leonardo port=5432 sslmode=disable"
	DB, err = gorm.Open(postgres.Open(connectionString))

	if err != nil {
		log.Panic("Erro de conexão com DB")
	}

}
