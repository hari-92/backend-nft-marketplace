package wallet_errors

import "gitlab.com/hari-92/nft-market-server/pkg/errors"

const (
	WalletErrorCodeUnknown = iota * errors.BaseErrorCodeWallet
	WalletErrorCodeNotFound
)

var (
	WalletErrorUnknown  = errors.NewAppError(WalletErrorCodeUnknown, "Unknown error")
	WalletErrorNotFound = errors.NewAppError(WalletErrorCodeNotFound, "Wallet not found")
)
