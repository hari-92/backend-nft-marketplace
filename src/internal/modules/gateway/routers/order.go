package gateway_routers

func (p *RegisterRoutersIn) RegisterOrderRoutes() {
	orderRouter := p.Engine.Group("/orders")

	// User
	orderRouter.GET("", p.OrderController.GetOrders)
	orderRouter.GET("/:order_id", p.OrderController.GetOrder)
	orderRouter.POST("", p.OrderController.PostOrder)
	orderRouter.DELETE("/:order_id", p.OrderController.CancelOrder)
	orderRouter.DELETE("/bulk", p.OrderController.BulkCancelOrders)
	orderRouter.POST("/validate", p.OrderController.PostValidateOrder)
	orderRouter.GET("/history", p.OrderController.GetHistory)

	// Admin
	orderRouter.POST("/admin/cancel", p.OrderController.PostAdminCancel)
}
