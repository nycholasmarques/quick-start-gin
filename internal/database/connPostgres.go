package database

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5"
	"github.com/nycholasmarques/quick-start-gin/config"
	"github.com/nycholasmarques/quick-start-gin/internal/logger"
	"go.uber.org/zap"
)

func ConnPostgres(cfg config.Config) (*pgx.Conn, error) {
	connStr := fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=%s",
		cfg.Database.Host,
		cfg.Database.Port,
		cfg.Database.User,
		cfg.Database.Password,
		cfg.Database.DBName,
		cfg.Database.SSLMode,
	)

	conn, err := pgx.Connect(context.Background(), connStr)
	if err != nil {
		logger.Log.Error("Error to connect database", zap.String("error", err.Error()))
		return nil, err
	}
	return conn, nil
}
