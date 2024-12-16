package tests

import (
	"context"
	"os"
	"testing"

	"github.com/Capstane/stream-auth-service/internal"
	"github.com/Capstane/stream-auth-service/internal/config"
	"github.com/Capstane/stream-auth-service/internal/topic"
	"github.com/redis/go-redis/v9"
)

func TestRedisConnect(t *testing.T) {
	cfg := config.LoadConfig()

	options := redis.Options{
		Addr: internal.FormatAddr(cfg.RedisHost, cfg.RedisPort),
	}
	// Instantiate client
	client := redis.NewClient(&options)
	if client == nil {
		t.Error("Invalid redis options")
		panic("No sense to continue")
	}
	defer client.Close()

	ctx := context.Background()
	_, err := client.Ping(ctx).Result()
	if err != nil {
		t.Error(err)
	}
}

func TestSendEmailMessageAsJson(t *testing.T) {
	cfg := config.LoadConfig()

	options := redis.Options{
		Addr: internal.FormatAddr(cfg.RedisHost, cfg.RedisPort),
	}
	// Instantiate client
	client := redis.NewClient(&options)
	if client == nil {
		t.Error("Invalid redis options")
		panic("No sense to continue")
	}
	defer client.Close()

	ctx := context.Background()
	_, err := client.Publish(ctx, cfg.RedisTopic, &topic.SmtpMessage{
		Type: topic.MailPlain.String(),
		Text: "Lorem Ipsum",
		To:   os.Getenv("TEST_SMTP_TO"),
	}).Result()

	if err != nil {
		t.Error(err)
	}
}
