package wallet_filters

type WalletFilter struct {
	UserID       uint32   `json:"user_id,omitempty" form:"user_id,omitempty"`
	InUserID     []uint32 `json:"in_user_id,omitempty" form:"in_user_id,omitempty"`
	NotInUserID  []uint32 `json:"not_in_user_id,omitempty" form:"not_in_user_id,omitempty"`
	ChainID      uint32   `json:"chain_id,omitempty" form:"chain_id,omitempty"`
	InChainID    []uint32 `json:"in_chain_id,omitempty" form:"in_chain_id,omitempty"`
	NotInChainID []uint32 `json:"not_in_chain_id,omitempty" form:"not_in_chain_id,omitempty"`
	Address      string   `json:"address,omitempty" form:"address,omitempty"`
	TokenID      uint32   `json:"token_id,omitempty" form:"token_id,omitempty"`
	InTokenID    []uint32 `json:"in_token_id,omitempty" form:"in_token_id,omitempty"`
	NotInTokenID []uint32 `json:"not_in_token_id,omitempty" form:"not_in_token_id,omitempty"`
}
