package kafka

import "github.com/golibs-starter/golib/web/event"

// Event represents ...
type Event[T any] struct {
	*event.AbstractEvent
	PayloadData T `json:"payload"`
}

// Payload return payload of event
func (e Event[T]) Payload() interface{} {
	return e.PayloadData
}

// String() convert event to string
func (e Event[T]) String() string {
	return e.ToString(e)
}
