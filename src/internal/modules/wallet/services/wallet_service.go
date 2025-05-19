package wallet_services

import (
	walletModels "gitlab.com/hari-92/nft-market-server/internal/modules/wallet/models"
	walletRepositories "gitlab.com/hari-92/nft-market-server/internal/modules/wallet/repositories"
)

type IWalletService interface {
	GetBalance(userID uint32) ([]*walletModels.Wallet, error)
}

func NewWalletService(
	walletRepository walletRepositories.IWalletRepository,
) IWalletService {
	return &WalletService{
		walletRepository: walletRepository,
	}
}

type WalletService struct {
	walletRepository walletRepositories.IWalletRepository
}

func (w *WalletService) GetBalance(userID uint32) ([]*walletModels.Wallet, error) {
	wallets, err := w.walletRepository.GetBalance(userID)
	if err != nil {
		return nil, err
	}

	return wallets, nil
}
