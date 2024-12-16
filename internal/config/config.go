package config

import (
	"os"
	"strconv"

	"github.com/joho/godotenv"
	"github.com/rs/zerolog/log"
)

type Config struct {
	LogLevel  int    `default:"4" envconfig:"LOG_LEVEL"`
	RedisHost string `binding:"required" envconfig:"REDIS_HOST"`
	RedisPort string `binding:"required" envconfig:"REDIS_PORT"`
	SmtpHost  string `binding:"required" envconfig:"SMTP_HOST"`
	SmtpPort  string `binding:"required" envconfig:"SMTP_PORT"`
}

func LoadConfig() Config {
	err := godotenv.Load()
	if err != nil {
		log.Info().Msg("Error loading .env file")
	}

	logLevel, _ := strconv.Atoi(os.Getenv("LOG_LEVEL"))

	return Config{
		LogLevel:  logLevel,
		RedisHost: os.Getenv("REDIS_HOST"),
		RedisPort: os.Getenv("REDIS_PORT"),
		SmtpHost:  os.Getenv("SMTP_HOST"),
		SmtpPort:  os.Getenv("SMTP_PORT"),
	}
}
