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
	fmt.Println("token_grpc GetTokens")
	return &pb.GetTokensResponse{}, nil
}

func (h *Handler) GetToken(ctx context.Context, req *pb.GetTokenRequest) (*pb.GetTokenResponse, error) {
	token, err := h.tokenService.GetToken(&tokenRequests.GetTokenRequest{
		ID: req.TokenId,
	})
	if err != nil {
		return nil, err
	}
	return &pb.GetTokenResponse{
		TokenId: uint32(token.ID),
	}, nil
}

func (h *Handler) PostToken(ctx context.Context, req *pb.PostTokenRequest) (*pb.PostTokenResponse, error) {
	token, err := h.tokenService.CreateToken(&tokenRequests.CreateToken{
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
	fmt.Println("token_grpc PutToken")
	return &pb.PutTokenResponse{}, nil
}

func (h *Handler) DeleteToken(ctx context.Context, req *pb.DeleteTokenRequest) (*pb.DeleteTokenResponse, error) {
	fmt.Println("token_grpc DeleteToken")
	return &pb.DeleteTokenResponse{}, nil
}

func (h *Handler) PostValidateToken(ctx context.Context, req *pb.PostValidateTokenRequest) (*pb.PostValidateTokenResponse, error) {
	fmt.Println("token_grpc PostValidateToken")
	return &pb.PostValidateTokenResponse{}, nil
}
