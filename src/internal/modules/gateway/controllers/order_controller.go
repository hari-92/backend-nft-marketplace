package gateway_controllers

import (
	"github.com/gin-gonic/gin"
)

type OrderController struct {
}

func NewOrderController() *OrderController {
	return &OrderController{}
}

// GetOrders: Get all and paginated orders of a user (GET /orders)
func (o *OrderController) GetOrders(ctx *gin.Context) {
	// TODO: Implement this
	ctx.JSON(200, gin.H{"message": "Get Orders"})
}

// GetOrder: Get an order by id (GET /orders/:order_id)
func (o *OrderController) GetOrder(ctx *gin.Context) {
	// TODO: Implement this
	ctx.JSON(200, gin.H{"message": "Get Order"})
}

// PostOrder: Create a new order (POST /orders)
func (o *OrderController) PostOrder(ctx *gin.Context) {
	// TODO: Implement this
	ctx.JSON(200, gin.H{"message": "Post Order"})
}

// DeleteOrder: Delete an order by id (DELETE /orders/:order_id)
func (o *OrderController) DeleteOrder(ctx *gin.Context) {
	// TODO: Implement this
	ctx.JSON(200, gin.H{"message": "Delete Order"})
}

// DeleteOrders: Delete all orders of a user (DELETE /orders)
func (o *OrderController) DeleteOrders(ctx *gin.Context) {
	// TODO: Implement this
	ctx.JSON(200, gin.H{"message": "Delete Orders"})
}

// PostValidateOrder: Validate an order (POST /orders/validate)
func (o *OrderController) PostValidateOrder(ctx *gin.Context) {
	// TODO: Implement this
	ctx.JSON(200, gin.H{"message": "Post Validate Order"})
}

// GetHistory: Get the history of orders of a user (GET /orders/history)
func (o *OrderController) GetHistory(ctx *gin.Context) {
	// TODO: Implement this
	ctx.JSON(200, gin.H{"message": "Get History"})
}

// PostAdminCancel: Cancel an order (POST /orders/admin/cancel)
func (o *OrderController) PostAdminCancel(ctx *gin.Context) {
	// TODO: Implement this
	ctx.JSON(200, gin.H{"message": "Post Admin Cancel"})
}
