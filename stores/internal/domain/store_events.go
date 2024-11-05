package domain

const (
	StoreCreatedEvent               = "stores.StoreCreated"
	StoreParticipationEnabledEvent  = "stores.StoreParticipationEnabled"
	StoreParticipationDisabledEvent = "stores.StoreParticipationDisabled"
	StoreRebrandedEvent             = "stores.StoreRebranded"
)

type StoreCreated struct {
	Store *Store
}

func (StoreCreated) EventName() string {
	return "stores.StoreCreated"
}

type StoreParticipationEnabled struct {
	Store *Store
}

func (StoreParticipationEnabled) EventName() string {
	return "stores.StoreParticipationEnabled"
}

type StoreParticipationDisabled struct {
	Store *Store
}

func (StoreParticipationDisabled) EventName() string {
	return "stores.StoreParticipationDisabled"
}
