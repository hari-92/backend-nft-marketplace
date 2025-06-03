package token_services

import (
	commonProducers "gitlab.com/hari-92/nft-market-server/internal/core/common_modules/producers"
	commonProducersEvent "gitlab.com/hari-92/nft-market-server/internal/core/common_modules/producers/events/token"
	tokenFilters "gitlab.com/hari-92/nft-market-server/internal/modules/token/filters"
	tokenModels "gitlab.com/hari-92/nft-market-server/internal/modules/token/models"
	tokenRepositories "gitlab.com/hari-92/nft-market-server/internal/modules/token/repositories"
	tokenRequests "gitlab.com/hari-92/nft-market-server/internal/modules/token/requests"
	tokenResponses "gitlab.com/hari-92/nft-market-server/internal/modules/token/responses"
)

type ITokenService interface {
	GetToken(request *tokenRequests.GetTokenRequest) (*tokenResponses.GetTokenResponse, error)
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

func (s *tokenService) GetToken(request *tokenRequests.GetTokenRequest) (*tokenResponses.GetTokenResponse, error) {
	token, err := s.tokenRepository.FindOne(&tokenFilters.TokenFilter{
		Id: uint32(request.ID),
	})
	if err != nil {
		return nil, err
	}
	return &tokenResponses.GetTokenResponse{
		ID: uint32(token.ID),
	}, nil
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
		ChainID:     uint32(request.ChainID),
	}
	token, err := s.tokenRepository.Create(token)
	if err != nil {
		return nil, err
	}
	commonProducers.GetProducerHubInstance().InitTokenEvent(&commonProducersEvent.InitTokenEvent{
		ID:          uint32(token.ID),
		Address:     token.Address,
		Symbol:      token.Symbol,
		TokenName:   token.Name,
		Description: token.Description,
		TotalSupply: token.TotalSupply,
		ChainID:     uint32(token.ChainID),
	})
	return &tokenResponses.CreateToken{
		ID: token.ID,
	}, nil
}
