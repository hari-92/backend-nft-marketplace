package wallet_models

type Wallet struct {
	UserID  uint32  `json:"user_id"`
	TokenID uint32  `json:"token_id"`
	Balance float64 `json:"balance"`
}
