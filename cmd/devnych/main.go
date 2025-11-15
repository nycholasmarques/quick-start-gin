package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/nycholasmarques/quick-start-gin/config"
	"github.com/nycholasmarques/quick-start-gin/internal/database"
	"github.com/nycholasmarques/quick-start-gin/internal/database/sqlc"
	"github.com/nycholasmarques/quick-start-gin/internal/logger"
	"github.com/nycholasmarques/quick-start-gin/internal/routes"
	"github.com/redis/go-redis/v9"
	"go.uber.org/zap"
)

// @title quick-start-gin API
// @version 0.1
// @description API documentation for quick-start-gin
// @host localhost:8080
// @BasePath /api/v1
// @schemes http
func main() {
	env := os.Getenv("APP_ENV")
	if env == "" {
		env = "dev"
	}
	logger.Init(env)
	defer logger.Log.Sync()

	if os.Getenv("GIN_MODE") == "release" {
		gin.SetMode(gin.ReleaseMode)
	}

	cfg := config.LoadConfig()

	connPostgres, err := database.ConnPostgres(*cfg)
	if err != nil {
		logger.Log.Fatal("failed to connect to Postgres", zap.String("error", err.Error()))
	}

	err = connPostgres.Ping(context.Background())
	if err != nil {
		logger.Log.Error("Error to ping postgres", zap.String("error", err.Error()))
		return
	}

	queries := sqlc.New(connPostgres)

	var redisClient *database.RedisClient
	var rateLimitRedis *database.RedisClient
	if os.Getenv("REDIS_ENABLED") == "true" {
		rdb, err := database.NewRedisClient(
			os.Getenv("REDIS_ADDR"),
			os.Getenv("REDIS_PASSWORD"),
			0,
		)
		if err != nil {
			logger.Log.Error("Redis connection failed", zap.String("error", err.Error()))
		} else {
			redisClient = rdb
			defer rdb.Client.Close()
		}
		rateRdb, err := database.NewRateLimitRedisClient(
			os.Getenv("REDIS_ADDR"),
			os.Getenv("REDIS_PASSWORD"),
		)
		if err != nil {
			logger.Log.Error("Rate limit Redis connection failed", zap.String("error", err.Error()))
		} else {
			rateLimitRedis = rateRdb
			defer rateRdb.Client.Close()
		}
	} else {
		logger.Log.Info("Redis not enabled, skipping connection")
	}

	router := gin.Default()

	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "HEAD", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Length", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	router.Use(gin.Recovery())
	router.SetTrustedProxies(nil)
	router.Static("/uploads", "./uploads")

	var redisBase, rateLimiterRedis *redis.Client
	if redisClient != nil {
		redisBase = redisClient.Client
	}
	if rateLimitRedis != nil {
		rateLimiterRedis = rateLimitRedis.Client
	}

	router = routes.SetupRoutes(router, queries, redisBase, rateLimiterRedis)

	port := cfg.Port
	if port == "" {
		port = "8080"
	}
	srv := &http.Server{
		Addr:    ":" + port,
		Handler: router,
	}

	fmt.Println("server init in http://localhost:" + port)
	fmt.Println("doc init in http://localhost:" + port + "/swagger/index.html")

	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			logger.Log.Fatal("listen: ", zap.String("error", err.Error()))
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit
	logger.Log.Info("Shutdown Server ...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		logger.Log.Fatal("Server Shutdown", zap.String("error", err.Error()))
	}

	logger.Log.Info("Server exiting")
}
