package rpc_ports

import (
	"context"

	discoveryClient "gitlab.com/hari-92/nft-market-server/internal/core/common_modules/discovery_client"
	pb "gitlab.com/hari-92/nft-market-server/pkg/grpc/proto_type"
)

type ITokenRpcPorts interface {
	GetTokens(ctx context.Context, request *pb.GetTokensRequest) (*pb.GetTokensResponse, error)
	GetToken(ctx context.Context, request *pb.GetTokenRequest) (*pb.GetTokenResponse, error)
	PostToken(ctx context.Context, request *pb.PostTokenRequest) (*pb.PostTokenResponse, error)
	PutToken(ctx context.Context, request *pb.PutTokenRequest) (*pb.PutTokenResponse, error)
	DeleteToken(ctx context.Context, request *pb.DeleteTokenRequest) (*pb.DeleteTokenResponse, error)
	PostValidateToken(ctx context.Context, request *pb.PostValidateTokenRequest) (*pb.PostValidateTokenResponse, error)
}

type tokenRpcPorts struct {
}

func NewTokenRpcPorts() ITokenRpcPorts {
	return &tokenRpcPorts{}
}

func (t *tokenRpcPorts) GetTokens(ctx context.Context, request *pb.GetTokensRequest) (*pb.GetTokensResponse, error) {
	res, err := discoveryClient.GetTokenClient().GetTokens(ctx, request)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (t *tokenRpcPorts) GetToken(ctx context.Context, request *pb.GetTokenRequest) (*pb.GetTokenResponse, error) {
	res, err := discoveryClient.GetTokenClient().GetToken(ctx, request)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (t *tokenRpcPorts) PostToken(ctx context.Context, request *pb.PostTokenRequest) (*pb.PostTokenResponse, error) {
	res, err := discoveryClient.GetTokenClient().PostToken(ctx, request)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (t *tokenRpcPorts) PutToken(ctx context.Context, request *pb.PutTokenRequest) (*pb.PutTokenResponse, error) {
	res, err := discoveryClient.GetTokenClient().PutToken(ctx, request)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (t *tokenRpcPorts) DeleteToken(ctx context.Context, request *pb.DeleteTokenRequest) (*pb.DeleteTokenResponse, error) {
	res, err := discoveryClient.GetTokenClient().DeleteToken(ctx, request)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (t *tokenRpcPorts) PostValidateToken(ctx context.Context, request *pb.PostValidateTokenRequest) (*pb.PostValidateTokenResponse, error) {
	res, err := discoveryClient.GetTokenClient().PostValidateToken(ctx, request)
	if err != nil {
		return nil, err
	}
	return res, nil
}
