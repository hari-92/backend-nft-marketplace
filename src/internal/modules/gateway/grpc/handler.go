package gateway_grpc

import (
	pb "gitlab.com/hari-92/nft-market-server/pkg/grpc/proto_type"
)

type Handler struct {
	pb.UnimplementedGatewayProtoServiceServer
}

func NewGrpcHandler() *Handler {
	return &Handler{}
}
