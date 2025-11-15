package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/nycholasmarques/quick-start-gin/internal/database/sqlc"
	_ "github.com/nycholasmarques/quick-start-gin/internal/docs"
	"github.com/redis/go-redis/v9"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func SetupRoutes(router *gin.Engine, queries *sqlc.Queries, redis *redis.Client, rateLimitRedis *redis.Client) *gin.Engine {
	{
		router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
		v1 := router.Group("/api/v1/")
		v1.GET("health", healthCheck)
	}
	return router
}

// HealthCheck godoc
// @Summary Health check
// @Description Returns API status
// @Tags health
// @Accept json
// @Produce json
// @Success 200 {object} map[string]string
// @Router /health [get]
func healthCheck(ctx *gin.Context) {
	ctx.JSON(200, gin.H{
		"status": "ok",
	})
}
