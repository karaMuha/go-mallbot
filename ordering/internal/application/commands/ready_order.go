package commands

import (
	"context"

	"eda-in-golang/internal/ddd"
	"eda-in-golang/ordering/internal/domain"
)

type ReadyOrder struct {
	ID string
}

type ReadyOrderHandler struct {
	orders          domain.OrderRepository
	invoices        domain.InvoiceRepository
	domainPublisher ddd.EventPublisher
}

func NewReadyOrderHandler(orders domain.OrderRepository, invoices domain.InvoiceRepository,
	domainPublisher ddd.EventPublisher,
) ReadyOrderHandler {
	return ReadyOrderHandler{
		orders:          orders,
		invoices:        invoices,
		domainPublisher: domainPublisher,
	}
}

func (h ReadyOrderHandler) ReadyOrder(ctx context.Context, cmd ReadyOrder) error {
	order, err := h.orders.Find(ctx, cmd.ID)
	if err != nil {
		return err
	}

	if err = order.Ready(); err != nil {
		return nil
	}

	if err = h.orders.Update(ctx, order); err != nil {
		return err
	}

	if err = h.orders.Update(ctx, order); err != nil {
		return err
	}

	if err = h.domainPublisher.Publish(ctx, order.GetEvents()...); err != nil {
		return err
	}

	return nil
}
