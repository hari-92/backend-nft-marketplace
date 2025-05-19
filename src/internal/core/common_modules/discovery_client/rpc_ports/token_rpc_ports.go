package rpc_ports

import (
	"context"

	pb "gitlab.com/hari-92/nft-market-server/pkg/grpc/proto_type"
)

type ITokenRpcPorts interface {
	GetTokens(ctx context.Context, request *pb.GetTokensRequest) (*pb.GetTokensResponse, error)
}

type tokenRpcPorts struct {
}

func NewTokenRpcPorts() ITokenRpcPorts {
	return &tokenRpcPorts{}
}

func (t *tokenRpcPorts) GetTokens(ctx context.Context, request *pb.GetTokensRequest) (*pb.GetTokensResponse, error) {
	// TODO: Implement this
	return nil, nil
}

func (t *tokenRpcPorts) GetToken(ctx context.Context, request *pb.GetTokenRequest) (*pb.GetTokenResponse, error) {
	// TODO: Implement this
	return nil, nil
}

func (t *tokenRpcPorts) PostToken(ctx context.Context, request *pb.PostTokenRequest) (*pb.PostTokenResponse, error) {
	// TODO: Implement this
	return nil, nil
}

func (t *tokenRpcPorts) PutToken(ctx context.Context, request *pb.PutTokenRequest) (*pb.PutTokenResponse, error) {
	// TODO: Implement this
	return nil, nil
}

func (t *tokenRpcPorts) DeleteToken(ctx context.Context, request *pb.DeleteTokenRequest) (*pb.DeleteTokenResponse, error) {
	// TODO: Implement this
	return nil, nil
}

func (t *tokenRpcPorts) PostValidateToken(ctx context.Context, request *pb.PostValidateTokenRequest) (*pb.PostValidateTokenResponse, error) {
	// TODO: Implement this
	return nil, nil
}
