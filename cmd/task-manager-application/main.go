package main

import (
	"github.com/somashekhar-m25/task-manager-application/internal/logger"
	router "github.com/somashekhar-m25/task-manager-application/internal/routers"
)

func init() {
	logger.GetLogger()
}

func main() {
	logger.ZapLogger.Info("starting task-manager-application...")
	router.StartApplication()
}
