package logger

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
		log: zap.New(conf.Core, conf.Options...)}
}

func NewLog(log *zap.Logger) Logger {
	return &Log{log: log}
}

type Log struct {
	log *zap.Logger
}

func (l *Log) Flush() error {
	return l.log.Sync()
}

func (l *Log) Trace(v ...interface{}) {
	l.log.Debug(fmt.Sprint(v...))
}

func (l *Log) Debug(v ...interface{}) {
	l.log.Debug(fmt.Sprint(v...))
}

// Info logs
func (l *Log) Info(v ...interface{}) {
	l.log.Info(fmt.Sprint(v...))
}

// Warn logs
func (l *Log) Warn(v ...interface{}) {
	l.log.Warn(fmt.Sprint(v...))
}

// Error logs
func (l *Log) Error(v ...interface{}) {
	l.log.Error(fmt.Sprint(v...))
}

// Debugz formats logs
func (l *Log) Debugz(msg string, fields ...zap.Field) {
	l.log.Debug(msg, fields...)
}

// Infoz formats logs
func (l *Log) Infoz(msg string, fields ...zap.Field) {
	l.log.Info(msg, fields...)
}

// Warnz formats logs
func (l *Log) Warnz(msg string, fields ...zap.Field) {
	l.log.Warn(msg, fields...)
}

// Errorz formats logs
func (l *Log) Errorz(msg string, fields ...zap.Field) {
	l.log.Error(msg, fields...)
}
