package pnl_errors

import "gitlab.com/hari-92/nft-market-server/pkg/errors"

const (
	PnlErrorCodeUnknown = iota * errors.BaseErrorCodePnl
	PnlErrorCodeNotFound
)

var (
	PnlErrorUnknown  = errors.NewAppError(PnlErrorCodeUnknown, "Unknown error")
	PnlErrorNotFound = errors.NewAppError(PnlErrorCodeNotFound, "PNL record not found")
)
