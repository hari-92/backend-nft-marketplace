package user_grpc

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

func NewGrpcServer(g *bootstrap.GrpcServer, handler *Handler) *Server {
	fmt.Println("NewGrpcServer")
	pb.RegisterUserProtoServiceServer(g.Server, handler)

	return &Server{
		GrpcServer: g,
		grpcServer: g.Server,
	}
}
