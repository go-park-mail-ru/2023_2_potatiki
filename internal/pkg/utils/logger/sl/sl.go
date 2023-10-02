package sl

import (
	"log/slog"
	"runtime"
	"strings"
)

func Err(err error) slog.Attr {
	return slog.Attr{
		Key:   "error",
		Value: slog.StringValue(err.Error()),
	}
}

func GFN() string {
	// GFN - GetFunctionName
	// return function name for generate true error message
	pc := make([]uintptr, 15)
	n := runtime.Callers(2, pc)
	frames := runtime.CallersFrames(pc[:n])
	frame, _ := frames.Next()
	values := strings.Split(frame.Function, "/")
	return values[len(values)-1]
}
