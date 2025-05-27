package user_services

import (
	"errors"
	userConverters "gitlab.com/hari-92/nft-market-server/internal/modules/user/converters"
	userRepositories "gitlab.com/hari-92/nft-market-server/internal/modules/user/repositories"
	userRequests "gitlab.com/hari-92/nft-market-server/internal/modules/user/requests"
	userResponses "gitlab.com/hari-92/nft-market-server/internal/modules/user/responses"
	"gorm.io/gorm"
)

type UserService interface {
	IsExist(req *userRequests.IsExistUserRequest) (bool, error)
	GetUser(req *userRequests.GetUserRequest) (*userResponses.GetUserResponse, error)
	CreateUser(req *userRequests.CreateUserRequest) (*userResponses.CreateUserResponse, error)
}

type userService struct {
	userRepository           userRepositories.UserRepository
	requestToFilterConverter *userConverters.RequestToFilterConverter
	modelToResponseConverter *userConverters.ModelToResponseConverter
	requestToModelConverter  *userConverters.RequestToModelConverter
}

func NewUserService(userRepository userRepositories.UserRepository) UserService {
	return &userService{
		userRepository:           userRepository,
		requestToFilterConverter: userConverters.NewRequestToFilterConverter(),
		modelToResponseConverter: userConverters.NewModelToResponseConverter(),
		requestToModelConverter:  userConverters.NewRequestToModelConverter(),
	}
}

func (u *userService) IsExist(req *userRequests.IsExistUserRequest) (bool, error) {
	user, err := u.userRepository.GetOne(u.requestToFilterConverter.ToUserFilter(&userRequests.GetUserRequest{
		Email: req.Username,
	}))
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return false, nil
		}
		return false, err
	}
	if user == nil {
		return false, nil
	}
	return true, nil
}

func (u *userService) GetUser(req *userRequests.GetUserRequest) (*userResponses.GetUserResponse, error) {
	user, err := u.userRepository.GetOne(u.requestToFilterConverter.ToUserFilter(req))
	if err != nil {
		return nil, err
	}
	return u.modelToResponseConverter.ToGetUserResponse(user), nil
}

func (u *userService) CreateUser(req *userRequests.CreateUserRequest) (*userResponses.CreateUserResponse, error) {
	user, err := u.userRepository.Create(u.requestToModelConverter.ToCreateUserModel(req))
	if err != nil {
		return nil, err
	}
	return u.modelToResponseConverter.ToCreateUserResponse(user), nil
}
