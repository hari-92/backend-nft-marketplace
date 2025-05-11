package wallet_services

import (
	walletRequests "gitlab.com/hari-92/nft-market-server/internal/modules/wallet/requests"
	walletResponses "gitlab.com/hari-92/nft-market-server/internal/modules/wallet/responses"
)

type WalletService interface {
	GetBalances(request *walletRequests.GetBalances) (*walletResponses.GetBalances, error)
}

func NewWalletService() WalletService {
	return &walletService{}
}

type walletService struct {
}

func (w *walletService) GetBalances(request *walletRequests.GetBalances) (*walletResponses.GetBalances, error) {
	return &walletResponses.GetBalances{}, nil
}
