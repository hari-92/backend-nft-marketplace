package user_models

import (
	"gitlab.com/hari-92/nft-market-server/internal/core/base"
)

type UserToken struct {
	base.Model
	UserID        uint32  `json:"user_id" gorm:"not null"`
	TokenID       uint32  `json:"token_id" gorm:"not null"`
	WalletID      uint32  `json:"wallet_id" gorm:"not null"`
	Balance       float64 `json:"balance" gorm:"type:decimal(18,8);not null;default:0.0"`
	LockedBalance float64 `json:"locked_balance" gorm:"type:decimal(18,8);not null;default:0.0"`
}
