package gateway_routers

func (p RegisterRoutersIn) RegisterTokenRoutes() {
	tokenRouter := p.Engine.Group("/tokens")

	// Guest
	tokenRouter.GET("/", p.TokenController.GetTokens)
	tokenRouter.GET("/:token_id", p.TokenController.GetToken)

	// Admin
	tokenRouter.POST("/", p.TokenController.PostToken)
	tokenRouter.PUT("/:token_id", p.TokenController.PutToken)
	tokenRouter.DELETE("/:token_id", p.TokenController.DeleteToken)
	tokenRouter.POST("/validate", p.TokenController.PostValidateToken)
}
