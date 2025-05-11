package gateway_routers

func (p RegisterRoutersIn) RegisterAuthRoutes() {
	group := p.Engine.Group("/auth")
	group.POST("/login", p.AuthController.Login)
	group.POST("/register", p.AuthController.Register)
	group.POST("/logout", p.AuthController.Logout)
}
