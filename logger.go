package log

import (
	"github.com/cihub/seelog"
)

var defaultConfig = `
<seelog>
	<outputs>
	<rollingfile formatid="fmt" type="date"
	 filename="./logger.log"
	 fullname="false" datepattern="20060102" maxrolls="7" />
	</outputs>
	<formats>
		<format id="fmt" format="%L [%Date(2006-01-02 15:04:05.000000000)] %Msg%n"/>
	</formats>
</seelog>`

func init() {
	err := log.RegisterCustomFormatter("L", logLevelFormatter)
	if err != nil {
		ErrorD("log init failed: ", err)
	}
	logger, _ := seelog.LoggerFromConfigAsBytes([]byte(defaultConfig))
	Replace(logger)
}

// ReplaceConfig replace logger from new config string
func ReplaceConfig(config string) {
	logger, _ := seelog.LoggerFromConfigAsBytes([]byte(config))
	Replace(logger)
}

// LoadLogger 通过配置文件初始化日志模块
func LoadLogger(conf string) (seelog.LoggerInterface, error) {
	return seelog.LoggerFromConfigAsFile(conf)
}

// Replace logger
func Replace(logger seelog.LoggerInterface) {
	seelog.ReplaceLogger(logger)
}

// Flush immediately processes all currently queued logs.
func Flush() {
	seelog.Flush()
}

// Trace logs 详细运行跟踪日志，可能影响程序性能，所以生产环境不配置输出，仅在开发测试环境使用
func Trace(v ...interface{}) {
	seelog.Trace(v...)
}

// Debug logs
func Debug(v ...interface{}) {
	seelog.Debug(v...)
}

// Info logs
func Info(v ...interface{}) {
	seelog.Info(v...)
}

// Warn logs
func Warn(v ...interface{}) {
	seelog.Warn(v...)
}

// Error logs
func Error(v ...interface{}) {
	seelog.Error(v...)
	//seelog.Error(fmt.Sprint(v...), "\nError stack:\n", string(debug.Stack()))
}

// Tracef logs 详细运行跟踪日志，可能影响程序性能，所以生产环境不配置输出，仅在开发测试环境使用
func Tracef(format string, params ...interface{}) {
	seelog.Tracef(format, params...)
}

// Debugf formats logs
func Debugf(format string, params ...interface{}) {
	seelog.Debugf(format, params...)
}

// Infof formats logs
func Infof(format string, params ...interface{}) {
	seelog.Infof(format, params...)
}

// Warnf formats logs
func Warnf(format string, params ...interface{}) {
	seelog.Warnf(format, params...)
}

// Errorf formats logs
func Errorf(format string, params ...interface{}) {
	seelog.Errorf(format, params...)
	// seelog.Error(fmt.Sprintf(format, params...), "\nStack:\n", string(debug.Stack()))
	//params = append(params, string(debug.Stack()))
	//seelog.Errorf(format+"\nError stack:\n%s", params...)
}

var logLevelToString = map[log.LogLevel]string{
	log.TraceLvl:    "T",
	log.DebugLvl:    "D",
	log.InfoLvl:     "I",
	log.WarnLvl:     "W",
	log.ErrorLvl:    "E",
	log.CriticalLvl: "C",
	log.Off:         "_",
}

func logLevelFormatter(params string) log.FormatterFunc {
	return func(message string, level log.LogLevel,
		context log.LogContextInterface) interface{} {
		levelStr, ok := logLevelToString[level]
		if !ok {
			return "!"
		}
		return levelStr
	}
}
