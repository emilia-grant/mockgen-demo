package internal

import (
	"context"
	"errors"
	"net/http"
)

var (
	ErrUnrecoverable = errors.New("time to panic")
)

// Service uses has a dependency that it uses to DoNetworkThing
type Service struct {
	mydep SomeDependencyContract
	h     *http.Client
}

// Since Service and dependency are in the same module, there's no difference to providing the
// interface here or in dependency.go but the interfaces are usually written by the module consumer.

//go:generate mockgen -source=service.go -destination ./mock/internal_mock.go
type SomeDependencyContract interface {
	DoNetworkThing(ctx context.Context, h *http.Client) error
}

func NewService(dep SomeDependencyContract, h *http.Client) Service {
	return Service{
		mydep: dep,
		h:     h,
	}
}

func (s Service) DoAThing(ctx context.Context) error {
	err := s.mydep.DoNetworkThing(ctx, s.h)
	if err == nil {
		return nil
	}
	// Simulated "something is VERY wrong" moment
	if errors.Is(err, ErrNilClient) {
		return ErrUnrecoverable
	}
	return err
}
