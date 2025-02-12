package event

import (
	"GourseAPI/internal/valueobjects"
	"context"
	"time"
)

// Bus defines the expected behaviour from an event bus.
type Bus interface {
	// Publish is the method used to publish new events.
	Publish(context.Context, []Event) error
	// Subscribe is the method used to subscribe new event handlers.
	Subscribe(Type, Handler)
}

//go:generate mockery --case=snake --outpkg=eventmocks --output=eventmocks --name=Bus

// Handler defines the expected behaviour from an event handler.
type Handler interface {
	Handle(context.Context, Event) error
}

// Type represents a domain event type.
type Type string

// Event represents a domain command.
type Event interface {
	ID() string
	AggregateID() string
	OccurredOn() time.Time
	Type() Type
}

type BaseEvent struct {
	eventID     valueobjects.UUID
	aggregateID string
	occurredOn  time.Time
}

func NewBaseEvent(aggregateID string) BaseEvent {
	return BaseEvent{
		eventID:     valueobjects.GenerateUUID(),
		aggregateID: aggregateID,
		occurredOn:  time.Now(),
	}
}

func (b BaseEvent) ID() string {
	return b.eventID.String()
}

func (b BaseEvent) OccurredOn() time.Time {
	return b.occurredOn
}

func (b BaseEvent) AggregateID() string {
	return b.aggregateID
}
