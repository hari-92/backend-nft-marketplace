package gateway_routers

func (p *RegisterRoutersIn) RegisterCandleRoutes() {
	candleRouter := p.Engine.Group("/candles")

	// Guest
	candleRouter.GET("", p.CandleController.GetCandles)
	candleRouter.GET("/:pair/:time_frame/stats", p.CandleController.GetCandleStats)
	candleRouter.GET("/:pair/:time_frame/latest", p.CandleController.GetCandleLatest)
	candleRouter.POST("/validate", p.CandleController.PostCandleValidate)
	candleRouter.GET("/:pair/aggregate", p.CandleController.GetCandlesAggregate)
}
