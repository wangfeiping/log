package log

import (
	"fmt"

	"go.uber.org/zap"
)

type Logger interface {
	Flush() error
	Trace(v ...interface{})
	Debug(v ...interface{})
	Info(v ...interface{})
	Warn(v ...interface{})
	Error(v ...interface{})
	Debugz(msg string, fields ...zap.Field)
	Infoz(msg string, fields ...zap.Field)
	Warnz(msg string, fields ...zap.Field)
	Errorz(msg string, fields ...zap.Field)
}

func NewLogger(conf *Conf) Logger {
	return &Log{
		logger: zap.New(conf.Core, conf.Options...)}
}

type Log struct {
	logger *zap.Logger
}

func (l *Log) Flush() error {
	return l.logger.Sync()
}

func (l *Log) Trace(v ...interface{}) {
	l.logger.Debug(fmt.Sprint(v...))
}

func (l *Log) Debug(v ...interface{}) {
	l.logger.Debug(fmt.Sprint(v...))
}

// Info logs
func (l *Log) Info(v ...interface{}) {
	l.logger.Info(fmt.Sprint(v...))
}

// Warn logs
func (l *Log) Warn(v ...interface{}) {
	l.logger.Warn(fmt.Sprint(v...))
}

// Error logs
func (l *Log) Error(v ...interface{}) {
	l.logger.Error(fmt.Sprint(v...))
}

// Debugz formats logs
func (l *Log) Debugz(msg string, fields ...zap.Field) {
	l.logger.Debug(msg, fields...)
}

// Infoz formats logs
func (l *Log) Infoz(msg string, fields ...zap.Field) {
	l.logger.Info(msg, fields...)
}

// Warnz formats logs
func (l *Log) Warnz(msg string, fields ...zap.Field) {
	l.logger.Warn(msg, fields...)
}

// Errorz formats logs
func (l *Log) Errorz(msg string, fields ...zap.Field) {
	l.logger.Error(msg, fields...)
}
