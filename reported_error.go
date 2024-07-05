package yikes

import (
	"errors"
	"log/slog"
)

var (
	_ error = reportedError{}
)

type reportedError struct {
	level slog.Level
	error
}

// Unwrap returns the underlying [error] that was reported.
func (reported reportedError) Unwrap() error {
	return reported.error
}

// AlreadyReported returns true if the [error] has already been reported at the
// same or higher [log/slog.Level]. If the [error] has been reported at a lower
// [log/slog.Level], it will be reported again at the higher level.
func AlreadyReported(err error, level slog.Level) bool {
	var reported reportedError
	return errors.As(err, &reported) && reported.level >= level
}
