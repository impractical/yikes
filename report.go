package yikes

import (
	"context"
	"log/slog"
)

// Reporter is a collection of state required to report errors to the operators
// of a program.
type Reporter struct {
	// Logger is a [log/slog.Logger] that is already configured and ready
	// to be used to log errors.
	//
	// If you plan to use [Reporter.Critical], the Logger should already be
	// configured to map the string value of [LevelCritical], or it'll show
	// up as "ERROR+8".
	Logger *slog.Logger
}

// Report calls [log/slog.Logger.Log] on the [log/slog.Logger] associated with
// the [Reporter]. The passed [error] will be logged under the "error" key
// automatically. It then wraps the passed [error] such that [AlreadyReported]
// returns true for the [log/slog.Level] and all lower levels.
//
// If [AlreadyReported] returns true for the [error] and [log/slog.Level]
// passed into Report, Report returns the [error] as it was passed in and takes
// no further actions.
func (reporter Reporter) Report(ctx context.Context, level slog.Level, message string, err error, args ...any) error { //nolint:revive // we're largely just copying slog.Log at this point, this is as succinct as it gets
	if AlreadyReported(err, level) {
		return err
	}
	args = append([]any{"error", err}, args...)
	reporter.Logger.Log(ctx, level, message, args...)
	return reportedError{
		level: level,
		error: err,
	}
}

// Error calls [Reporter.Report] using the [log/slog.LevelError] level.
func (reporter Reporter) Error(ctx context.Context, message string, err error, args ...any) error {
	return reporter.Report(ctx, slog.LevelError, message, err, args...)
}

// Warn calls [Reporter.Report] using the [log/slog.LevelWarn] level.
func (reporter Reporter) Warn(ctx context.Context, message string, err error, args ...any) error {
	return reporter.Report(ctx, slog.LevelWarn, message, err, args...)
}

// Critical calls [Reporter.Report] using the [LevelCritical] level.
func (reporter Reporter) Critical(ctx context.Context, message string, err error, args ...any) error {
	return reporter.Report(ctx, LevelCritical, message, err, args...)
}
