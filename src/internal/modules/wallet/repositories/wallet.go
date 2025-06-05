package wallet_repositories

import (
	walletFilters "gitlab.com/hari-92/nft-market-server/internal/modules/wallet/filters"
	walletModels "gitlab.com/hari-92/nft-market-server/internal/modules/wallet/models"
	walletScopes "gitlab.com/hari-92/nft-market-server/internal/modules/wallet/repositories/scopes"
	"gorm.io/gorm"
)

type IWalletRepository interface {
	GetBalance(userID uint32) ([]*walletModels.Wallet, error)
	GetOne(filter *walletFilters.WalletFilter) (*walletModels.Wallet, error)
	GetMany(filter *walletFilters.WalletFilter) ([]*walletModels.Wallet, error)
	Create(wallet *walletModels.Wallet) (*walletModels.Wallet, error)
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

func (w *WalletRepository) GetOne(filter *walletFilters.WalletFilter) (*walletModels.Wallet, error) {
	var wallet walletModels.Wallet
	if err := w.db.Model(&wallet).Scopes(walletScopes.WalletScope(filter)).First(&wallet).Error; err != nil {
		return nil, err
	}

	return &wallet, nil
}

func (w *WalletRepository) GetMany(filter *walletFilters.WalletFilter) ([]*walletModels.Wallet, error) {
	var wallets []*walletModels.Wallet
	if err := w.db.Model(&walletModels.Wallet{}).Scopes(walletScopes.WalletScope(filter)).Find(&wallets).Error; err != nil {
		return nil, err
	}

	return wallets, nil
}

func (w *WalletRepository) GetBalance(userID uint32) ([]*walletModels.Wallet, error) {
	// Create mock data with 10 different tokens
	var mockWallets []*walletModels.Wallet

	return mockWallets, nil
}

func (w *WalletRepository) Create(wallet *walletModels.Wallet) (*walletModels.Wallet, error) {
	if err := w.db.Create(&wallet).Error; err != nil {
		return nil, err
	}
	return wallet, nil
}
