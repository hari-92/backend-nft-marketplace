package token_models

import "gitlab.com/hari-92/nft-market-server/internal/core/base"

type Token struct {
	base.Model
	Address     string `json:"address" gorm:"type:varchar(42);not null"`
	Symbol      string `json:"symbol" gorm:"type:varchar(10);not null"`
	Name        string `json:"name" gorm:"type:varchar(100);not null"`
	Description string `json:"description" gorm:"type:text"`
	Decimals    int    `json:"decimals" gorm:"type:int;not null"`
	TotalSupply uint64 `json:"total_supply" gorm:"type:bigint;not null"`
	ChainID     int    `json:"chain_id" gorm:"type:int;not null"`
	IsActive    bool   `json:"is_active" gorm:"type:boolean;default:false"`
}
