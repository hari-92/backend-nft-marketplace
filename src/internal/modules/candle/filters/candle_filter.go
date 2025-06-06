package candle_filters

import "time"

type CandlesFilter struct {
	StartTime time.Time `json:"start_time" form:"start_time"`
	EndTime   time.Time `json:"end_time" form:"end_time"`
	PairID    uint64    `json:"pair_id" form:"pair_id"`
}
