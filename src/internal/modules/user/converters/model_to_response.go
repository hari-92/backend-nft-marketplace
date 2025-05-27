package user_converters

import (
	userModels "gitlab.com/hari-92/nft-market-server/internal/modules/user/models"
	userResponses "gitlab.com/hari-92/nft-market-server/internal/modules/user/responses"
)

type ModelToResponseConverter struct{}

func NewModelToResponseConverter() *ModelToResponseConverter {
	return &ModelToResponseConverter{}
}

func (c *ModelToResponseConverter) ToGetUserResponse(model *userModels.User) *userResponses.GetUserResponse {
	if model == nil {
		return nil
	}

	return &userResponses.GetUserResponse{
		ID:    uint32(model.ID),
		Email: model.Email,
	}
}

func (c *ModelToResponseConverter) ToCreateUserResponse(model *userModels.User) *userResponses.CreateUserResponse {
	if model == nil {
		return nil
	}

	return &userResponses.CreateUserResponse{
		ID: uint32(model.ID),
	}
}
