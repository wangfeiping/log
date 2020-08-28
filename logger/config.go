package logger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type LogConfig interface {
	NewLogger() Logger
}

type Conf struct {
	Core    zapcore.Core
	Options []zap.Option
}

func (c *Conf) NewLogger() Logger {
	return NewLogger(c)
}
