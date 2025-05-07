package constant

import (
	"google.golang.org/grpc"

	pb "gitlab.com/hari-92/nft-market-server/pkg/grpc/proto_type"
)

type ServiceID string

const (
	User ServiceID = "user"
)

func (s ServiceID) String() string {
	return string(s)
}

func UserClient(cc grpc.ClientConnInterface) interface{} {
	return pb.NewUserProtoServiceClient(cc)
}

var RegisterGrpcClient = map[ServiceID]func(cc grpc.ClientConnInterface) interface{}{
	User: UserClient,
}
