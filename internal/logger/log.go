package logger

import (
	"sync"

	"go.uber.org/zap"
)

var (
	ZapLogger *zap.Logger
	once      sync.Once
)

func GetLogger() *zap.Logger {
	once.Do(func() {
		logger, err := zap.NewProduction()
		if err != nil {
			panic("failed to initialize logger: " + err.Error())
		}
		ZapLogger = logger
	})
	return ZapLogger
}
