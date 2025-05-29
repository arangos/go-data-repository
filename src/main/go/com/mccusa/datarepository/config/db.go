package config

import (
	"fmt"
	"go-data-repository/src/main/go/com/mccusa/datarepository/model"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect(cfg *Config) *gorm.DB {
	var err error

	// Build connection string with SSL mode
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=require",
		cfg.DBHost,
		cfg.DBPort,
		cfg.DBUser,
		cfg.DBPassword,
		cfg.DBName)

	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("failed to connect to DB: %v", err)
	}
	// Auto-migrate your models
	if err := DB.AutoMigrate(&model.Client{}); err != nil {
		log.Fatalf("AutoMigrate failed: %v", err)
	}
	return DB
}
