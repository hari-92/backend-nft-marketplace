package gateway_services

import (
	commonDto "gitlab.com/hari-92/nft-market-server/internal/core/common_dto"
	gatewayErrors "gitlab.com/hari-92/nft-market-server/internal/modules/gateway/errors"
	gatewayInstance "gitlab.com/hari-92/nft-market-server/internal/modules/gateway/instance"
	gatewayRequest "gitlab.com/hari-92/nft-market-server/internal/modules/gateway/requests/auth"
	gatewayResponse "gitlab.com/hari-92/nft-market-server/internal/modules/gateway/responses/auth"
)

type AuthService interface {
	Login(request *gatewayRequest.Login) (response *gatewayResponse.Login, err error)
	Register(request *gatewayRequest.Register) (response *gatewayResponse.RegisterResponse, err error)
	Logout(request *gatewayRequest.Logout) (response *gatewayResponse.Logout, err error)
	GoogleCallback(request *gatewayRequest.GoogleCallbackRequest) (response *gatewayResponse.GoogleCallbackResponse, err error)
}

func NewAuthService() AuthService {
	return &authService{}
}

type authService struct {
}

func (s *authService) Login(request *gatewayRequest.Login) (response *gatewayResponse.Login, err error) {
	return &gatewayResponse.Login{
		AccessToken: "",
	}, nil
}

func (s *authService) Register(req *gatewayRequest.Register) (response *gatewayResponse.RegisterResponse, err error) {
	exist, err := gatewayInstance.UserRpcPortGateway.IsExistUser(req.Username)
	if err != nil {
		return nil, err
	}

	if exist {
		return nil, gatewayErrors.GatewayErrorUserAlreadyExists
	}

	user, err := gatewayInstance.UserRpcPortGateway.CreateUser(&commonDto.CreateUserPayload{
		Username: req.Username,
		Password: req.Password,
	})
	if err != nil {
		return nil, err
	}
	getToken := func(x *commonDto.CreateUserResponse) struct {
		AccessToken  string
		RefreshToken string
	} {
		return struct {
			AccessToken  string
			RefreshToken string
		}{AccessToken: x.Email, RefreshToken: x.Email}
	}

	return &gatewayResponse.RegisterResponse{
		AccessToken:  getToken(user).AccessToken,
		RefreshToken: getToken(user).RefreshToken,
	}, nil
}

func (s *authService) Logout(request *gatewayRequest.Logout) (response *gatewayResponse.Logout, err error) {
	return &gatewayResponse.Logout{}, nil
}

func (s *authService) GoogleCallback(request *gatewayRequest.GoogleCallbackRequest) (response *gatewayResponse.GoogleCallbackResponse, err error) {
	return &gatewayResponse.GoogleCallbackResponse{}, nil
}
