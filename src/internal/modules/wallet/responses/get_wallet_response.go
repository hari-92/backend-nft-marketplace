package wallet_responses

type GetWalletResponse struct {
	UserID        uint32  `json:"user_id,omitempty"`
	Address       string  `json:"address,omitempty"`
	ChainID       uint32  `json:"chain_id,omitempty"`
	TokenID       uint32  `json:"token_id,omitempty"`
	Balance       float64 `json:"balance,omitempty"`
	LockedBalance float64 `json:"locked_balance,omitempty"`
}
