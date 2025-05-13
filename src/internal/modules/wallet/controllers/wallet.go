package wallet_controllers

import (
	walletServices "gitlab.com/hari-92/nft-market-server/internal/modules/wallet/services"
)

type WalletController struct {
	walletService walletServices.IWalletService
}

func NewWalletController() *WalletController {
	return &WalletController{}
}
