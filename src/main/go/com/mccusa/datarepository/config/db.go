package config

import (
	"fmt"
	"go-data-repository/src/main/go/com/mccusa/datarepository/model"
	"log"

	"github.com/jmoiron/sqlx"
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

func ConnectCustomDb(cfg *Config) *sqlx.DB {
	// Connect to the database with sqlx
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=require",
		cfg.DBHost,
		cfg.DBPort,
		cfg.DBUser,
		cfg.DBPassword,
		cfg.DBName,
	)
	customDb, err := sqlx.Connect("postgres", dsn)
	if err != nil {
		log.Fatalf("Failed to connect to database with sqlx: %v", err)
	}
	log.Printf("Connected to database at %s", customDb.DriverName())
	return customDb
}
