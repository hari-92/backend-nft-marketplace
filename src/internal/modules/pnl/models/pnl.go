package pnl_models

import (
	"time"

	"gitlab.com/hari-92/nft-market-server/internal/core/base"
)

type PNL struct {
	base.Model
	UserID        uint32    `json:"user_id" gorm:"not null"`
	PairID        uint32    `json:"pair_id" gorm:"not null"`
	RealizedPnl   float64   `json:"realized_pnl" gorm:"type:decimal(18,8);not null"`
	UnrealizedPnl float64   `json:"unrealized_pnl" gorm:"type:decimal(18,8);not null"`
	CalculatedAt  time.Time `json:"calculated_at" gorm:"not null"`
}
