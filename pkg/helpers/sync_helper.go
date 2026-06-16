package helpers

import (
	"context"

	"golang.org/x/sync/errgroup"
)

func NewErrGroup(ctx context.Context, workerCount int) (*errgroup.Group, context.Context) {
	group, ctx := errgroup.WithContext(ctx)
	group.SetLimit(workerCount)
	return group, ctx
}
