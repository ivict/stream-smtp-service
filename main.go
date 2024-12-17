package main

import (
	"os"

	"github.com/Capstane/stream-mail-service/internal/config"
	"github.com/Capstane/stream-mail-service/internal/stream"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func main() {
	// Initialize Zerolog logger with output to stdout
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stdout})

	cfg := config.LoadConfig()
	zerolog.SetGlobalLevel(zerolog.Level(cfg.LogLevel))

	err := stream.ListenRedisStream(cfg)
	if err != nil {
		log.Error().Err(err)
	}
}
