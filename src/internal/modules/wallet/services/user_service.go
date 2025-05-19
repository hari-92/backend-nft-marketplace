package wallet_services

import (
	wallet_repositories "gitlab.com/hari-92/nft-market-server/internal/modules/wallet/repositories"
)

type UserService interface {
}

func NewUserService() UserService {
	return &userService{}
}

type userService struct {
	walletRepository wallet_repositories.IWalletRepository
}
