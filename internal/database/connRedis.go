package database

import (
	"context"
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/redis/go-redis/v9"
)

type RedisClient struct {
	Client *redis.Client
}

type RedisConfig struct {
	Addr         string
	Password     string
	DB           int
	PoolSize     int
	PoolTimeout  time.Duration
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
	DialTimeout  time.Duration
	MaxRetries   int
	MinIdleConns int
}

func NewRedisClient(addr, password string, db int) (*RedisClient, error) {
	config := RedisConfig{
		Addr:         addr,
		Password:     password,
		DB:           db,
		PoolSize:     100,
		PoolTimeout:  30 * time.Second,
		ReadTimeout:  3 * time.Second,
		WriteTimeout: 3 * time.Second,
		DialTimeout:  5 * time.Second,
		MaxRetries:   3,
		MinIdleConns: 10,
	}

	return NewRedisClientWithConfig(config)
}

func NewRedisClientWithConfig(config RedisConfig) (*RedisClient, error) {
	client := redis.NewClient(&redis.Options{
		Addr:         config.Addr,
		Password:     config.Password,
		DB:           config.DB,
		PoolSize:     config.PoolSize,
		PoolTimeout:  config.PoolTimeout,
		ReadTimeout:  config.ReadTimeout,
		WriteTimeout: config.WriteTimeout,
		DialTimeout:  config.DialTimeout,
		MaxRetries:   config.MaxRetries,
		MinIdleConns: config.MinIdleConns,

		MaxActiveConns:  config.PoolSize,
		ConnMaxIdleTime: 5 * time.Minute,
	})

	ctx, cancel := context.WithTimeout(context.Background(), config.DialTimeout)
	defer cancel()

	_, err := client.Ping(ctx).Result()
	if err != nil {
		return nil, fmt.Errorf("failed to connect to Redis: %w", err)
	}

	return &RedisClient{Client: client}, nil
}

func NewRateLimitRedisClient(addr, password string) (*RedisClient, error) {
	return NewRedisClient(addr, password, 1)
}

func NewRedisClientFromEnv(db int) (*RedisClient, error) {
	poolSize, _ := strconv.Atoi(getEnv("REDIS_POOL_SIZE", "100"))
	maxRetries, _ := strconv.Atoi(getEnv("REDIS_MAX_RETRIES", "3"))
	minIdleConns, _ := strconv.Atoi(getEnv("REDIS_MIN_IDLE_CONNS", "10"))

	config := RedisConfig{
		Addr:         getEnv("REDIS_ADDR", "localhost:6379"),
		Password:     getEnv("REDIS_PASSWORD", ""),
		DB:           db,
		PoolSize:     poolSize,
		PoolTimeout:  30 * time.Second,
		ReadTimeout:  3 * time.Second,
		WriteTimeout: 3 * time.Second,
		DialTimeout:  5 * time.Second,
		MaxRetries:   maxRetries,
		MinIdleConns: minIdleConns,
	}

	return NewRedisClientWithConfig(config)
}

func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}
