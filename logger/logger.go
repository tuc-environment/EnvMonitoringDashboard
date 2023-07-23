package logger

import (
	"go.uber.org/zap"
)

type Logger struct {
	*zap.Logger
}

func NewLogger() (*Logger, error) {
	logger, err := zap.NewDevelopment()
	return &Logger{logger}, err
}
