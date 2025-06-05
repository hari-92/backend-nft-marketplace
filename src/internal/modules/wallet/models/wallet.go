package wallet_models

import "gitlab.com/hari-92/nft-market-server/internal/core/base"

type Wallet struct {
	base.Model
	UserID              uint32            `json:"user_id" gorm:"not null"`
	Address             string            `json:"address" gorm:"not null"`
	EncryptedPrivateKey string            `json:"encrypted_private_key" gorm:"not null"`
	StructAddressType   StructAddressType `json:"struct_address_type" gorm:"not null"`
	ChainID             uint32            `json:"chain_id" gorm:"not null"`
}

type StructAddressType string

const (
	StructAddressTypeEVM StructAddressType = "evm"
	StructAddressTypeSol StructAddressType = "sol"
)
