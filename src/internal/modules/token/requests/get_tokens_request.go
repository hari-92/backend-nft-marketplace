package token_requests

type GetTokens struct {
	UserID uint `json:"user_id" binding:"required"`
}
