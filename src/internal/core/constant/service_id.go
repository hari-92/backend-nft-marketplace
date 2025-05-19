package constant

import (
	"google.golang.org/grpc"

	pb "gitlab.com/hari-92/nft-market-server/pkg/grpc/proto_type"
)

type ServiceID string

const (
	User   ServiceID = "user"
	Wallet ServiceID = "wallet"
	Token  ServiceID = "token"
)

func (s ServiceID) String() string {
	return string(s)
}

func UserClient(cc grpc.ClientConnInterface) interface{} {
	return pb.NewUserProtoServiceClient(cc)
}

func WalletClient(cc grpc.ClientConnInterface) interface{} {
	return pb.NewWalletProtoServiceClient(cc)
}

func TokenClient(cc grpc.ClientConnInterface) interface{} {
	return pb.NewTokenProtoServiceClient(cc)
}

var RegisterGrpcClient = map[ServiceID]func(cc grpc.ClientConnInterface) interface{}{
	User:   UserClient,
	Wallet: WalletClient,
	Token:  TokenClient,
}
