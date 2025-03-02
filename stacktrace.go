package yikes

import (
	"runtime"
	"strconv"
	"strings"
)

func getStacktrace() string {
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

	var out strings.Builder
	for frame, more := frames.Next(); more; frame, more = frames.Next() {
		out.WriteString(frame.Function)
		out.WriteString("\n\t")
		out.WriteString(frame.File)
		out.WriteString(":")
		out.WriteString(strconv.FormatInt(int64(frame.Line), 10))
		out.WriteString("\n")
	}

	return strings.TrimSuffix(out.String(), "\n")
}
