package log

import (
	"github.com/cihub/seelog"
)

var defaultConfig = `
<seelog minlevel="trace">
	<outputs formatid="fmt">
	    <console />
	</outputs>
	<formats>
		<format id="fmt" format="%L [%Date(2006-01-02 15:04:05.000000000)] %Msg%n"/>
	</formats>
</seelog>`

var defaultRollingFileConfig = `
<seelog minlevel="debug">
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
	err := seelog.RegisterCustomFormatter("L", logLevelFormatter)
	if err != nil {
		Errorf("log init failed: %v", err)
	}
	logger, _ := seelog.LoggerFromConfigAsBytes([]byte(defaultConfig))
	replace(logger)
}

// Config replace logger from config string
func Config(config string) {
	logger, err := seelog.LoggerFromConfigAsBytes([]byte(config))
	if err != nil {
		Errorf("log config error: ", err)
		return
	}
	replace(logger)
}

// Load replace logger from config file
func Load(conf string) {
	logger, err := seelog.LoggerFromConfigAsFile(conf)
	if err != nil {
		Errorf("log config load error: ", err)
		return
	}
	replace(logger)
}

// DefaultRollingFileConfig returns config for rolling file
func DefaultRollingFileConfig() string {
	return defaultRollingFileConfig
}

// replace logger
func replace(logger seelog.LoggerInterface) {
	seelog.ReplaceLogger(logger)
}

// Flush immediately processes all currently queued logs.
func Flush() {
	seelog.Flush()
}

// Trace logs
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

// Tracef logs
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

var logLevelToString = map[seelog.LogLevel]string{
	seelog.TraceLvl:    "T",
	seelog.DebugLvl:    "D",
	seelog.InfoLvl:     "I",
	seelog.WarnLvl:     "W",
	seelog.ErrorLvl:    "E",
	seelog.CriticalLvl: "C",
	seelog.Off:         "_",
}

func logLevelFormatter(params string) seelog.FormatterFunc {
	return func(message string, level seelog.LogLevel,
		context seelog.LogContextInterface) interface{} {
		levelStr, ok := logLevelToString[level]
		if !ok {
			return "!"
		}
		return levelStr
	}
}
