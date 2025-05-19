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

// GetBalance: Get balance of a user (GET /wallets/balance)
func (w *WalletController) GetBalance(ctx *gin.Context) {
	// TODO: Get userID from context
	userID := 1
	res, err := gatewayInstance.WalletRpcPortGateway.GetBalance(ctx, uint32(userID))
	if err != nil {
		ctx.JSON(500, gin.H{"message": "Server Internal Error", "error": err.Error()})
		return
	}

	ctx.JSON(200, gin.H{"message": "Get Balance", "data": res})
}

// PostWallet: Create a new wallet for a user (POST /wallet)
func (w *WalletController) PostWallet(ctx *gin.Context) {
	// TODO: Implement this
	ctx.JSON(200, gin.H{"message": "Post Wallet"})
}

// PostDeposit: Deposit money to a wallet (POST /wallets/deposit)
func (w *WalletController) PostDeposit(ctx *gin.Context) {
	// TODO: Implement this
	ctx.JSON(200, gin.H{"message": "Post Deposit"})
}

// PostWithdraw: Withdraw money from a wallet (POST /wallets/withdraw)
func (w *WalletController) PostWithdraw(ctx *gin.Context) {
	// TODO: Implement this
	ctx.JSON(200, gin.H{"message": "Post Withdraw"})
}

// GetHistory: Get history of a wallet (GET /wallets/history)
func (w *WalletController) GetHistory(ctx *gin.Context) {
	// TODO: Implement this
	ctx.JSON(200, gin.H{"message": "Get History"})
}

// PostTransfer: Transfer money between wallets (POST /wallets/transfer)
func (w *WalletController) PostTransfer(ctx *gin.Context) {
	// TODO: Implement this
	ctx.JSON(200, gin.H{"message": "Post Transfer"})
}

// GetTransaction: Get a transaction by txID (GET /wallets/transactions/:tx_id)
func (w *WalletController) GetTransaction(ctx *gin.Context) {
	// TODO: Implement this
	ctx.JSON(200, gin.H{"message": "Get Transaction"})
}

// PostFreeze: Freeze a wallet (use with admin) (POST /wallets/freeze)
func (w *WalletController) PostFreeze(ctx *gin.Context) {
	// TODO: Implement this
	ctx.JSON(200, gin.H{"message": "Post Freeze"})
}

// PostEstimateFee: Estimate fee for a transaction (POST /wallets/estimate-fee)
func (w *WalletController) PostEstimateFee(ctx *gin.Context) {
	// TODO: Implement this
	ctx.JSON(200, gin.H{"message": "Post Estimate Fee"})
}
