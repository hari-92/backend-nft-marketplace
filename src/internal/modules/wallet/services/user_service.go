package wallet_services

import (
	wallet_models "gitlab.com/hari-92/nft-market-server/internal/modules/wallet/models"
	wallet_repositories "gitlab.com/hari-92/nft-market-server/internal/modules/wallet/repositories"
)

type UserService interface {
	GetBalances(userID uint32) ([]*wallet_models.Wallet, error)
}

func NewUserService() UserService {
	return &userService{}
}

type userService struct {
	walletRepository wallet_repositories.IWalletRepository
}

func (s *userService) GetBalances(userID uint32) ([]*wallet_models.Wallet, error) {
	return s.walletRepository.GetBalances(userID)
}
