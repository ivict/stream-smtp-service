package tests

import (
	"testing"

	"github.com/redis/go-redis/v9"
)

func RestConnectTest(t *testing.T) {
	options := redis.Options{}
	// Instantiate client
	client := redis.NewClient(&options)
	if client == nil {
		t.Error("Invalid redis options")
		panic("No sense to continue")
	}
	defer client.Close()
}
