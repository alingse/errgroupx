package errgroupx

import (
	"context"

	"golang.org/x/sync/errgroup"
)

type GroupX struct {
	*errgroup.Group
}

func WithContext(ctx context.Context) (*GroupX, context.Context) {
	group, ctx := errgroup.WithContext(ctx)
	return &GroupX{Group: group}, ctx
}

func (x *GroupX) Wait() error {
	return x.Group.Wait()
}

func (x *GroupX) Go(f func() error) {
	x.Group.Go(f)
}
