package controllers

import (
	"github.com/gin-gonic/gin"
	"gitlab.com/hari-92/nft-market-server/internal/modules/gateway/instance"
)

type UserController struct {
}

func NewUserController(router *gin.Engine) *UserController {
	return &UserController{}
}

func (c *UserController) HelloWorld(ctx *gin.Context) {
	res, err := instance.UserRpcPortGateway.HelloWorld("halllo")
	if err != nil {
		ctx.JSON(500, gin.H{"message": "Server Internal Error", "error": err.Error()})
		return
	}
	ctx.JSON(200, gin.H{"message": "Hello World", "status": "success", "data": res})
}
