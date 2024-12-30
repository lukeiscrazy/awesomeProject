package database

import (
	"awesomeProject/models"
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

var DB *gorm.DB

func Connect() {
	dsn := "host=localhost user=postgres password=Lulu0516 dbname=postgres port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database: ", err)
	}

	DB = db

	// 自动迁移
	err = DB.AutoMigrate(
		&models.User{},
		&models.Follow{},
	)
	if err != nil {
		log.Fatal("Failed to migrate models: ", err)
	}

	fmt.Println("Database connected and migrated successfully!")
}
