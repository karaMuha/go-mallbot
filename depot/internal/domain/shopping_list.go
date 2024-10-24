package domain

import (
	"eda-in-golang/internal/ddd"

	"github.com/stackus/errors"
)

var (
	ErrShoppingCannotBeCancelled = errors.Wrap(errors.ErrBadRequest, "the shopping list cannot be cancelled")
)

type ShoppingListStatus string

const (
	ShoppingListIsUnknown   ShoppingListStatus = ""
	ShoppingListIsAvailable ShoppingListStatus = "available"
	ShoppingListIsAssigned  ShoppingListStatus = "assigned"
	ShoppingListIsActive    ShoppingListStatus = "active"
	ShoppingListIsCompleted ShoppingListStatus = "completed"
	ShoppingListIsCancelled ShoppingListStatus = "cancelled"
)

func (s ShoppingListStatus) String() string {
	switch s {
	case ShoppingListIsAvailable, ShoppingListIsAssigned, ShoppingListIsActive, ShoppingListIsCompleted, ShoppingListIsCancelled:
		return string(s)
	default:
		return ""
	}
}

func ToShoppingListStatus(status string) ShoppingListStatus {
	switch status {
	case ShoppingListIsAvailable.String():
		return ShoppingListIsAvailable
	case ShoppingListIsAssigned.String():
		return ShoppingListIsAssigned
	case ShoppingListIsActive.String():
		return ShoppingListIsActive
	case ShoppingListIsCompleted.String():
		return ShoppingListIsCompleted
	case ShoppingListIsCancelled.String():
		return ShoppingListIsCancelled
	default:
		return ShoppingListIsUnknown
	}
}

type ShoppingList struct {
	ddd.Aggregate
	OrderID       string
	Stops         Stops
	AssignedBotID string
	Status        ShoppingListStatus
}

func CreateShopping(id, orderID string) *ShoppingList {
	shoppingList := &ShoppingList{
		Aggregate: &ddd.AggregateBase{
			ID: id,
		},
		OrderID: orderID,
		Status:  ShoppingListIsAvailable,
		Stops:   make(Stops),
	}

	shoppingList.AddEvent(&ShoppingListCreated{
		ShoppingList: shoppingList,
	})

	return shoppingList
}

func (sl *ShoppingList) AddItem(store *Store, product *Product, quantity int) error {
	if _, exists := sl.Stops[store.ID]; !exists {
		sl.Stops[store.ID] = &Stop{
			StoreName:     store.Name,
			StoreLocation: store.Location,
			Items:         make(Items),
		}
	}

	return sl.Stops[store.ID].AddItem(product, quantity)
}

func (sl *ShoppingList) Cancel() error {
	// validate status

	sl.Status = ShoppingListIsCancelled

	sl.AddEvent(&ShoppingListCanceled{
		ShoppingList: sl,
	})

	return nil
}

func (sl *ShoppingList) Assign(id string) error {
	// validate status

	sl.AssignedBotID = id
	sl.Status = ShoppingListIsAssigned

	sl.AddEvent(&ShoppingListAssigned{
		ShoppingList: sl,
	})

	return nil
}

func (sl *ShoppingList) Complete() error {
	// validate status

	sl.Status = ShoppingListIsCompleted

	sl.AddEvent(&ShoppingListCompleted{
		ShoppingList: sl,
	})

	return nil
}
