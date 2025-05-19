package wallet_models

import "gitlab.com/hari-92/nft-market-server/internal/core/base"

type Wallet struct {
	base.Model
	UserID  uint32  `json:"user_id"`
	TokenID uint32  `json:"token_id"`
	Balance float64 `json:"balance"`
}
