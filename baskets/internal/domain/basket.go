package domain

import (
	"eda-in-golang/internal/ddd"
	"sort"

	"github.com/stackus/errors"
)

var (
	ErrBasketHasNoItems         = errors.Wrap(errors.ErrBadRequest, "the basket has no items")
	ErrBasketCannotBeModified   = errors.Wrap(errors.ErrBadRequest, "the basket cannot be modified")
	ErrBasketCannotBeCancelled  = errors.Wrap(errors.ErrBadRequest, "the basket cannot be cancelled")
	ErrQuantityCannotBeNegative = errors.Wrap(errors.ErrBadRequest, "the item quantity cannot be negative")
	ErrBasketIDCannotBeBlank    = errors.Wrap(errors.ErrBadRequest, "the basket id cannot be blank")
	ErrPaymentIDCannotBeBlank   = errors.Wrap(errors.ErrBadRequest, "the payment id cannot be blank")
	ErrCustomerIDCannotBeBlank  = errors.Wrap(errors.ErrBadRequest, "the customer id cannot be blank")
)

type BasketStatus string

const (
	BasketIsUnknown    BasketStatus = ""
	BasketIsOpen       BasketStatus = "open"
	BasketIsCancelled  BasketStatus = "cancelled"
	BasketIsCheckedOut BasketStatus = "checked_out"
)

func (s BasketStatus) String() string {
	switch s {
	case BasketIsOpen, BasketIsCancelled, BasketIsCheckedOut:
		return string(s)
	default:
		return ""
	}
}

type Basket struct {
	ddd.Aggregate
	CustomerID string
	PaymentID  string
	Items      []Item
	Status     BasketStatus
}

func StartBasket(id, customerID string) (*Basket, error) {
	if id == "" {
		return nil, ErrBasketIDCannotBeBlank
	}

	if customerID == "" {
		return nil, ErrCustomerIDCannotBeBlank
	}

	basket := &Basket{
		Aggregate: &ddd.AggregateBase{
			ID: id,
		},
		CustomerID: customerID,
		Status:     BasketIsOpen,
		Items:      []Item{},
	}

	basket.AddEvent(&BasketStarted{
		Basket: basket,
	})

	return basket, nil
}

func (b Basket) IsCancellable() bool {
	return b.Status == BasketIsOpen
}

func (b Basket) IsOpen() bool {
	return b.Status == BasketIsOpen
}

func (b *Basket) Cancel() error {
	if !b.IsCancellable() {
		return ErrBasketCannotBeCancelled
	}

	b.Status = BasketIsCancelled
	b.Items = []Item{}

	b.AddEvent(&BasketCanceled{
		Basket: b,
	})

	return nil
}

func (b *Basket) Checkout(paymentID string) error {
	if !b.IsOpen() {
		return ErrBasketCannotBeModified
	}

	if len(b.Items) == 0 {
		return ErrBasketHasNoItems
	}

	if paymentID == "" {
		return ErrPaymentIDCannotBeBlank
	}

	b.PaymentID = paymentID
	b.Status = BasketIsCheckedOut

	b.AddEvent(&BasketCheckedOut{
		Basket: b,
	})

	return nil
}

func (b *Basket) AddItem(store *Store, product *Product, quantity int) error {
	if !b.IsOpen() {
		return ErrBasketCannotBeModified
	}

	if quantity < 0 {
		return ErrQuantityCannotBeNegative
	}

	for i, item := range b.Items {
		if item.ProductID == product.ID && item.StoreID == product.StoreID {
			b.Items[i].Quantity += quantity
			return nil
		}
	}

	b.Items = append(b.Items, Item{
		StoreID:      store.ID,
		ProductID:    product.ID,
		StoreName:    store.Name,
		ProductName:  product.Name,
		ProductPrice: product.Price,
		Quantity:     quantity,
	})

	sort.Slice(b.Items, func(i, j int) bool {
		return b.Items[i].StoreName <= b.Items[j].StoreName && b.Items[i].ProductName < b.Items[j].ProductName
	})

	b.AddEvent(&BasketItemAdded{
		Basket: b,
	})

	return nil
}

func (b *Basket) RemoveItem(product *Product, quantity int) error {
	if !b.IsOpen() {
		return ErrBasketCannotBeModified
	}

	if quantity < 0 {
		return ErrQuantityCannotBeNegative
	}

	for i, item := range b.Items {
		if item.ProductID == product.ID && item.StoreID == product.StoreID {
			b.Items[i].Quantity -= quantity

			if b.Items[i].Quantity < 1 {
				b.Items = append(b.Items[:i], b.Items[i+1:]...)
			}
			return nil
		}
	}

	b.AddEvent(&BasketItemRemoved{
		Basket: b,
	})

	return nil
}
