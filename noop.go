package yikes

import (
	"context"
	"log/slog"
)

// noopHandler provides a [*log/slog.Logger] handler that discards any logs
// passed to it.
type noopHandler struct{}

// Enabled always returns false, indicating the handler doesn't handle any
// [log/slog.Level].
func (noopHandler) Enabled(_ context.Context, _ slog.Level) bool {
	return false
}

// Handle always returns nil, and is a no-op.
func (noopHandler) Handle(_ context.Context, _ slog.Record) error {
	return nil
}

// WithAttrs always ignores the passed [log/slog.Attrs], returning the receiver
// unmodified.
func (n noopHandler) WithAttrs(_ []slog.Attr) slog.Handler {
	return n
}

// WithGroup always ignores the group name, returning the receiver unmodified.
func (n noopHandler) WithGroup(_ string) slog.Handler {
	return n
}
