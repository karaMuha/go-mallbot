package ddd

import (
	"context"
	"sync"
)

type EventHandler[T Event] interface {
	HandleEvent(ctx context.Context, event T) error
}

type EventHandlerFunc[T Event] func(ctx context.Context, event T) error

type EventSubscriber[T Event] interface {
	Subscribe(name string, handler EventHandler[T])
}

type EventPublisher[T Event] interface {
	Publish(ctx context.Context, events ...T) error
}

type EventDispatcher[T Event] struct {
	handlers map[string][]EventHandler[T]
	mu       sync.Mutex
}

var _ interface {
	EventSubscriber[Event]
	EventPublisher[Event]
} = (*EventDispatcher[Event])(nil)

func NewEventDispatcher[T Event]() *EventDispatcher[T] {
	return &EventDispatcher[T]{
		handlers: make(map[string][]EventHandler[T]),
	}
}

func (h *EventDispatcher[T]) Subscribe(name string, handler EventHandler[T]) {
	h.mu.Lock()
	defer h.mu.Unlock()

	h.handlers[name] = append(h.handlers[name], handler)
}

func (h *EventDispatcher[T]) Publish(ctx context.Context, events ...T) error {
	for _, event := range events {
		for _, handler := range h.handlers[event.EventName()] {
			err := handler.HandleEvent(ctx, event)
			if err != nil {
				return err
			}
		}
	}
	return nil
}
