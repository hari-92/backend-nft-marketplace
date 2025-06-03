package token_requests

type GetTokensRequest struct {
	ChainID uint32 `json:"chain_id" binding:"required"`
}
