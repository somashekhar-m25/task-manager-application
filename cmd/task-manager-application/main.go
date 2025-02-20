package main

import (
	"github.com/somashekhar-m25/task-manager-application/internal/config"
	"github.com/somashekhar-m25/task-manager-application/internal/logger"
	router "github.com/somashekhar-m25/task-manager-application/internal/routers"
)

func init() {
	logger.GetLogger()
	config.LoadEnv()
}

func main() {
	logger.ZapLogger.Info("starting task-manager-application...")
	router.StartApplication()
	logger.ZapLogger.Info("closing task-manager-application...")
}
