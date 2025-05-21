package order_models

import "gitlab.com/hari-92/nft-market-server/internal/core/base"

type Order struct {
	base.Model
	UserID         uint32  `json:"user_id" gorm:"not null"`
	PairID         uint32  `json:"pair_id" gorm:"not null"`
	OrderType      string  `json:"order_type" gorm:"type:varchar(10);not null;check:order_type in ('LIMIT','MARKET')"`
	Side           string  `json:"side" gorm:"type:varchar(10);not null;check:side in ('BUY','SELL')"`
	Price          float64 `json:"price" gorm:"type:decimal(18,8);not null"`
	Quantity       float64 `json:"quantity" gorm:"type:decimal(18,8);not null"`
	FilledQuantity float64 `json:"filled_quantity" gorm:"type:decimal(18,8);default:0.0"`
	Status         string  `json:"status" gorm:"type:varchar(10);not null;check:status in ('OPEN','PENDING','FILLED','CANCELLED')"`
}
