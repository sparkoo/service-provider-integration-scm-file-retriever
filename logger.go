package main

import (
	"go.uber.org/zap"
	"log"
)

var logger *log.Logger

func init() {
	var logger *zap.Logger
	logger, _ = zap.NewProduction()
	if logger != nil {
		zap.ReplaceGlobals(logger)
	}

}
