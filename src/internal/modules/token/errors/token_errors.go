package token_errors

import "gitlab.com/hari-92/nft-market-server/pkg/errors"

const (
	TokenErrorCodeUnknown = iota * errors.BaseErrorCodeToken
	TokenErrorCodeNotFound
)

var (
	TokenErrorUnknown  = errors.NewAppError(TokenErrorCodeUnknown, "Unknown error")
	TokenErrorNotFound = errors.NewAppError(TokenErrorCodeNotFound, "Token not found")
)
