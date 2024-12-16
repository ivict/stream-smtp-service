package tests

import (
	"context"
	"testing"

	"github.com/Capstane/stream-auth-service/internal"
	"github.com/Capstane/stream-auth-service/internal/config"
	"github.com/redis/go-redis/v9"
)

func TestRestConnect(t *testing.T) {
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
