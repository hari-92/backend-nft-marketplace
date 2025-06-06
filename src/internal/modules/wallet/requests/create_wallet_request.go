package wallet_requests

type CreateWalletRequest struct {
	UserID uint32 `json:"user_id" validate:"required"`
}
