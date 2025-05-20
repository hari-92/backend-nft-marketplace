package candle_controllers

import (
	candleServices "gitlab.com/hari-92/nft-market-server/internal/modules/candle/services"
)

type CandleController struct {
	candleService candleServices.CandleService
}

func NewCandleController(candleService candleServices.CandleService) *CandleController {
	return &CandleController{
		candleService: candleService,
	}
}
