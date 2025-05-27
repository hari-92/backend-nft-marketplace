package user_converters

import (
	userModels "gitlab.com/hari-92/nft-market-server/internal/modules/user/models"
	userRequests "gitlab.com/hari-92/nft-market-server/internal/modules/user/requests"
)

type RequestToModelConverter struct{}

func NewRequestToModelConverter() *RequestToModelConverter {
	return &RequestToModelConverter{}
}

func (c *RequestToModelConverter) ToCreateUserModel(req *userRequests.CreateUserRequest) *userModels.User {
	return &userModels.User{
		Email:    req.Email,
		Username: req.Email,
	}
}
