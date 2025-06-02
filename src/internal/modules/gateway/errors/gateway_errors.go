package gateway_errors

import "gitlab.com/hari-92/nft-market-server/pkg/errors"

const (
	GatewayErrorCodeUnknown = iota * errors.BaseErrorCodeGateway
	GatewayErrorCodeUnauthorized
	GatewayErrorCodeInvalidRequest
	GatewayErrorCodeUserAlreadyExists
	GatewayErrorCodeLoginFailed
)

var (
	GatewayErrorUnknown           = errors.NewAppError(GatewayErrorCodeUnknown, "Unknown error")
	GatewayErrorUnauthorized      = errors.NewAppError(GatewayErrorCodeUnauthorized, "Unauthorized access")
	GatewayErrorInvalidRequest    = errors.NewAppError(GatewayErrorCodeInvalidRequest, "Invalid request")
	GatewayErrorUserAlreadyExists = errors.NewAppError(GatewayErrorCodeUserAlreadyExists, "User already exists")
	GatewayErrorLoginFailed       = errors.NewAppError(GatewayErrorCodeLoginFailed, "Login failed, please check your credentials")
)
