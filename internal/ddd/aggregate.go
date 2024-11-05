package ddd

const (
	AggregateNameKey    = "aggregate-name"
	AggregateIDKey      = "aggregate-id"
	AggregateVersionKey = "aggregate-version"
)

type AggregateEvent interface {
	Event
	AggregateName() string
	AggregateID() string
	AggregateVersion() int
}

type Aggregate struct {
	Entity
	events []AggregateEvent
}

type aggregateEvent struct {
	event
}

type AggregateBase struct {
	ID     string
	events []Event
}

func NewAggregate(id, name string) Aggregate {
	return Aggregate{
		Entity: NewEntity(id, name),
		events: make([]AggregateEvent, 0),
	}
}

func (a Aggregate) AggregateName() string {
	return a.name
}

func (a Aggregate) Events() []AggregateEvent {
	return a.events
}

func (a Aggregate) ClearEvents() {
	a.events = []AggregateEvent{}
}

func (a *Aggregate) AddEvent(name string, payload EventPayload, options ...EventOption) {
	options = append(
		options,
		Metadata{
			AggregateNameKey: a.name,
			AggregateIDKey:   a.id,
		},
	)
	a.events = append(
		a.events,
		aggregateEvent{
			event: newEvent(name, payload, options...),
		},
	)
}

func (a *Aggregate) setEvents(events []AggregateEvent) {
	a.events = events
}

func (e aggregateEvent) AggregateName() string {
	return e.metadata.Get(AggregateNameKey).(string)
}

func (e aggregateEvent) AggregateID() string {
	return e.metadata.Get(AggregateIDKey).(string)
}

func (e aggregateEvent) AggregateVersion() int {
	return e.metadata.Get(AggregateVersionKey).(int)
}
