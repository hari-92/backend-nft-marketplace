package candle_stick_errors

import "gitlab.com/hari-92/nft-market-server/pkg/errors"

const (
	CandleStickErrorCodeUnknown = iota * errors.BaseErrorCodeCandleStick
	CandleStickErrorCodeNotFound
)

var (
	CandleStickErrorUnknown  = errors.NewAppError(CandleStickErrorCodeUnknown, "Unknown error")
	CandleStickErrorNotFound = errors.NewAppError(CandleStickErrorCodeNotFound, "Candle stick not found")
)
