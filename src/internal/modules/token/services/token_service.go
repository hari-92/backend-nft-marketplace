package token_services

import (
	tokenRepositories "gitlab.com/hari-92/nft-market-server/internal/modules/token/repositories"
	tokenRequests "gitlab.com/hari-92/nft-market-server/internal/modules/token/requests"
	tokenResponses "gitlab.com/hari-92/nft-market-server/internal/modules/token/responses"
)

type TokenService interface {
	GetTokens(request *tokenRequests.GetTokens) (*tokenResponses.GetTokens, error)
}

func NewTokenService(
	tokenRepository tokenRepositories.TokenRepository,
) TokenService {
	return &tokenService{
		tokenRepository: tokenRepository,
	}
}

type tokenService struct {
	tokenRepository tokenRepositories.TokenRepository
}

func (s *tokenService) GetTokens(request *tokenRequests.GetTokens) (*tokenResponses.GetTokens, error) {
	return &tokenResponses.GetTokens{}, nil
}
