package commands

import (
	"context"

	"eda-in-golang/internal/ddd"
	"eda-in-golang/stores/internal/domain"
)

type RemoveProduct struct {
	ID string
}

type RemoveProductHandler struct {
	stores          domain.StoreRepository
	products        domain.ProductRepository
	domainPublisher ddd.EventPublisher[ddd.AggregateEvent]
}

func NewRemoveProductHandler(stores domain.StoreRepository, products domain.ProductRepository, domainPublisher ddd.EventPublisher[ddd.AggregateEvent]) RemoveProductHandler {
	return RemoveProductHandler{
		stores:          stores,
		products:        products,
		domainPublisher: domainPublisher,
	}
}

func (h RemoveProductHandler) RemoveProduct(ctx context.Context, cmd RemoveProduct) error {
	product, err := h.products.FindProduct(ctx, cmd.ID)
	if err != nil {
		return err
	}

	if err = h.products.RemoveProduct(ctx, cmd.ID); err != nil {
		return err
	}

	product.RemoveProduct()

	// publish domain events
	if err = h.domainPublisher.Publish(ctx, product.Events()...); err != nil {
		return err
	}

	return nil
}
