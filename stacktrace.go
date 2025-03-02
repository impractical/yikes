package yikes

import (
	"log/slog"
	"runtime"
	"strconv"
)

func getStacktrace() slog.Attr {
	buf := make([]uintptr, 1024)
	// 0 is runtime.Callers, 1 is getStacktrace, 2 is either report or
	// topLevelReport, 3 is the exported function of yikes that was called,
	// 4 is the caller.
	offset := 4
	var numFrames int
	for {
		numFrames = runtime.Callers(offset, buf)
		if numFrames < len(buf) {
			break
		}
		buf = make([]uintptr, len(buf)*2)
	}

	frames := runtime.CallersFrames(buf[:numFrames])

	var result []slog.Attr
	var index int
	for frame, more := frames.Next(); more; frame, more = frames.Next() {
		result = append(result, slog.Group(strconv.Itoa(index),
			slog.String("function", frame.Function),
			slog.String("file", frame.File),
			slog.Int("line", frame.Line),
		))
	}

	return slog.Any("stacktrace", slog.GroupValue(result...))
}
