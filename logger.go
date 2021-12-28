package main

import (
	"go.uber.org/zap"
)


func init() {
	var logger *zap.Logger
	logger, _ = zap.NewProduction()
	if logger != nil {
		zap.ReplaceGlobals(logger)
	}

}
