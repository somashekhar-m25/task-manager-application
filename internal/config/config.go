package config

import (
	"github.com/joho/godotenv"
	"github.com/somashekhar-m25/task-manager-application/internal/logger"
	"go.uber.org/zap"
)

func LoadEnv() {
	logger.ZapLogger.Info("loading .env file")
	err := godotenv.Load()
	if err != nil {
		logger.ZapLogger.Panic("failed to load .env file", zap.String("error: ", err.Error()))
	}
	logger.ZapLogger.Info("successfull loaded .env file")
}
