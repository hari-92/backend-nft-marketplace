package controllers

import (
	"github.com/gin-gonic/gin"
)

type UserController struct {
}

func NewUserController(router *gin.Engine) *UserController {
	return &UserController{}
}

func (c *UserController) HelloWorld(ctx *gin.Context) {
	ctx.JSON(200, gin.H{"message": "Hello World", "status": "success"})
}
