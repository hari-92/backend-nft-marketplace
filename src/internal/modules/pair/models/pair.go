package pair_models

import "gitlab.com/hari-92/nft-market-server/internal/core/base"

type Pair struct {
	base.Model
	BaseTokenID  uint32 `json:"base_token_id" gorm:"not null"`
	QuoteTokenID uint32 `json:"quote_token_id" gorm:"not null"`
	Symbol       string `json:"symbol" gorm:"type:varchar(10);not null"`
}
