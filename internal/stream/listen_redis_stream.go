package stream

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/Capstane/stream-auth-service/internal"
	"github.com/Capstane/stream-auth-service/internal/config"
	"github.com/redis/go-redis/v9"
	"github.com/rs/zerolog/log"
)

func printLogo() {
	fmt.Println("STREAMS SMTP SERVICE V1")
}

func ListenRedisStream(cfg config.Config) error {

	options := redis.Options{
		Addr: internal.FormatAddr(cfg.RedisHost, cfg.RedisPort),
	}

	// Instantiate client
	client := redis.NewClient(&options)
	if client == nil {
		return fmt.Errorf("invalid redis options, for configuration: %v", cfg)
	}
	defer client.Close()

	ctx := context.Background()

	// Check that stream exists
	existedStreams, _, err := client.Scan(ctx, 0, cfg.RedisStream, -1).Result()
	if err != nil {
		return fmt.Errorf("can't read streams %v: %v", cfg.RedisStream, err)
	}
	if len(existedStreams) == 0 {
		return fmt.Errorf("stream %v doesn't exists", cfg.RedisStream)
	}

	signalChannel := make(chan os.Signal, 1)
	signal.Notify(
		signalChannel,
		os.Interrupt,
		syscall.SIGHUP,
		syscall.SIGQUIT,
		syscall.SIGTERM,
		syscall.SIGSEGV,
	)

	go processStream(client, ctx, &cfg)

	signalEvent := <-signalChannel
	switch signalEvent {
	case syscall.SIGQUIT,
		syscall.SIGTERM,
		syscall.SIGINT,
		syscall.SIGKILL:
		return nil
	case syscall.SIGHUP:
		return fmt.Errorf("signal hang up")
	case syscall.SIGSEGV:
		return fmt.Errorf("segmentation violation")
	}
	return nil
}

func processStream(client *redis.Client, ctx context.Context, cfg *config.Config) error {
	defer func() {
		syscall.Kill(syscall.Getpid(), syscall.SIGINT)
	}()

	// MailSenders channel
	chMailSenders := make(chan redis.XMessage, cfg.MaxSimultaneousSmtpConnections)
	go func() {
		for streamMessage := range chMailSenders {
			smtpMessage, err := SmtpMessageUnmarshal(streamMessage.Values)
			if err != nil {
				log.Error().Err(err)
				continue
			}

			err = sendEmail(*smtpMessage, cfg)
			if err == nil {
				_, err := client.XDel(ctx, cfg.RedisStream, streamMessage.ID).Result()
				if err != nil {
					log.Error().Err(err)
				}
			}
		}
	}()

	printLogo()
	circularBuffer := NewCircularBuffer(cfg.MaxSimultaneousSmtpConnections)
	for {
		xStreams, err := client.XRead(ctx, &redis.XReadArgs{
			Streams: []string{cfg.RedisStream},
			Block:   0,
			ID:      "+",
			Count:   int64(cfg.MaxSimultaneousSmtpConnections),
		}).Result()

		if err != nil {
			return fmt.Errorf("can't read stream %v: %v", cfg.RedisStream, err)
		}

		for _, xStream := range xStreams {
			for _, message := range xStream.Messages {
				if !circularBuffer.Has(message.ID) {
					chMailSenders <- message
				}
				circularBuffer.Put(message.ID)
			}
		}

	}
}
