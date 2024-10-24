package application

import (
	"context"
	"eda-in-golang/internal/ddd"
)

type DomainEventHandlers interface {
	OnOrderCreated(context.Context, ddd.Event) error
	OnOrderReadied(context.Context, ddd.Event) error
	OnOrderCanceled(context.Context, ddd.Event) error
	OnOrderCompleted(context.Context, ddd.Event) error
}

type ignoreUnimplementedDomainEvents struct{}

var _ DomainEventHandlers = (*ignoreUnimplementedDomainEvents)(nil)

func (ignoreUnimplementedDomainEvents) OnOrderCreated(ctx context.Context, event ddd.Event) error {
	return nil
}

func (ignoreUnimplementedDomainEvents) OnOrderReadied(ctx context.Context, event ddd.Event) error {
	return nil
}

func (ignoreUnimplementedDomainEvents) OnOrderCanceled(ctx context.Context, event ddd.Event) error {
	return nil
}

func (ignoreUnimplementedDomainEvents) OnOrderCompleted(ctx context.Context, event ddd.Event) error {
	return nil
}
