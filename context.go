package yikes

import (
	"context"
	"log/slog"
)

type ctxKey struct{}

var (
	reporterKey = ctxKey{}
)

// In injects a [Reporter] into a [context.Context] and returns the modified
// Context.
func In(ctx context.Context, reporter Reporter) context.Context {
	return context.WithValue(ctx, reporterKey, reporter)
}

// From extracts a [Reporter] from a [context.Context] that was created by
// [In]. If the Context wasn't created by [In], a no-op Reporter will be
// returned that doesn't do anything when its methods are called.
func From(ctx context.Context) Reporter {
	val := ctx.Value(reporterKey)
	if val != nil {
		if reporter, ok := val.(Reporter); ok {
			return reporter
		}
	}
	return Reporter{
		Logger: slog.New(noopHandler{}),
	}
}

// Set is a helper function that updates the [log/slog.Logger] on the
// [Reporter] in the passed [context.Context] to have the passed key value
// attributes associated with it.
func Set(ctx context.Context, args ...any) context.Context {
	return In(ctx, Reporter{Logger: From(ctx).Logger.With(args...)})
}
