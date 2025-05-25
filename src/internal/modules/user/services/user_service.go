package user_services

import (
	userConverters "gitlab.com/hari-92/nft-market-server/internal/modules/user/converters"
	userRepositories "gitlab.com/hari-92/nft-market-server/internal/modules/user/repositories"
	userRequests "gitlab.com/hari-92/nft-market-server/internal/modules/user/requests"
	userResponses "gitlab.com/hari-92/nft-market-server/internal/modules/user/responses"
)

type UserService interface {
	GetUser(req *userRequests.GetUserRequest) (*userResponses.GetUserResponse, error)
}

type userService struct {
	userRepository           userRepositories.UserRepository
	requestToFilterConverter *userConverters.RequestToFilterConverter
	modelToResponseConverter *userConverters.ModelToResponseConverter
}

func NewUserService(userRepository userRepositories.UserRepository) UserService {
	return &userService{
		userRepository:           userRepository,
		requestToFilterConverter: userConverters.NewRequestToFilterConverter(),
		modelToResponseConverter: userConverters.NewModelToResponseConverter(),
	}
}

func (u *userService) GetUser(req *userRequests.GetUserRequest) (*userResponses.GetUserResponse, error) {
	user, err := u.userRepository.GetOne(u.requestToFilterConverter.ToUserFilter(req))
	if err != nil {
		return nil, err
	}
	return u.modelToResponseConverter.ToGetUserResponse(user), nil
}
