package wallet_repositories

import (
	wallet_models "gitlab.com/hari-92/nft-market-server/internal/modules/wallet/models"
	"gorm.io/gorm"
)

type IWalletRepository interface {
	GetBalance(userID uint32) ([]*wallet_models.Wallet, error)
}

func NewWalletRepository(
	db *gorm.DB,
) IWalletRepository {
	return &WalletRepository{
		db: db,
	}
}

type WalletRepository struct {
	db *gorm.DB
}

func (w *WalletRepository) GetBalance(userID uint32) ([]*wallet_models.Wallet, error) {
	// Create mock data with 10 different tokens
	mockWallets := []*wallet_models.Wallet{
		{
			UserID:  userID,
			TokenID: 1,
			Balance: 1000.50,
		},
		{
			UserID:  userID,
			TokenID: 2,
			Balance: 500.75,
		},
		{
			UserID:  userID,
			TokenID: 3,
			Balance: 2500.25,
		},
		{
			UserID:  userID,
			TokenID: 4,
			Balance: 750.00,
		},
		{
			UserID:  userID,
			TokenID: 5,
			Balance: 1500.80,
		},
		{
			UserID:  userID,
			TokenID: 6,
			Balance: 3000.00,
		},
		{
			UserID:  userID,
			TokenID: 7,
			Balance: 800.50,
		},
		{
			UserID:  userID,
			TokenID: 8,
			Balance: 1200.25,
		},
		{
			UserID:  userID,
			TokenID: 9,
			Balance: 950.75,
		},
		{
			UserID:  userID,
			TokenID: 10,
			Balance: 2000.00,
		},
	}

	return mockWallets, nil
}
