package qa

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/outofforest/parallel"
)

// NewGroup creates new group for test.
func NewGroup(ctx context.Context, t *testing.T) *parallel.Group {
	group := parallel.NewGroup(ctx)
	t.Cleanup(func() {
		group.Exit(nil)
		require.NoError(t, group.Wait())
	})
	return group
}
