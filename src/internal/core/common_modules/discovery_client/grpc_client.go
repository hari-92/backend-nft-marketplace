package discovery_client

import (
	"gitlab.com/hari-92/nft-market-server/internal/core/constant"
	pb "gitlab.com/hari-92/nft-market-server/pkg/grpc/proto_type"
)

var (
	userClient pb.UserProtoServiceClient
)

func GetUserClient() pb.UserProtoServiceClient {
	if userClient != nil {
		return userClient
	}

	client, err := discoveryGrpcClient.GetClient(constant.User)
	if err != nil {
		panic(err)
	}

	userClient = client.(pb.UserProtoServiceClient)
	return userClient
}
