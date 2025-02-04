package log

import (
	"fmt"
	"sync/atomic"

	"github.com/spf13/viper"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	lumberjack "gopkg.in/natefinch/lumberjack.v2"

	"github.com/wangfeiping/log/logger"
)

// no-lint
const (
	FlagLogFile   = "log.file"
	FlagLogSize   = "log.size"
	FlagLogBackup = "log.backup"
)

var log logger.Logger
var errorCount uint64 = 0

func init() {
	// config := zap.NewProductionConfig()
	config := newConfig()

	callerSkip := zap.AddCallerSkip(1)
	l, err := config.Build(callerSkip)
	if err != nil {
		panic(fmt.Sprintf("log init failed: %v", err))
	}
	log = logger.NewLog(l)
}

// Config replace logger from config string
func Config(c logger.LogConfig) {
	if c == nil {
		c = defaultConfig()
	}
	log = c.NewLogger()
}

func defaultConfig() logger.LogConfig {
	rolling := lumberjack.Logger{
		Filename:   viper.GetString(FlagLogFile),
		MaxSize:    viper.GetInt(FlagLogSize),
		MaxBackups: viper.GetInt(FlagLogBackup),
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
	return &logger.Conf{Core: core, Options: opts}
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
	log.Flush()
}

// Trace logs
func Trace(v ...interface{}) {
	log.Trace(v)
}

// Debug logs
func Debug(v ...interface{}) {
	log.Debug(v)
}

// Info logs
func Info(v ...interface{}) {
	log.Info(v)
}

// Warn logs
func Warn(v ...interface{}) {
	log.Warn(v)
}

// Error logs
func Error(v ...interface{}) {
	atomic.AddUint64(&errorCount, 1)
	log.Error(v)
}

// Tracef logs
func Tracef(format string, params ...interface{}) {
	log.Debug(fmt.Sprintf(format, params...))
}

// Debugf formats logs
func Debugf(format string, params ...interface{}) {
	log.Debug(fmt.Sprintf(format, params...))
}

// Infof formats logs
func Infof(format string, params ...interface{}) {
	log.Info(fmt.Sprintf(format, params...))
}

// Warnf formats logs
func Warnf(format string, params ...interface{}) {
	log.Warn(fmt.Sprintf(format, params...))
}

// Errorf formats logs
func Errorf(format string, params ...interface{}) {
	atomic.AddUint64(&errorCount, 1)
	log.Error(fmt.Sprintf(format, params...))
}

// Debugz formats logs
func Debugz(msg string, fields ...zap.Field) {
	log.Debugz(msg, fields...)
}

// Infoz formats logs
func Infoz(msg string, fields ...zap.Field) {
	log.Infoz(msg, fields...)
}

// Warnz formats logs
func Warnz(msg string, fields ...zap.Field) {
	log.Warnz(msg, fields...)
}

// Errorz formats logs
func Errorz(msg string, fields ...zap.Field) {
	atomic.AddUint64(&errorCount, 1)
	log.Errorz(msg, fields...)
}

// Output error logs and panic
func Panicf(format string, params ...interface{}) {
	err := fmt.Sprintf(format, params...)
	Error(err)
	Flush()
	panic(err)
}

func ErrorCount() uint64 {
	return atomic.LoadUint64(&errorCount)
}
