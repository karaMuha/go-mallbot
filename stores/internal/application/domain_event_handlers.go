package application

import (
	"context"
	"eda-in-golang/internal/ddd"
)

type DomainEventHandlers interface {
	OnStoreCreated(context.Context, ddd.Event) error
	OnStoreParticipationEnabled(context.Context, ddd.Event) error
	OnStoreParticipationDisabled(context.Context, ddd.Event) error
	OnProductAdded(context.Context, ddd.Event) error
	OnProductRemoved(context.Context, ddd.Event) error
}

type ignoreUnimplementedDomainEvents struct{}

var _ DomainEventHandlers = (*ignoreUnimplementedDomainEvents)(nil)

func (ignoreUnimplementedDomainEvents) OnStoreCreated(ctx context.Context, event ddd.Event) error {
	return nil
}

func (ignoreUnimplementedDomainEvents) OnStoreParticipationEnabled(ctx context.Context, event ddd.Event) error {
	return nil
}

func (ignoreUnimplementedDomainEvents) OnStoreParticipationDisabled(ctx context.Context, event ddd.Event) error {
	return nil
}

func (ignoreUnimplementedDomainEvents) OnProductAdded(ctx context.Context, event ddd.Event) error {
	return nil
}

func (ignoreUnimplementedDomainEvents) OnProductRemoved(ctx context.Context, event ddd.Event) error {
	return nil
}
