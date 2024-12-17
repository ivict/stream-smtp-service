package tests

import (
	"context"
	"os"
	"strconv"
	"testing"
	"time"

	"github.com/Capstane/stream-mail-service/internal"
	"github.com/Capstane/stream-mail-service/internal/config"
	"github.com/Capstane/stream-mail-service/internal/stream"
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

	millis := time.Now().UnixNano() / 1e6
	messageId := strconv.Itoa(int(millis)) + "-*"

	ctx := context.Background()
	_, err := client.XAdd(ctx, &redis.XAddArgs{
		Stream: cfg.RedisStream,
		MaxLen: 1,
		Values: stream.SmtpMessage{
			Type:    stream.MailPlain.String(),
			Subject: "TestSendEmailMessageAsJson",
			Text:    "Lorem Ipsum",
			To:      os.Getenv("TEST_SMTP_TO"),
		}.Marshal(),
		ID: messageId,
	}).Result()
	if err != nil {
		t.Error(err)
	}
}
