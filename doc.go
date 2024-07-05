// Package yikes provides functionality for reporting errors to operators.
//
// It is based on the design explored in [this post], specifically the error
// reporting for operators bit. The idea is that operators want errors reported
// wherever they occur in the call stack, not from some middleware that
// requires the error to be aware of its own stack trace and any context
// surrounding the error than operator may be interested in.
//
// It provides a [Reporter] type that can be used to report errors. It uses the
// [log/slog.Logger] type to report errors, but the reporting doesn't have to
// be to a log line; the Logger could be configured, for example, to report to
// an error reporting service.
//
// The error return from the [Reporter.Report] method (and its aliases,
// [Reporter.Error], [Reporter.Warn], and [Reporter.Critical]) is the same
// error but wrapped in a type that indicates to yikes that the error has
// already been reported and should not be reported again, unless it's being
// reported at a higher level. This is done so callers of helper functions
// within your application don't need to worry about whether the helper
// reported the error or not; they can just call [Reporter.Report] and have it
// do the deduplication itself. It will be rereported at a higher level so any
// automation that triggers at higher levels will still trigger, even if a
// function deeper in the call stack didn't have the context to understand the
// error should be reported at a higher level. If the application, for some
// reason, needs to know whether an error was reported or not,
// [AlreadyReported] can provide that information.
//
// The intended use of this package is, whenever encountering an error an
// operator would be interested in, to replace this:
//
//	if err != nil {
//		return err
//	}
//
// with this:
//
//	if err != nil {
//		return yikes.From(ctx).Error(ctx, "some error occurred", err)
//	}
//
// optionally passing any context the operator would be interested in as
// arguments to [Reporter.Error].
//
// [this post]: https://paddy.carvers.com/posts/go-errors
package yikes
