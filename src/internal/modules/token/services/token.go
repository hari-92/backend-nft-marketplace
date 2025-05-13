package token_services

import (
	tokenRequests "gitlab.com/hari-92/nft-market-server/internal/modules/token/requests"
	tokenResponses "gitlab.com/hari-92/nft-market-server/internal/modules/token/responses"
)

type TokenService interface {
	GetTokens(request *tokenRequests.GetTokens) (*tokenResponses.GetTokens, error)
}

func NewTokenService() TokenService {
	return &tokenService{}
}

type tokenService struct {
}

func (s *tokenService) GetTokens(request *tokenRequests.GetTokens) (*tokenResponses.GetTokens, error) {
	return &tokenResponses.GetTokens{}, nil
}
