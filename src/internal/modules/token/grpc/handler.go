package token_grpc

import (
	"context"
	"fmt"

	pb "gitlab.com/hari-92/nft-market-server/pkg/grpc/proto_type"
)

type Handler struct {
	pb.UnimplementedTokenProtoServiceServer
}

func NewHandler() *Handler {
	return &Handler{}
}

func (h *Handler) GetTokens(ctx context.Context, req *pb.GetTokensRequest) (*pb.GetTokensResponse, error) {
	fmt.Println("token_grpc GetTokens")
	return &pb.GetTokensResponse{}, nil
}

func (h *Handler) GetToken(ctx context.Context, req *pb.GetTokenRequest) (*pb.GetTokenResponse, error) {
	fmt.Println("token_grpc GetToken")
	return &pb.GetTokenResponse{}, nil
}

func (h *Handler) PostToken(ctx context.Context, req *pb.PostTokenRequest) (*pb.PostTokenResponse, error) {
	fmt.Println("token_grpc PostToken")
	return &pb.PostTokenResponse{}, nil
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
