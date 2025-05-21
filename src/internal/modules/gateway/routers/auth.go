package gateway_routers

func (p RegisterRoutersIn) RegisterAuthRoutes() {
	group := p.Engine.Group("/auth")

	// Guest
	group.POST("/login", p.AuthController.Login)

	// User
	group.POST("/register", p.AuthController.Register)
	group.POST("/logout", p.AuthController.Logout)

	// Callback
	group.POST("/callback/google", p.AuthController.GoogleCallback)
}
