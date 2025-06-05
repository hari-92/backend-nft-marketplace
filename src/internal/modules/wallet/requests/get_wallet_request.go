package wallet_requests

type GetWalletRequest struct {
	UserID  uint32 `json:"user_id,omitempty" form:"user_id,omitempty"`
	Address string `json:"address,omitempty" form:"address,omitempty"`
	ChainID uint32 `json:"chain_id,omitempty" form:"chain_id,omitempty"`
	TokenID uint32 `json:"token_id,omitempty" form:"token_id,omitempty"`
}
