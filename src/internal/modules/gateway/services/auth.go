package gateway_services

import (
	gatewayRequest "gitlab.com/hari-92/nft-market-server/internal/modules/gateway/requests/auth"
	gatewayResponse "gitlab.com/hari-92/nft-market-server/internal/modules/gateway/responses/auth"
)

type AuthService interface {
	Login(request *gatewayRequest.Login) (response *gatewayResponse.Login, err error)
	Register(request *gatewayRequest.Register) (response *gatewayResponse.Register, err error)
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

func (s *authService) Register(request *gatewayRequest.Register) (response *gatewayResponse.Register, err error) {
	return &gatewayResponse.Register{}, nil
}

func (s *authService) Logout(request *gatewayRequest.Logout) (response *gatewayResponse.Logout, err error) {
	return &gatewayResponse.Logout{}, nil
}

func (s *authService) GoogleCallback(request *gatewayRequest.GoogleCallbackRequest) (response *gatewayResponse.GoogleCallbackResponse, err error) {
	return &gatewayResponse.GoogleCallbackResponse{}, nil
}
