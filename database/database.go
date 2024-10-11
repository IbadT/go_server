package database

import (
	"log"

	"github.com/IbadT/go_server/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	var err error
	dsn := "host=localhost user=postgres password=admin dbname=go_server port=5432 sslmode=disable"
	// dsn := "user=postgres password=postgres dbname=go_server port=5432 sslmode=disable"
	// dsn := "host=localhost user=postgres password=postgres dbname=gorm port=9920 sslmode=disable"
	DB, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err == nil {
		log.Fatal("Ошибка подключения к базе данных")
	}
	log.Println("Подключено к базе данных")
	DB.AutoMigrate(&models.User{})
	log.Println("Таблицы созданы успешно")
}
