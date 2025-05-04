package grpc

import (
	"gitlab.com/hari-92/nft-market-server/internal/modules/discovery/properties"
	pb "gitlab.com/hari-92/nft-market-server/pkg/grpc/proto_type"
)

type Handler struct {
	pb.UnimplementedDiscoveryProtoServiceServer
	serviceManager *properties.ServicesProperties
}

func NewGrpcHandler(serviceManager *properties.ServicesProperties) *Handler {
	return &Handler{
		serviceManager: serviceManager,
	}
}
