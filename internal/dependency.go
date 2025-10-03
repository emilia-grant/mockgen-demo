package internal

import (
	"context"
	"errors"
	"fmt"
	"net/http"
)

var (
	ErrNilClient = errors.New("nil client")
	ErrCtx       = errors.New("error in context")
)

type SomeDependency struct {
}

func NewSomeComponent() *SomeDependency {
	return &SomeDependency{}
}

func (m *SomeDependency) DoNetworkThing(ctx context.Context, h *http.Client) error {
	if h == nil {
		return ErrNilClient
	}
	if ctx.Err() != nil {
		return fmt.Errorf("%w: %w", ErrCtx, ctx.Err())
	}
	return nil
}

func (m *SomeDependency) DoPrivateThing() error {
	return nil
}
