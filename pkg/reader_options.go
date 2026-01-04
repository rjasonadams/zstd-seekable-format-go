package seekable

import (
	"github.com/SaveTheRbtz/zstd-seekable-format-go/pkg/env"
)

type rOption func(*readerImpl) error

// WithRLogger sets the logger for the reader.
func WithRLogger(l Logger) rOption {
	return func(r *readerImpl) error { r.logger = l; return nil }
}

func WithREnvironment(e env.REnvironment) rOption {
	return func(r *readerImpl) error { r.env = e; return nil }
}
