package pair_grpc

import (
	"fmt"

	"gitlab.com/hari-92/nft-market-server/pkg/bootstrap"
	pb "gitlab.com/hari-92/nft-market-server/pkg/grpc/proto_type"
	"google.golang.org/grpc"
)

type Server struct {
	GrpcServer *bootstrap.GrpcServer
	grpcServer *grpc.Server
}

func NewGrpcServer(grpcServer *bootstrap.GrpcServer, grpcHandler *Handler) *Server {
	fmt.Println("NewGrpcServer pair_grpc")
	pb.RegisterPairProtoServiceServer(grpcServer.Server, grpcHandler)

	return &Server{
		GrpcServer: grpcServer,
		grpcServer: grpcServer.Server,
	}
}
