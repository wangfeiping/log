package log

import (
	"bytes"
	"errors"
	"fmt"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	lumberjack "gopkg.in/natefinch/lumberjack.v2"
)

var defaultConfig = ``

var rollingFileConfig = `zap:defaultRollingFile`

var logger *zap.Logger

func init() {
	// config := zap.NewProductionConfig()
	config := newConfig()

	var err error
	logger, err = config.Build()
	if err != nil {
		panic(fmt.Sprintf("log init failed: %v", err))
	}
}

// Config replace logger from config string
func Config(config string) {
	if config == rollingFileConfig {
		rolling := lumberjack.Logger{
			Filename:   "./logger.log",
			MaxSize:    512,
			MaxBackups: 10,
			MaxAge:     7,
			Compress:   true,
		}

		encoderConfig := zapcore.EncoderConfig{
			TimeKey:    "time",
			LevelKey:   "level",
			NameKey:    "logger",
			CallerKey:  "caller",
			MessageKey: "msg",
			// StacktraceKey:  "stacktrace",
			LineEnding:     zapcore.DefaultLineEnding,
			EncodeLevel:    zapcore.CapitalColorLevelEncoder,
			EncodeTime:     zapcore.ISO8601TimeEncoder,
			EncodeDuration: zapcore.SecondsDurationEncoder,
			EncodeCaller:   zapcore.ShortCallerEncoder,
			EncodeName:     zapcore.FullNameEncoder,
		}

		level := zap.NewAtomicLevelAt(zap.DebugLevel)

		core := zapcore.NewCore(
			zapcore.NewJSONEncoder(encoderConfig),
			zapcore.NewMultiWriteSyncer(
				// zapcore.AddSync(os.Stdout),
				zapcore.AddSync(&rolling)),
			level,
		)

		caller := zap.AddCaller()
		development := zap.Development()
		filed := zap.Fields(zap.String("service", "logger"))
		logger = zap.New(core, caller, development, filed)
	}
}

func newConfig() zap.Config {
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

	// log level
	level := zap.NewAtomicLevelAt(zap.DebugLevel)

	config := zap.Config{
		Level:         level,
		Development:   true,
		Encoding:      "console", // "console/json",
		EncoderConfig: encoderConfig,
		InitialFields: map[string]interface{}{"service": "logger"},
		// OutputPaths:      []string{"stdout", "./logger.log"},
		OutputPaths:      []string{"stdout"},
		ErrorOutputPaths: []string{"stderr"},
	}
	return config
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
		fmt.Printf("log error(debug): %v\n", err)
		return
	}
	logger.Debug(msg, fields...)
}

// Info logs
func Info(v ...interface{}) {
	msg, fields, err := parse(v...)
	if err != nil {
		fmt.Printf("log error(info): %v\n", err)
		return
	}
	logger.Info(msg, fields...)
}

// Warn logs
func Warn(v ...interface{}) {
	msg, fields, err := parse(v...)
	if err != nil {
		fmt.Printf("log error(warn): %v\n", err)
		return
	}
	logger.Warn(msg, fields...)
}

// Error logs
func Error(v ...interface{}) {
	msg, fields, err := parse(v...)
	if err != nil {
		fmt.Printf("log error(error): %v\n", err)
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

// DebugZ formats logs
func DebugZ(msg string, fields ...zap.Field) {
	logger.Debug(msg, fields...)
}

// InfoZ formats logs
func InfoZ(msg string, fields ...zap.Field) {
	logger.Info(msg, fields...)
}

// WarnZ formats logs
func WarnZ(msg string, fields ...zap.Field) {
	logger.Warn(msg, fields...)
}

// ErrorZ formats logs
func ErrorZ(msg string, fields ...zap.Field) {
	logger.Error(msg, fields...)
}

func parse(v ...interface{}) (string, []zap.Field, error) {
	if msg, ok := v[0].(string); ok {
		var fields []zap.Field
		if len(v)%2 == 1 {
			return msg, fields, nil
		}
	}
	return "", nil, errors.New(join("unable to build log:", v[1:]...))
}

func join(msg string, v ...interface{}) string {
	var buf bytes.Buffer
	buf.WriteString(msg)
	for _, s := range v {
		buf.WriteString(" ")
		buf.WriteString(fmt.Sprint(s))
	}
	return buf.String()
}
