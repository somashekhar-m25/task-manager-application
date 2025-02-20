package config

import (
	"path/filepath"

	"github.com/joho/godotenv"
	"github.com/somashekhar-m25/task-manager-application/internal/logger"
	"go.uber.org/zap"
)

func LoadEnv() {
	logger.ZapLogger.Info("loading .env file")
	envPath, err := filepath.Abs("../../")
	if err != nil {
		logger.ZapLogger.Error("error getting absolute .env path", zap.String("error: ", err.Error()))
	}
	envPath = filepath.Join(envPath, "internal", "config", "application.env")
	err = godotenv.Load(envPath)
	if err != nil {
		logger.ZapLogger.Panic("failed to load .env file", zap.String("error: ", err.Error()))
	}
	logger.ZapLogger.Info("successfull loaded .env file")
}
