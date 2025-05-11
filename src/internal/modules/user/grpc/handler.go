package user_grpc

import (
	"context"

	pb "gitlab.com/hari-92/nft-market-server/pkg/grpc/proto_type"
)

type Handler struct {
	pb.UnimplementedUserProtoServiceServer
}

func NewGrpcHandler() *Handler {
	return &Handler{}
}

func (h *Handler) HelloWorld(ctx context.Context, req *pb.HelloWorldRequest) (*pb.HelloWorldResponse, error) {
	return &pb.HelloWorldResponse{
		Message: "Hello world from user service",
	}, nil
}
