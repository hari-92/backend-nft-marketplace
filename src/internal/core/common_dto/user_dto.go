package common_dto

import "gitlab.com/hari-92/nft-market-server/internal/core/constant"

type CreateUserPayload struct {
	Username         string                    `json:"username"`
	Password         string                    `json:"password"`
	AuthenticateType constant.AuthenticateType `json:"authenticate_type"`
}

type CreateUserResponse struct {
	Id    uint32 `json:"id"`
	Email string `json:"email"`
}
