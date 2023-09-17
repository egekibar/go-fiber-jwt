package database

import (
	"fmt"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"os"
	"test-app/models"
)

var DB *gorm.DB

func Connect() {
	connection, err := gorm.Open(sqlite.Open("database/database.sqlite"), &gorm.Config{})

	if err != nil {
		panic("could not connect to the database")
	}

	DB = connection

	connection.AutoMigrate(&models.Users{})
}

func Init() {
	sqlitePath := "database/database.sqlite"
	if _, err := os.Stat(sqlitePath); os.IsNotExist(err) {
		dosya, err := os.Create(sqlitePath)
		if err != nil {
			fmt.Println("Veritabanı dosyası oluşturulmadı:", err)
			return
		}
		defer dosya.Close()
	}
	return
}
