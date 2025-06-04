package token_services

import (
	commonProducers "gitlab.com/hari-92/nft-market-server/internal/core/common_modules/producers"
	commonProducersEvent "gitlab.com/hari-92/nft-market-server/internal/core/common_modules/producers/events/token"
	tokenErrors "gitlab.com/hari-92/nft-market-server/internal/modules/token/errors"
	tokenFilters "gitlab.com/hari-92/nft-market-server/internal/modules/token/filters"
	tokenModels "gitlab.com/hari-92/nft-market-server/internal/modules/token/models"
	tokenRepositories "gitlab.com/hari-92/nft-market-server/internal/modules/token/repositories"
	tokenRequests "gitlab.com/hari-92/nft-market-server/internal/modules/token/requests"
	tokenResponses "gitlab.com/hari-92/nft-market-server/internal/modules/token/responses"
)

type ITokenService interface {
	GetToken(request *tokenRequests.GetTokenRequest) (*tokenResponses.GetTokenResponse, error)
	GetTokens(request *tokenRequests.GetTokensRequest) ([]*tokenResponses.GetTokensResponse, error)
	CreateToken(request *tokenRequests.CreateTokenRequest) (*tokenResponses.CreateTokenResponse, error)
	UpdateTokenByID(id uint32, request *tokenRequests.UpdateTokenRequest) (*tokenResponses.UpdateTokenResponse, error)
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
		ID:          uint32(token.ID),
		Address:     token.Address,
		Symbol:      token.Symbol,
		Decimals:    int32(token.Decimals),
		TotalSupply: token.TotalSupply,
		Name:        token.Name,
		Description: token.Description,
		ChainID:     token.ChainID,
	}, nil
}

func (s *tokenService) GetTokens(request *tokenRequests.GetTokensRequest) ([]*tokenResponses.GetTokensResponse, error) {
	tokens, err := s.tokenRepository.FindMany(&tokenFilters.TokenFilter{
		ChainID: request.ChainID,
	})
	if err != nil {
		return nil, err
	}

	responseTokens := make([]*tokenResponses.GetTokensResponse, len(tokens))
	for i, token := range tokens {
		responseTokens[i] = &tokenResponses.GetTokensResponse{
			ID:          uint32(token.ID),
			Address:     token.Address,
			Symbol:      token.Symbol,
			Decimals:    int32(token.Decimals),
			TotalSupply: token.TotalSupply,
			Name:        token.Name,
			Description: token.Description,
			ChainID:     token.ChainID,
		}
	}
	return responseTokens, nil
}

func (s *tokenService) CreateToken(request *tokenRequests.CreateTokenRequest) (*tokenResponses.CreateTokenResponse, error) {
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
	return &tokenResponses.CreateTokenResponse{
		ID: token.ID,
	}, nil
}

func (s *tokenService) UpdateTokenByID(id uint32, request *tokenRequests.UpdateTokenRequest) (*tokenResponses.UpdateTokenResponse, error) {
	model, err := s.tokenRepository.FindOne(&tokenFilters.TokenFilter{
		Id: id,
	})
	if err != nil {
		return nil, err
	}
	if model == nil {
		return nil, tokenErrors.TokenErrorNotFound
	}

	if request.Address != nil {
		model.Address = *request.Address
	}
	if request.Name != nil {
		model.Name = *request.Name
	}
	if request.Symbol != nil {
		model.Symbol = *request.Symbol
	}
	if request.Description != nil {
		model.Description = *request.Description
	}
	if request.ChainID != nil {
		model.ChainID = *request.ChainID
	}
	token, err := s.tokenRepository.Save(model)

	if err != nil {
		return nil, err
	}

	return &tokenResponses.UpdateTokenResponse{
		ID: uint32(token.ID),
	}, nil
}
