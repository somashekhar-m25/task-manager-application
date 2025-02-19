package database

import (
	"errors"
	"fmt"
	"os"

	"github.com/somashekhar-m25/task-manager-application/internal/logger"
	"go.uber.org/zap"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type DataBase struct {
	DB *gorm.DB
}

func ConnectToDB() (*DataBase, error) {
	//read database credentials from env
	dbHost := os.Getenv("DB_HOST")
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")
	dbPort := os.Getenv("DB_PORT")

	//postgres database source name
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", dbHost, dbUser, dbPassword, dbName, dbPort)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		logger.ZapLogger.Error("database connection failed", zap.String("dsn: ", dsn), zap.String("error: ", err.Error()))
		return nil, errors.New("database connection failed")
	}

	//get db instance
	sqlDB, err := db.DB()
	if err != nil {
		logger.ZapLogger.Error("failed to get database instance", zap.String("error: ", err.Error()))
		return nil, errors.New("failed to get database instance")
	}

	//ping db to check connection alive
	if err = sqlDB.Ping(); err != nil {
		logger.ZapLogger.Error("failed to ping database", zap.String("error: ", err.Error()))
		return nil, errors.New("failed to ping database")
	}

	logger.ZapLogger.Info("database connection successfull")
	return &DataBase{
		DB: db,
	}, nil
}
