package common_producers_event_gateway

type UserCreatedEvent struct {
	UserID uint32 `json:"user_id"`
}

func (e *UserCreatedEvent) Name() string {
	return "UserCreatedEvent"
}

func (e *UserCreatedEvent) Data() interface{} {
	return e
}
