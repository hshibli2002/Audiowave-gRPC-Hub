package config

import (
	"github.com/joho/godotenv"
	"github.com/spf13/viper"
	"log"
)

type Config struct {
	DBHost     string
	DBPort     string
	DBUser     string
	DBPassword string
	DBName     string
}

func LoadConfig() (*Config, error) {
	// Load the .env file
	if err := godotenv.Load(); err != nil {
		log.Printf("No .env file found")
	}

	viper.AutomaticEnv() // Override values from environment variables

	var config Config
	config.DBHost = viper.GetString("DB_HOST")
	config.DBPort = viper.GetString("DB_PORT")
	config.DBUser = viper.GetString("DB_USER")
	config.DBPassword = viper.GetString("DB_PASSWORD")
	config.DBName = viper.GetString("DB_NAME")

	return &config, nil
}
