package common_producers_event_gateway

type TrackUserLoginEvent struct {
	UserID uint32 `json:"user_id"`
}

func (e *TrackUserLoginEvent) Name() string {
	return "TrackUserLoginEvent"
}

func (e *TrackUserLoginEvent) Data() interface{} {
	return e
}
