package candle_responses

type GetCandlesResponse struct {
	OpenTime         uint64  `json:"open_time"`
	Open             float64 `json:"open"`
	High             float64 `json:"high"`
	Low              float64 `json:"low"`
	Close            float64 `json:"close"`
	Volume           float64 `json:"volume"`
	CloseTime        uint64  `json:"close_time"`
	QuoteAssetVolume float64 `json:"quote_asset_volume"`
	NumberOfTrades   uint64  `json:"number_of_trades"`
}
