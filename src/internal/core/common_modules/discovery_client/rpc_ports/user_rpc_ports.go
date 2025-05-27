package rpc_ports

import (
	"context"

	commonDto "gitlab.com/hari-92/nft-market-server/internal/core/common_dto"
	"gitlab.com/hari-92/nft-market-server/internal/core/common_modules/discovery_client"
	pb "gitlab.com/hari-92/nft-market-server/pkg/grpc/proto_type"
)

type IUserRpcPorts interface {
	HelloWorld(s string) (string, error)
	IsExistUser(username string) (bool, error)
	CreateUser(user *commonDto.CreateUserPayload) (*commonDto.CreateUserResponse, error)
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

func (u *userRpcPorts) IsExistUser(username string) (bool, error) {
	res, err := discovery_client.GetUserClient().IsExistUser(
		context.Background(),
		&pb.IsExistUserRequest{
			Username: username,
		},
	)
	if err != nil {
		return false, err
	}
	if res == nil {
		return false, nil
	}
	return res.GetIsExist(), nil
}

func (u *userRpcPorts) CreateUser(user *commonDto.CreateUserPayload) (*commonDto.CreateUserResponse, error) {
	res, err := discovery_client.GetUserClient().CreateUser(
		context.Background(),
		&pb.CreateUserRequest{
			Email:    user.Username,
			Password: user.Password,
		},
	)
	if err != nil {
		return nil, err
	}

	return &commonDto.CreateUserResponse{
		Id:    res.GetId(),
		Email: res.GetEmail(),
	}, nil
}
