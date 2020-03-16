package encoder

import (
	"runtime"

	"go.uber.org/zap/zapcore"
)

func ProxyCallerEncoder(caller zapcore.EntryCaller,
	enc zapcore.PrimitiveArrayEncoder) {
	c := zapcore.NewEntryCaller(runtime.Caller(6))
	enc.AppendString(c.TrimmedPath())
}
