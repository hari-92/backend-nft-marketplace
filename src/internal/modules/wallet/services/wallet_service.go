package wallet_services

import (
	core_crypto "gitlab.com/hari-92/nft-market-server/internal/core/crypto"
	crypto_evm "gitlab.com/hari-92/nft-market-server/internal/core/crypto/evm"
	core_properties "gitlab.com/hari-92/nft-market-server/internal/core/properties"
	walletFilters "gitlab.com/hari-92/nft-market-server/internal/modules/wallet/filters"
	walletModels "gitlab.com/hari-92/nft-market-server/internal/modules/wallet/models"
	walletRepositories "gitlab.com/hari-92/nft-market-server/internal/modules/wallet/repositories"
	walletRequests "gitlab.com/hari-92/nft-market-server/internal/modules/wallet/requests"
	walletResponses "gitlab.com/hari-92/nft-market-server/internal/modules/wallet/responses"
)

type IWalletService interface {
	GetBalance(userID uint32) ([]*walletModels.Wallet, error)
	GetOne(filter *walletRequests.GetWalletRequest) (*walletResponses.GetWalletResponse, error)
	CreateWallet(req *walletRequests.CreateWalletRequest) (*walletResponses.CreateWalletResponse, error)
}

func NewWalletService(
	walletRepository walletRepositories.IWalletRepository,
	secretKeyProperties core_properties.SecretKeyProperties,
) IWalletService {
	return &WalletService{
		walletRepository:    walletRepository,
		secretKeyProperties: secretKeyProperties,
	}
}

type WalletService struct {
	walletRepository    walletRepositories.IWalletRepository
	secretKeyProperties core_properties.SecretKeyProperties
}

func (w *WalletService) GetBalance(userID uint32) ([]*walletModels.Wallet, error) {
	wallets, err := w.walletRepository.GetBalance(userID)
	if err != nil {
		return nil, err
	}

	return wallets, nil
}

func (w *WalletService) GetOne(filter *walletRequests.GetWalletRequest) (*walletResponses.GetWalletResponse, error) {
	wallet, err := w.walletRepository.GetOne(&walletFilters.WalletFilter{
		UserID:  filter.UserID,
		ChainID: filter.ChainID,
		Address: filter.Address,
		TokenID: filter.TokenID,
	})
	if err != nil {
		return nil, err
	}

	if wallet == nil {
		return nil, nil
	}

	response := &walletResponses.GetWalletResponse{
		UserID:  wallet.UserID,
		Address: wallet.Address,
		ChainID: wallet.ChainID,
	}

	return response, nil
}

func (w *WalletService) CreateWallet(req *walletRequests.CreateWalletRequest) (*walletResponses.CreateWalletResponse, error) {
	privateKey, address, err := crypto_evm.GenerateWallet()
	if err != nil {
		return nil, err
	}
	newWallet := &walletModels.Wallet{
		UserID:              req.UserID,
		EncryptedPrivateKey: core_crypto.Encrypt(privateKey, w.secretKeyProperties.SecretKey),
		Address:             address,
		StructAddressType:   walletModels.StructAddressTypeEVM,
		ChainID:             1, // Assuming 1 is the EVM chain ID, adjust as necessary
	}

	wallet, err := w.walletRepository.Create(newWallet)
	if err != nil {
		return nil, err
	}

	response := &walletResponses.CreateWalletResponse{
		ID:     uint32(wallet.ID),
		UserID: newWallet.UserID,
	}

	return response, nil
}
