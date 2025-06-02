package gateway_request

import (
	"gitlab.com/hari-92/nft-market-server/internal/core/constant"
)

type Register struct {
	Username         string                    `json:"username"`
	Password         string                    `json:"password"`
	AuthenticateType constant.AuthenticateType `json:"authenticate_type"`
}
