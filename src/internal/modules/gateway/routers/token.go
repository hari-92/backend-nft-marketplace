package gateway_routers

func (p RegisterRoutersIn) RegisterTokenRoutes() {
	tokenRouter := p.Engine.Group("/api/v1/tokens")

	// Guest
	tokenRouter.GET("/", p.TokenController.GetTokens)
	tokenRouter.GET("/:id", p.TokenController.GetToken)

	// Admin
	tokenRouter.POST("/", p.TokenController.PostToken)
	tokenRouter.PUT("/:id", p.TokenController.PutToken)
	tokenRouter.DELETE("/:id", p.TokenController.DeleteToken)
	tokenRouter.POST("/validate", p.TokenController.PostValidateToken)
}
