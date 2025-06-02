package token_services

import (
	commonProducers "gitlab.com/hari-92/nft-market-server/internal/core/common_modules/producers"
	commonProducersEvent "gitlab.com/hari-92/nft-market-server/internal/core/common_modules/producers/events/token"
	tokenModels "gitlab.com/hari-92/nft-market-server/internal/modules/token/models"
	tokenRepositories "gitlab.com/hari-92/nft-market-server/internal/modules/token/repositories"
	tokenRequests "gitlab.com/hari-92/nft-market-server/internal/modules/token/requests"
	tokenResponses "gitlab.com/hari-92/nft-market-server/internal/modules/token/responses"
)

type ITokenService interface {
	GetTokens(request *tokenRequests.GetTokens) (*tokenResponses.GetTokens, error)
	CreateToken(request *tokenRequests.CreateToken) (*tokenResponses.CreateToken, error)
}

func NewTokenService(
	tokenRepository tokenRepositories.TokenRepository,
) ITokenService {
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

func (s *tokenService) CreateToken(request *tokenRequests.CreateToken) (*tokenResponses.CreateToken, error) {
	token := &tokenModels.Token{
		Address:     request.Address,
		Name:        request.Name,
		Symbol:      request.Symbol,
		Description: request.Description,
		Decimals:    request.Decimals,
		TotalSupply: request.TotalSupply,
		ChainID:     request.ChainID,
	}
	token, err := s.tokenRepository.Create(token)
	if err != nil {
		return nil, err
	}
	commonProducers.InitTokenEvent(&commonProducersEvent.InitTokenEvent{
		TokenId: token.ID,
	})
	return &tokenResponses.CreateToken{
		ID: token.ID,
	}, nil
}
