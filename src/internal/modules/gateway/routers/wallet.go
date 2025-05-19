package gateway_routers

func (p RegisterRoutersIn) RegisterWalletRoutes() {
	walletRouter := p.Engine.Group("/wallets")

	// User
	walletRouter.GET("/balance", p.WalletController.GetBalance)
	walletRouter.POST("/", p.WalletController.PostWallet)
	walletRouter.POST("/deposit", p.WalletController.PostDeposit)
	walletRouter.POST("/withdraw", p.WalletController.PostWithdraw)
	walletRouter.GET("/history", p.WalletController.GetHistory)
	walletRouter.POST("/transfer", p.WalletController.PostTransfer)
	walletRouter.GET("/transactions/:tx_id", p.WalletController.GetTransaction)
	walletRouter.POST("/estimate-fee", p.WalletController.PostEstimateFee)

	// Admin
	walletRouter.POST("/freeze", p.WalletController.PostFreeze)
}
