package repository

import (
	"fmt"
	"log"

	"web/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Config struct {
	Host     string
	Port     string
	Username string
	Password string
	DBName   string
	SSLMode  string
}

var DB *gorm.DB

func ConnectDB() {
	cfg := Config{
		Host:     "127.0.0.1",
		Port:     "5433",
		Username: "postgres",
		Password: "postgres",
		DBName:   "postgres",
		SSLMode:  "disable",
	}
	pcd := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=%s",
		cfg.Host, cfg.Port, cfg.Username, cfg.DBName, cfg.Password, cfg.SSLMode)

	db, err := gorm.Open(postgres.Open(pcd), &gorm.Config{})
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println("Connect to DB: Successfully!")

	db.AutoMigrate(&models.Picture{})

	DB = db
}
