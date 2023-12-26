package database

import (
	"github.com/pequenojoohn/go-gin/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	DB  *gorm.DB
	err error
)

func ConnectedDataBase() {
	dsn := "host=localhost user=root password=root dbname=root port=5432 sslmode=disable"
	conn, err := gorm.Open(postgres.Open(dsn))

	if err != nil {
		panic(err.Error())
	}

	conn.AutoMigrate(&models.Aluno{})

	DB = conn
}
