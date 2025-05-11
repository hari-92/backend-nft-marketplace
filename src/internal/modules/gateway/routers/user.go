package gateway_routers

func (p RegisterRoutersIn) RegisterUserRoutes() {
	userRouter := p.Engine.Group("/user")
	userRouter.GET("/", p.UserController.HelloWorld)
}
