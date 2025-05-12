package candle_stick_controllers

import (
	candleStickServices "gitlab.com/hari-92/nft-market-server/internal/modules/candle_stick/services"
)

type CandleStickController struct {
	candleStickService candleStickServices.CandleStickService
}

func NewCandleStickController(candleStickService candleStickServices.CandleStickService) *CandleStickController {
	return &CandleStickController{
		candleStickService: candleStickService,
	}
}
