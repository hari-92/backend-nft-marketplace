package discovery_grpc

import (
	"context"
	"fmt"

	discoveryGrpc "gitlab.com/hari-92/nft-market-server/internal/modules/discovery/properties"
	pb "gitlab.com/hari-92/nft-market-server/pkg/grpc/proto_type"
)

type Handler struct {
	pb.UnimplementedDiscoveryProtoServiceServer
	serviceManager *discoveryGrpc.ServicesProperties
}

func NewGrpcHandler(serviceManager *discoveryGrpc.ServicesProperties) *Handler {
	return &Handler{
		serviceManager: serviceManager,
	}
}

func (h *Handler) Discovery(ctx context.Context, req *pb.DiscoveryServiceRequest) (*pb.DiscoveryServiceResponse, error) {
	service, ok := h.serviceManager.Services[req.GetServiceID()]
	if !ok {
		return nil, fmt.Errorf("service %s not found", req.GetServiceID())
	}

	return &pb.DiscoveryServiceResponse{
		Host: service.Host,
		Port: service.Port,
	}, nil
}
