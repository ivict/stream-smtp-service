package config

import (
	"os"
	"strconv"

	"github.com/joho/godotenv"
	"github.com/rs/zerolog/log"
)

type Config struct {
	LogLevel  int8   `default:"4" envconfig:"LOG_LEVEL"`
	RedisHost string `binding:"required" envconfig:"REDIS_HOST"`
	RedisPort string `binding:"required" envconfig:"REDIS_PORT"`
	SmtpHost  string `binding:"required" envconfig:"SMTP_HOST"`
	SmtpPort  string `binding:"required" envconfig:"SMTP_PORT"`
	SmtpFrom  string `binding:"required" envconfig:"SMTP_FROM"`
}

func LoadConfig() Config {
	err := godotenv.Load()
	if err != nil {
		log.Info().Msg("Error loading .env file")
	}

	logLevel, _ := strconv.Atoi(os.Getenv("LOG_LEVEL"))

	return Config{
		LogLevel:  int8(logLevel),
		RedisHost: os.Getenv("REDIS_HOST"),
		RedisPort: os.Getenv("REDIS_PORT"),
		SmtpHost:  os.Getenv("SMTP_HOST"),
		SmtpPort:  os.Getenv("SMTP_PORT"),
		SmtpFrom:  os.Getenv("SMTP_FROM"),
	}
}
