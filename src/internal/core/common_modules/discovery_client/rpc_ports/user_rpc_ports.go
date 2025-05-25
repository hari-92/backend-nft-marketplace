package rpc_ports

import (
	"context"

	"gitlab.com/hari-92/nft-market-server/internal/core/common_modules/discovery_client"
	pb "gitlab.com/hari-92/nft-market-server/pkg/grpc/proto_type"
)

type IUserRpcPorts interface {
	HelloWorld(s string) (string, error)
	GetOneUser(ctx context.Context, request *pb.GetOneUserRequest) (*pb.GetOneUserResponse, error)
}

type userRpcPorts struct {
}

func NewUserRpcPorts() IUserRpcPorts {
	return &userRpcPorts{}
}

func (u *userRpcPorts) HelloWorld(s string) (string, error) {
	res, err := discovery_client.GetUserClient().HelloWorld(context.Background(), &pb.HelloWorldRequest{})
	if err != nil {
		return "", err
	}
	return res.GetMessage(), nil
}

func (u *userRpcPorts) GetOneUser(ctx context.Context, request *pb.GetOneUserRequest) (*pb.GetOneUserResponse, error) {
	res, err := discovery_client.GetUserClient().GetOneUser(ctx, request)
	if err != nil {
		return nil, err
	}
	return res, nil
}
