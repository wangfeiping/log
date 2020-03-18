package log

import (
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
        
        callerSkip := zap.AddCallerSkip(1)
	var err error
	logger, err = config.Build(callerSkip)
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
			// EncodeName:     zapcore.FullNameEncoder,
			EncodeCaller: zapcore.ShortCallerEncoder,
		}

		level := zap.NewAtomicLevelAt(zap.DebugLevel)

		core := zapcore.NewCore(
			// zapcore.NewJSONEncoder(encoderConfig),
			zapcore.NewConsoleEncoder(encoderConfig),
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
		EncodeLevel:    zapcore.CapitalColorLevelEncoder,
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
	logger.Debug(fmt.Sprint(v...))
}

// Debug logs
func Debug(v ...interface{}) {
	logger.Debug(fmt.Sprint(v...))
}

// Info logs
func Info(v ...interface{}) {
	logger.Info(fmt.Sprint(v...))
}

// Warn logs
func Warn(v ...interface{}) {
	logger.Warn(fmt.Sprint(v...))
}

// Error logs
func Error(v ...interface{}) {
	logger.Error(fmt.Sprint(v...))
}

// Tracef logs
func Tracef(format string, params ...interface{}) {
	logger.Debug(fmt.Sprintf(format, params...))
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

// Debugz formats logs
func Debugz(msg string, fields ...zap.Field) {
	logger.Debug(msg, fields...)
}

// Infoz formats logs
func Infoz(msg string, fields ...zap.Field) {
	logger.Info(msg, fields...)
}

// Warnz formats logs
func Warnz(msg string, fields ...zap.Field) {
	logger.Warn(msg, fields...)
}

// Errorz formats logs
func Errorz(msg string, fields ...zap.Field) {
	logger.Error(msg, fields...)
}
