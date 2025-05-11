package wallet_controllers

import (
	"github.com/gin-gonic/gin"
	wallet_requests "gitlab.com/hari-92/nft-market-server/internal/modules/wallet/requests"
	walletServices "gitlab.com/hari-92/nft-market-server/internal/modules/wallet/services"
)

type WalletController struct {
	walletService walletServices.WalletService
}

func NewWalletController() *WalletController {
	return &WalletController{}
}

func (w *WalletController) GetBalances(ctx *gin.Context) {
	res, err := w.walletService.GetBalances(&wallet_requests.GetBalances{})
	if err != nil {
		ctx.JSON(500, gin.H{"message": "Server Internal Error", "error": err.Error()})
		return
	}

	ctx.JSON(200, gin.H{"message": "Get Balances", "data": res})
}
