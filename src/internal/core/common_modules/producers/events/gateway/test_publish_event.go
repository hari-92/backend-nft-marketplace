package common_producers_event_gateway

type TestPublishEvent struct {
	UserID  uint32 `json:"user_id"`
	Message string `json:"message"`
}

func (t TestPublishEvent) Name() string {
	return "TestPublishEvent"
}

func (t TestPublishEvent) Data() interface{} {
	return &t
}
