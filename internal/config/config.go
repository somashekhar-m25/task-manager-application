package config

import (
	"path/filepath"

	"github.com/joho/godotenv"
	"github.com/rs/zerolog/log"
)

func LoadEnv() {
	log.Logger.Info().Msg("loadin .env files...")
	envPath, err := filepath.Abs("../../")
	if err != nil {
		log.Logger.Error().Err(err).Msg("error getting absolute .env path")
	}
	envPath = filepath.Join(envPath, "internal", "config", "application.env")
	err = godotenv.Load(envPath)
	if err != nil {
		log.Logger.Info().Msg("failed to load .env file")
	}
	log.Logger.Info().Msg("successfull loaded .env file...")
}
