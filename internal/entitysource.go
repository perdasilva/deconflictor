package internal

import (
	"context"

	"github.com/operator-framework/deppy/pkg/deppy"
	"github.com/operator-framework/deppy/pkg/deppy/input"
)

// entity sources are not necessary - so creating a dummy one
var _ input.EntitySource = &DummyEntitySource{}

type DummyEntitySource struct {
}

func (d DummyEntitySource) Get(ctx context.Context, id deppy.Identifier) *input.Entity {
	//TODO implement me
	panic("implement me")
}

func (d DummyEntitySource) Filter(ctx context.Context, filter input.Predicate) (input.EntityList, error) {
	//TODO implement me
	panic("implement me")
}

func (d DummyEntitySource) GroupBy(ctx context.Context, fn input.GroupByFunction) (input.EntityListMap, error) {
	//TODO implement me
	panic("implement me")
}

func (d DummyEntitySource) Iterate(ctx context.Context, fn input.IteratorFunction) error {
	//TODO implement me
	panic("implement me")
}
