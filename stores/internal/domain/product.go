package domain

import (
	"eda-in-golang/internal/ddd"

	"github.com/stackus/errors"
)

var (
	ErrProductNameIsBlank     = errors.Wrap(errors.ErrBadRequest, "the product name cannot be blank")
	ErrProductPriceIsNegative = errors.Wrap(errors.ErrBadRequest, "the product price cannot be negative")
)

type Product struct {
	ddd.Aggregate
	StoreID     string
	Name        string
	Description string
	SKU         string
	Price       float64
}

func CreateProduct(id, storeID, name, description, sku string, price float64) (*Product, error) {
	if name == "" {
		return nil, ErrProductNameIsBlank
	}

	if price < 0 {
		return nil, ErrProductPriceIsNegative
	}

	product := &Product{
		Aggregate: &ddd.AggregateBase{
			ID: id,
		},
		StoreID:     storeID,
		Name:        name,
		Description: description,
		SKU:         sku,
		Price:       price,
	}

	product.AddEvent(&ProductAdded{
		Product: product,
	})

	return product, nil
}

func (p *Product) RemoveProduct() {
	p.AddEvent(&ProductRemoved{
		Product: p,
	})
}
