package redis

import (
	"log"
	"net"
	"os"
	"time"

	"github.com/bungysheep/news-api/pkg/configs"
	"github.com/go-redis/redis/v7"
)

var (
	// RedisClient - Redis client
	RedisClient *redis.Client
)

// CreateRedisClient - Creates redis client
func CreateRedisClient() error {
	log.Printf("Creating redis client...")

	redisURL, err := resolveRedisURL()
	if err != nil {
		return err
	}

	resolveRedisAuth, err := resolveRedisAuth()
	if err != nil {
		return err
	}

	client := redis.NewClient(&redis.Options{
		Addr:     redisURL,
		Password: resolveRedisAuth,
		DB:       0,
	})

	RedisClient = client

	for i := 0; i < configs.NUMBERDIALATTEMPT; i++ {
		_, err = client.Ping().Result()
		if err != nil {
			opErr, ok := err.(*net.OpError)
			if !ok || opErr.Op != "dial" {
				return err
			}
			time.Sleep(5 * time.Second)
		}
	}

	return err
}

func resolveRedisURL() (string, error) {
	redisURL := os.Getenv("REDIS_URL")
	if redisURL != "" {
		return redisURL, nil
	}

	return configs.REDISURL, nil
}

func resolveRedisAuth() (string, error) {
	redisAuth := os.Getenv("REDIS_AUTH")
	if redisAuth != "" {
		return redisAuth, nil
	}

	return configs.REDISAUTH, nil
}
