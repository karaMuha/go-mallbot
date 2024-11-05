package domain

const (
	ProductAddedEvent          = "stores.ProductAdded"
	ProductRebrandedEvent      = "stores.ProductRebranded"
	ProductPriceIncreasedEvent = "stores.ProductPriceIncreased"
	ProductPriceDecreasedEvent = "stores.ProductPriceDecreased"
	ProductRemovedEvent        = "stores.ProductRemoved"
)

type ProductAdded struct {
	Product *Product
}

func (ProductAdded) EventName() string {
	return "stores.ProductAdded"
}

type ProductRemoved struct {
	Product *Product
}

func (ProductRemoved) EventName() string {
	return "stores.ProductRemoved"
}
