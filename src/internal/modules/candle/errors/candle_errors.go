package candle_errors

import "gitlab.com/hari-92/nft-market-server/pkg/errors"

const (
	CandleErrorCodeUnknown = iota * errors.BaseErrorCodeCandle
	CandleErrorCodeNotFound
)

var (
	CandleErrorUnknown  = errors.NewAppError(CandleErrorCodeUnknown, "Unknown error")
	CandleErrorNotFound = errors.NewAppError(CandleErrorCodeNotFound, "Candle not found")
)
