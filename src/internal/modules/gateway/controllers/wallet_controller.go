package gateway_controllers

import (
	"github.com/gin-gonic/gin"
	gatewayInstance "gitlab.com/hari-92/nft-market-server/internal/modules/gateway/instance"
)

type WalletController struct {
}

func NewWalletController() *WalletController {
	return &WalletController{}
}

// GetBalances (GET /wallet/balances)
func (w *WalletController) GetBalances(ctx *gin.Context) {
	res, err := gatewayInstance.WalletRpcPortGateway.GetBalances(ctx, 1)
	if err != nil {
		ctx.JSON(500, gin.H{"message": "Server Internal Error", "error": err.Error()})
		return
	}

	ctx.JSON(200, gin.H{"message": "Get Balances", "data": res})
}
