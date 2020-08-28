package log

import (
	"fmt"

	"github.com/spf13/viper"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	lumberjack "gopkg.in/natefinch/lumberjack.v2"
)

// no-lint
const (
	FlagLogFile = "log.file"
	FlagSize    = "log.size"
)

var logger Logger

func init() {
	// config := zap.NewProductionConfig()
	config := newConfig()

	callerSkip := zap.AddCallerSkip(1)
	l, err := config.Build(callerSkip)
	if err != nil {
		panic(fmt.Sprintf("log init failed: %v", err))
	}
	logger = &Log{logger: l}
}

// Config replace logger from config string
func Config(c LogConfig) {
	if c == nil {
		c = defaultConfig()
	}
	logger = c.NewLogger()
}

func defaultConfig() LogConfig {
	rolling := lumberjack.Logger{
		Filename:   viper.GetString(FlagLogFile),
		MaxSize:    viper.GetInt(FlagSize),
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
		EncodeLevel:    zapcore.CapitalLevelEncoder,
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

	var opts []zap.Option
	opts = append(opts, zap.AddCaller())
	opts = append(opts, zap.AddCallerSkip(1))
	opts = append(opts, zap.Development())
	// logger = zap.New(core,
	// 	caller, callerSkip, development)
	return &Conf{Core: core, Options: opts}
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
		// InitialFields: map[string]interface{}{"service": "logger"},
		// OutputPaths:      []string{"stdout", "./logger.log"},
		OutputPaths:      []string{"stdout"},
		ErrorOutputPaths: []string{"stderr"},
	}
	return config
}

// Load replace logger from config file
func Load(conf string) {

}

// Flush immediately processes all currently queued logs.
func Flush() {
	logger.Flush()
}

// Trace logs
func Trace(v ...interface{}) {
	logger.Trace(v)
}

// Debug logs
func Debug(v ...interface{}) {
	logger.Debug(v)
}

// Info logs
func Info(v ...interface{}) {
	logger.Info(v)
}

// Warn logs
func Warn(v ...interface{}) {
	logger.Warn(v)
}

// Error logs
func Error(v ...interface{}) {
	logger.Error(v)
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
	logger.Debugz(msg, fields...)
}

// Infoz formats logs
func Infoz(msg string, fields ...zap.Field) {
	logger.Infoz(msg, fields...)
}

// Warnz formats logs
func Warnz(msg string, fields ...zap.Field) {
	logger.Warnz(msg, fields...)
}

// Errorz formats logs
func Errorz(msg string, fields ...zap.Field) {
	logger.Errorz(msg, fields...)
}
