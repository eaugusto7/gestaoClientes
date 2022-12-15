package db

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	Database *gorm.DB
	err      error
)

func ConectaBanco() {
	stringDeConexao := "host=localhost user=user password=password dbname=database port=5432 sslmode=disable"
	Database, err = gorm.Open(postgres.Open(stringDeConexao))
	if err != nil {
		log.Panic("Erro ao conectar com banco de dados")
	}
}
