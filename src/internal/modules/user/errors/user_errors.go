package user_errors

import "gitlab.com/hari-92/nft-market-server/pkg/errors"

const (
	UserErrorCodeUnknown = iota * errors.BaseErrorCodeUser
	UserErrorCodeNotFound
)

var (
	UserErrorUnknown  = errors.NewAppError(UserErrorCodeUnknown, "Unknown error")
	UserErrorNotFound = errors.NewAppError(UserErrorCodeNotFound, "User not found")
)
