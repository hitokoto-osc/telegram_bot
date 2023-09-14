package logging

import (
	"go.uber.org/zap"
)

const loggerKey = "logger"

var logger *zap.Logger

// GetLogger returns a global logger
func GetLogger() *zap.Logger {
	return logger
}

func setZapGlobalLogger() {
	zap.ReplaceGlobals(logger)
}
