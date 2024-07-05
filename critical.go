package yikes

import (
	"log/slog"
)

const (
	// LevelCritical is a [log/slog.Level] that indicates a critical error.
	// A critical error is more serious than an error, and warrants a
	// noisier notification (email, message in chat, etc.) than an error on
	// every occurrence. It should be used incredibly sparingly, and only
	// in the most catastrophic cases; things that indicate security
	// problems, a loss of revenue, data corruption, or other things that
	// require manual intervention from an operator.
	LevelCritical slog.Level = 16
)
