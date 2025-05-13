package trading_pair_controllers

import (
	tradingPairServices "gitlab.com/hari-92/nft-market-server/internal/modules/trading_pair/services"
)

type TradingPairController struct {
	tradingPairService tradingPairServices.TradingPairService
}

func NewTradingPairController(tradingPairService tradingPairServices.TradingPairService) *TradingPairController {
	return &TradingPairController{
		tradingPairService: tradingPairService,
	}
}
