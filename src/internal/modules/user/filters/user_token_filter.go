package user_filters

type UserTokenFilter struct {
	UserID        uint32
	TokenID       uint32
	InTokenID     []uint32
	NotInTokenID  []uint32
	WalletID      uint32
	InWalletID    []uint32
	NotInWalletID []uint32
}
