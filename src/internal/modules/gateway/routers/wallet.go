package gateway_routers

func (p RegisterRoutersIn) RegisterWalletRoutes() {
	walletRouter := p.Engine.Group("/wallet")
	walletRouter.GET("/balances", p.WalletController.GetBalances)
}
