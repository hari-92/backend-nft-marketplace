package gateway_errors

import "gitlab.com/hari-92/nft-market-server/pkg/errors"

const (
	GatewayErrorCodeUnknown = iota * errors.BaseErrorCodeGateway
	GatewayErrorCodeUnauthorized
	GatewayErrorCodeInvalidRequest
)

var (
	GatewayErrorUnknown        = errors.NewAppError(GatewayErrorCodeUnknown, "Unknown error")
	GatewayErrorUnauthorized   = errors.NewAppError(GatewayErrorCodeUnauthorized, "Unauthorized access")
	GatewayErrorInvalidRequest = errors.NewAppError(GatewayErrorCodeInvalidRequest, "Invalid request")
)
