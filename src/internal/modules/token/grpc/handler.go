package token_grpc

import (
	"context"
	"fmt"

	pb "gitlab.com/hari-92/nft-market-server/pkg/grpc/proto_type"

	tokenRequests "gitlab.com/hari-92/nft-market-server/internal/modules/token/requests"
	tokenServices "gitlab.com/hari-92/nft-market-server/internal/modules/token/services"
)

type Handler struct {
	pb.UnimplementedTokenProtoServiceServer
	tokenService tokenServices.ITokenService
}

func NewHandler(
	tokenService tokenServices.ITokenService,
) *Handler {
	return &Handler{
		tokenService: tokenService,
	}
}

func (h *Handler) GetTokens(ctx context.Context, req *pb.GetTokensRequest) (*pb.GetTokensResponse, error) {
	tokens, err := h.tokenService.GetTokens(&tokenRequests.GetTokensRequest{
		ChainID: req.GetChainId(),
	})
	if err != nil {
		return nil, err
	}
	if tokens == nil {
		return &pb.GetTokensResponse{}, nil
	}

	responseTokens := make([]*pb.GetTokensResponse_Token, len(tokens))
	for i, token := range tokens {
		responseTokens[i] = &pb.GetTokensResponse_Token{
			Id:          uint32(token.ID),
			Address:     token.Address,
			Symbol:      token.Symbol,
			Decimals:    token.Decimals,
			TotalSupply: token.TotalSupply,
			Name:        token.Name,
			Description: token.Description,
			ChainId:     token.ChainID,
		}
	}

	return &pb.GetTokensResponse{
		Tokens: responseTokens,
	}, nil
}

func (h *Handler) GetToken(ctx context.Context, req *pb.GetTokenRequest) (*pb.GetTokenResponse, error) {
	token, err := h.tokenService.GetToken(&tokenRequests.GetTokenRequest{
		ID: req.GetId(),
	})
	if err != nil {
		return nil, err
	}
	return &pb.GetTokenResponse{
		Id:          token.ID,
		Address:     token.Address,
		Symbol:      token.Symbol,
		Decimals:    int32(token.Decimals),
		TotalSupply: token.TotalSupply,
		Name:        token.Name,
		Description: token.Description,
		ChainId:     token.ChainID,
	}, nil
}

func (h *Handler) PostToken(ctx context.Context, req *pb.PostTokenRequest) (*pb.PostTokenResponse, error) {
	token, err := h.tokenService.CreateToken(&tokenRequests.CreateTokenRequest{
		Address:     req.Address,
		Symbol:      req.Symbol,
		Name:        req.Name,
		Description: req.Description,
		Decimals:    int(req.Decimals),
		TotalSupply: uint64(req.TotalSupply),
		ChainID:     uint32(req.ChainId),
	})
	if err != nil {
		return nil, err
	}
	return &pb.PostTokenResponse{
		Id: uint32(token.ID),
	}, nil
}

func (h *Handler) PutToken(ctx context.Context, req *pb.PutTokenRequest) (*pb.PutTokenResponse, error) {
	token, err := h.tokenService.UpdateTokenByID(
		req.GetId(),
		&tokenRequests.UpdateTokenRequest{
			Address:     req.Address,
			Symbol:      req.Symbol,
			Name:        req.Name,
			Description: req.Description,
			ChainID:     req.ChainId,
		},
	)
	if err != nil {
		return nil, err
	}
	return &pb.PutTokenResponse{
		Id: uint32(token.ID),
	}, nil
}

func (h *Handler) DeleteToken(ctx context.Context, req *pb.DeleteTokenRequest) (*pb.DeleteTokenResponse, error) {
	fmt.Println("token_grpc DeleteToken")
	return &pb.DeleteTokenResponse{}, nil
}

func (h *Handler) PostValidateToken(ctx context.Context, req *pb.PostValidateTokenRequest) (*pb.PostValidateTokenResponse, error) {
	fmt.Println("token_grpc PostValidateToken")
	return &pb.PostValidateTokenResponse{}, nil
}
