package database

import (
	"errors"
	"fmt"
	"os"

	"github.com/rs/zerolog/log"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectToDB() (*gorm.DB, error) {
	log.Logger.Info().Msg("initiating database connection...")
	//read database credentials from env
	dbHost := os.Getenv("DB_HOST")
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")
	dbPort := os.Getenv("DB_PORT")

	//postgres database source name
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s", dbHost, dbUser, dbPassword, dbName, dbPort)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Logger.Error().Err(err).Msg("database connection failed")
		return nil, errors.New("database connection failed")
	}

	//get db instance
	sqlDB, err := db.DB()
	if err != nil {
		log.Logger.Error().Err(err).Msg("failed to get database instance")
		return nil, errors.New("failed to get database instance")
	}

	//ping db to check connection alive
	if err = sqlDB.Ping(); err != nil {
		log.Logger.Error().Err(err).Msg("failed to ping database")
		return nil, errors.New("failed to ping database")
	}

	log.Logger.Info().Msg("database connection successfull...")
	return db, nil
}
