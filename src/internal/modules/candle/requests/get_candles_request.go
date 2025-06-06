package candle_requests

type GetCandlesRequest struct {
	Symbol    string `json:"symbol"`
	Interval  string `json:"interval"`
	StartTime uint64 `json:"start_time"`
	EndTime   uint64 `json:"end_time"`
}
