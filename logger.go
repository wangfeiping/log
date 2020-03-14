package log

import (
	"errors"
	"fmt"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var defaultConfig = ``

var rollingFileConfig = ``

var logger *zap.Logger

func init() {
	encoderConfig := zapcore.EncoderConfig{
		TimeKey:    "time",
		LevelKey:   "level",
		NameKey:    "logger",
		MessageKey: "msg",
		CallerKey:  "caller",
		// StacktraceKey:  "stacktrace",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.LowercaseLevelEncoder,
		EncodeTime:     zapcore.ISO8601TimeEncoder,
		EncodeDuration: zapcore.SecondsDurationEncoder,
		EncodeCaller:   zapcore.ShortCallerEncoder,
	}

	// 设置日志级别
	level := zap.NewAtomicLevelAt(zap.DebugLevel)

	config := zap.Config{
		Level:            level,
		Development:      true,
		Encoding:         "console", // "console/json",
		EncoderConfig:    encoderConfig,
		InitialFields:    map[string]interface{}{"service": "logger"},
		OutputPaths:      []string{"stdout", "./logger.log"},
		ErrorOutputPaths: []string{"stderr"},
	}

	var err error
	logger, err = config.Build()
	if err != nil {
		panic(fmt.Sprintf("log init failed: %v", err))
	}
}

// Config replace logger from config string
func Config(config string) {

}

// Load replace logger from config file
func Load(conf string) {

}

// RollingFileConfig returns config for rolling file
func RollingFileConfig() string {
	return rollingFileConfig
}

// Flush immediately processes all currently queued logs.
func Flush() {
	logger.Sync()
}

// Trace logs
func Trace(v ...interface{}) {
}

// Debug logs
func Debug(v ...interface{}) {
	msg, fields, err := parse(v...)
	if err != nil {
		fmt.Printf("log parse error: %v\n", err)
		return
	}
	logger.Debug(msg, fields...)
}

func parse(v ...interface{}) (string, []zap.Field, error) {
	if msg, ok := v[0].(string); ok {
		var fields []zap.Field
		return msg, fields, nil
	}
	return "", nil, errors.New("unable to build log: " + fmt.Sprint(v[1:]))

}

// Info logs
func Info(v ...interface{}) {
	msg, fields, err := parse(v...)
	if err != nil {
		fmt.Printf("log parse error: %v\n", err)
		return
	}
	logger.Info(msg, fields...)
}

// Warn logs
func Warn(v ...interface{}) {
	msg, fields, err := parse(v...)
	if err != nil {
		fmt.Printf("log parse error: %v\n", err)
		return
	}
	logger.Warn(msg, fields...)
}

// Error logs
func Error(v ...interface{}) {
	msg, fields, err := parse(v...)
	if err != nil {
		fmt.Printf("log parse error: %v\n", err)
		return
	}
	logger.Error(msg, fields...)
}

// Tracef logs
func Tracef(format string, params ...interface{}) {
}

// Debugf formats logs
func Debugf(format string, params ...interface{}) {
	logger.Debug(fmt.Sprintf(format, params...))
}

// Infof formats logs
func Infof(format string, params ...interface{}) {
	logger.Info(fmt.Sprintf(format, params...))
}

// Warnf formats logs
func Warnf(format string, params ...interface{}) {
	logger.Warn(fmt.Sprintf(format, params...))
}

// Errorf formats logs
func Errorf(format string, params ...interface{}) {
	logger.Error(fmt.Sprintf(format, params...))
}
