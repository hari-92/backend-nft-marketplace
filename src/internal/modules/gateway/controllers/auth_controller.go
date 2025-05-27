package gateway_controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/golibs-starter/golib/log"
	gatewayRequest "gitlab.com/hari-92/nft-market-server/internal/modules/gateway/requests/auth"
	gateway_services "gitlab.com/hari-92/nft-market-server/internal/modules/gateway/services"
)

type AuthController struct {
	authService gateway_services.AuthService
}

func NewAuthController(
	authService gateway_services.AuthService,
) *AuthController {
	return &AuthController{
		authService: authService,
	}
}

// Login (POST /api/v1/auth/login)
func (c *AuthController) Login(ctx *gin.Context) {
	res, err := c.authService.Login(&gatewayRequest.Login{})
	if err != nil {
		ctx.JSON(500, gin.H{"message": "Server Internal Error", "error": err.Error()})
		return
	}
	ctx.JSON(200, gin.H{"message": "Login successfully", "data": res})
}

// Register (POST /api/v1/auth/register)
func (c *AuthController) Register(ctx *gin.Context) {
	var request gatewayRequest.Register
	if err := ctx.ShouldBindJSON(&request); err != nil {
		ctx.JSON(400, gin.H{"message": "Bad Request", "error": err.Error()})
		return
	}
	res, err := c.authService.Register(&request)
	if err != nil {
		ctx.JSON(500, gin.H{"message": "Server Internal Error", "error": err.Error()})
		return
	}
	ctx.JSON(200, gin.H{"message": "Register successfully", "data": res})
}

// Logout (POST /api/v1/auth/logout)
func (c *AuthController) Logout(ctx *gin.Context) {
	res, err := c.authService.Logout(&gatewayRequest.Logout{})
	if err != nil {
		ctx.JSON(500, gin.H{"message": "Server Internal Error", "error": err.Error()})
		return
	}
	ctx.JSON(200, gin.H{"message": "Logout successfully", "data": res})
}

// GoogleCallback callback from google oauth (GET /auth/callback/google)
func (c *AuthController) GoogleCallback(ctx *gin.Context) {
	log.Info("Google Callback called", ctx.Request)
	res, err := c.authService.GoogleCallback(&gatewayRequest.GoogleCallbackRequest{})
	if err != nil {
		ctx.JSON(500, gin.H{"message": "Server Internal Error", "error": err.Error()})
		return
	}
	ctx.JSON(200, gin.H{"message": "Google Callback successfully", "data": res})
}
