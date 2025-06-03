package token_requests

type UpdateTokenRequest struct {
	Address     *string `json:"address"`
	Symbol      *string `json:"symbol"`
	Name        *string `json:"name"`
	Description *string `json:"description"`
	ChainID     *uint32 `json:"chain_id"`
}
