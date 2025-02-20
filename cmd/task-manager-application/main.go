package main

import (
	"github.com/rs/zerolog/log"
	"github.com/somashekhar-m25/task-manager-application/internal/config"
	router "github.com/somashekhar-m25/task-manager-application/internal/routers"
)

func init() {
	config.LoadEnv()
}

func main() {
	log.Logger.Info().Msg("starting task-manager-application...")
	router.StartApplication()
	log.Logger.Info().Msg("stoping task-manager-application...")
}
