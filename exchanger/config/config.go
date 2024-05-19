package config

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)

type Config struct {
	RedisHost      string
	RedisPort      string
	CurrencyApiUrl string
	ServicePort    string
}

func LoadConfig() (*Config, error) {
	err := godotenv.Load()

	if err != nil {
		log.Fatal("Error loading .env file")
	}

	config := &Config{
		RedisHost:      os.Getenv("REDIS_HOST"),
		RedisPort:      os.Getenv("REDIS_PORT"),
		CurrencyApiUrl: os.Getenv("CURRENCY_API_URL"),
		ServicePort:    os.Getenv("EXCHANGER_SERVICE_PORT"),
	}

	return config, nil
}
