package token_controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	tokenRequests "gitlab.com/hari-92/nft-market-server/internal/modules/token/requests"
	tokenServices "gitlab.com/hari-92/nft-market-server/internal/modules/token/services"
)

type TokenController struct {
	tokenService tokenServices.TokenService
}

func NewTokenController(tokenService tokenServices.TokenService) *TokenController {
	return &TokenController{
		tokenService: tokenService,
	}
}

func (co *TokenController) GetTokens(ctx *gin.Context) {
	res, err := co.tokenService.GetTokens(&tokenRequests.GetTokens{})
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "Server Internal Error", "error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "Get Tokens", "data": res})
}
