package token_controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	tokenRequests "gitlab.com/hari-92/nft-market-server/internal/modules/token/requests"
	tokenServices "gitlab.com/hari-92/nft-market-server/internal/modules/token/services"
)

type TokenController struct {
	tokenService tokenServices.ITokenService
}

func NewTokenController(tokenService tokenServices.ITokenService) *TokenController {
	return &TokenController{
		tokenService: tokenService,
	}
}

func (co *TokenController) GetTokens(ctx *gin.Context) {
	chainId := ctx.Query("chain_id")
	chainIdUint, err := strconv.ParseUint(chainId, 10, 32)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "Bad Request", "error": "Chain ID is not a valid uint32"})
		return
	}
	res, err := co.tokenService.GetTokens(&tokenRequests.GetTokensRequest{
		ChainID: uint32(chainIdUint),
	})
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "Server Internal Error", "error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "Get Tokens", "data": res})
}
