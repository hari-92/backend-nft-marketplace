package gateway_routers

func (p *RegisterRoutersIn) RegisterPairRoutes() {
	pairRouter := p.Engine.Group("/pairs")

	// Guest
	pairRouter.GET("", p.PairController.GetPairs)
	pairRouter.GET("/:pair_id", p.PairController.GetPair)

	// Admin
	pairRouter.POST("", p.PairController.PostPair)
	pairRouter.PUT("/:pair_id", p.PairController.PutPair)
	pairRouter.DELETE("/:pair_id", p.PairController.DeletePair)
	pairRouter.POST("/validate", p.PairController.PostValidatePair)
}
