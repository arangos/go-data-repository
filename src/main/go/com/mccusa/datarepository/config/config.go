package config

import (
	"fmt"

	"github.com/spf13/viper"
)

// Config holds all application‐wide settings.
type Config struct {
	ServerPort string // e.g. "8080"
	BaseURL    string // base URL for the application
	APIKey     string // ActiveCampaign API token
	DBHost     string // database host
	DBPort     string // database port
	DBUser     string // database user
	DBPassword string // database password
	DBName     string // database name
}

// LoadConfig reads configuration from config.yaml
func LoadConfig() (*Config, error) {
	v := viper.New()

	// Basic viper config
	v.SetConfigName("application-prod")                      // name of config file (without extension)
	v.SetConfigType("yaml")                                  // YAML format
	v.AddConfigPath("../../../../../../src/main/resources/") // absolute path
	//v.AddConfigPath(".")                                     // only look in current directory

	// Environment variables
	v.AutomaticEnv()

	if err := v.ReadInConfig(); err != nil {
		fmt.Printf("⚠️  Error reading config file: %v\n", err)
		return nil, fmt.Errorf("failed to read config: %w", err)
	}

	fmt.Printf("Using config file: %s\n", v.ConfigFileUsed())

	cfg := &Config{
		DBHost:     v.GetString("database.host"),
		DBPort:     v.GetString("database.port"),
		DBUser:     v.GetString("database.username"),
		DBPassword: v.GetString("database.password"),
		DBName:     v.GetString("database.schema"),
		ServerPort: v.GetString("server.port"),
		BaseURL:    v.GetString("mccusa.baseUrl"),
		APIKey:     v.GetString("activecampaign.api.key"),
	}

	// Debug output
	fmt.Printf("Config values:\n")
	fmt.Printf("ServerPort: %s\n", cfg.ServerPort)
	fmt.Printf("BaseURL: %s\n", cfg.BaseURL)
	fmt.Printf("APIKey length: %d\n", len(cfg.APIKey))

	// Validate required fields
	if cfg.DBHost == "" {
		return nil, fmt.Errorf("database.url must be set")
	}
	if cfg.APIKey == "" {
		return nil, fmt.Errorf("activecampaign.api.key must be set")
	}
	return cfg, nil
}
