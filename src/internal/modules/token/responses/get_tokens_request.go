package token_responses

type GetTokensResponse struct {
	ID          uint32 `json:"id"`
	Address     string `json:"address"`
	Symbol      string `json:"symbol"`
	Decimals    int32  `json:"decimals"`
	TotalSupply uint64 `json:"total_supply"`
	Name        string `json:"name"`
	Description string `json:"description"`
	ChainID     uint32 `json:"chain_id"`
}
