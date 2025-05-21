package gateway_routers

func (p RegisterRoutersIn) RegisterPnlRoutes() {
	pnlRouter := p.Engine.Group("/v1/pnl")

	// User
	pnlRouter.GET("/realized", p.PnlController.GetRealized)
	pnlRouter.GET("/unrealized", p.PnlController.GetUnrealized)
	pnlRouter.GET("/summary", p.PnlController.GetSummary)
	pnlRouter.GET("/portfolio", p.PnlController.GetPortfolio)
	pnlRouter.POST("/recalculate", p.PnlController.PostRecalculate)
	pnlRouter.GET("/history", p.PnlController.GetHistory)
	pnlRouter.POST("/validate", p.PnlController.PostValidate)
	pnlRouter.GET("/pair/:pair_id", p.PnlController.GetPnlPair)
	pnlRouter.GET("/export", p.PnlController.GetExport)
}
