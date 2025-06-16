package qa

import (
	"context"
	"testing"

	"github.com/outofforest/logger"
)

// NewContext creates new context for tests.
func NewContext(t *testing.T) context.Context {
	return logger.WithLogger(t.Context(), logger.New(logger.DefaultConfig))
}
