package wallet_services

type UserService interface {
}

func NewUserService() UserService {
	return &userService{}
}

type userService struct {
}
