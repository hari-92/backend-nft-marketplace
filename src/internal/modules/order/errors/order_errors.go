package order_errors

import "gitlab.com/hari-92/nft-market-server/pkg/errors"

const (
	OrderErrorCodeUnknown = iota * errors.BaseErrorCodeOrder
	OrderErrorCodeNotFound
)

var (
	OrderErrorUnknown  = errors.NewAppError(OrderErrorCodeUnknown, "Unknown error")
	OrderErrorNotFound = errors.NewAppError(OrderErrorCodeNotFound, "Order not found")
)
