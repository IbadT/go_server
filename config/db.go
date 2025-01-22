package config

import (
	"log"

	"github.com/IbadT/go_server/model"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var database *gorm.DB
var e error

func DatabaseInit() {
	dsn := "host=db user=postgres password=postgres dbname=go_server port=5432 sslmode=disable"
	database, e = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if e != nil {
		log.Fatalf("Не удалось подключиться к базе данных: %v", e)
	}

	database.Exec(`CREATE EXTENSION IF NOT EXISTS "uuid-ossp";`)
	database.AutoMigrate(
		&model.User{},
		&model.Avatar{},
		&model.Like{},
		&model.Post{},
		&model.Role{},
		&model.Comment{},
		&model.Friend{},
	)
}

func DB() *gorm.DB {
	return database
}
