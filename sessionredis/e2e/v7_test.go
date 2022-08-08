package e2e_test

import (
	"os"
	"testing"

	"github.com/go-redis/redis/v9"
	"github.com/mr-linch/go-tg-store/sessionredis"
	"github.com/stretchr/testify/assert"
)

func TestV7(t *testing.T) {
	redisDSN := os.Getenv("REDIS_V7_DSN")
	if redisDSN == "" {
		t.Skip("REDIS_V7_DSN is not set")
	}

	opts, err := redis.ParseURL(redisDSN)
	assert.NoError(t, err)

	client := redis.NewClient(opts)

	store := sessionredis.NewStore[*redis.StatusCmd, *redis.StringCmd, *redis.IntCmd](client)
	logic(t, store)
}
