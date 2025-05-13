package order_controllers

import (
	orderServices "gitlab.com/hari-92/nft-market-server/internal/modules/order/services"
)

type OrderController struct {
	orderService orderServices.OrderService
}

func NewOrderController(orderService orderServices.OrderService) *OrderController {
	return &OrderController{
		orderService: orderService,
	}
}
