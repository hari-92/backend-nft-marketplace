package trading_pair_errors

import "gitlab.com/hari-92/nft-market-server/pkg/errors"

const (
	TradingPairErrorCodeUnknown = iota * errors.BaseErrorCodeTradingPair
	TradingPairErrorCodeNotFound
)

var (
	TradingPairErrorUnknown  = errors.NewAppError(TradingPairErrorCodeUnknown, "Unknown error")
	TradingPairErrorNotFound = errors.NewAppError(TradingPairErrorCodeNotFound, "Trading pair not found")
)
