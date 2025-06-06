package common_producers_event

type CreatedCandleEvent struct {
	PairID uint64 `json:"id"`
}

func (c CreatedCandleEvent) Name() string {
	return "CreatedCandleEvent"
}

func (c CreatedCandleEvent) Data() interface{} {
	return c
}
