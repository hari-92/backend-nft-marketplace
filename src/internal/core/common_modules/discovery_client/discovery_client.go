package discovery_client

import (
	"context"
	"fmt"
	"gitlab.com/hari-92/nft-market-server/internal/core/common_modules/discovery_client/properties"
	"gitlab.com/hari-92/nft-market-server/internal/core/constant"
	pb "gitlab.com/hari-92/nft-market-server/pkg/grpc/proto_type"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type IDiscoveryGrpcClient interface {
	GetClient(serviceID constant.ServiceID) (interface{}, error)
}

type DiscoveryGrpcClient struct {
	client        pb.DiscoveryProtoServiceClient
	clientManager map[constant.ServiceID]interface{}
}

func NewDiscoveryClient(properties *properties.GrpcServerDiscoveryProperties) IDiscoveryGrpcClient {
	var conn *grpc.ClientConn
	address := fmt.Sprintf("%s:%d", properties.Host, properties.Port)
	conn, err := grpc.NewClient(address, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		panic(err)
	}

	client := pb.NewDiscoveryProtoServiceClient(conn)
	return &DiscoveryGrpcClient{
		client:        client,
		clientManager: make(map[constant.ServiceID]interface{}),
	}
}

func (d *DiscoveryGrpcClient) registerClient(serviceID constant.ServiceID) (interface{}, error) {
	discovery, err := d.client.Discovery(context.Background(), &pb.DiscoveryServiceRequest{
		ServiceID: serviceID.String(),
	})
	if err != nil {
		return nil, err
	}

	address := fmt.Sprintf("%s:%d", discovery.Host, discovery.Port)
	conn, err := grpc.NewClient(address, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}

	d.clientManager[serviceID] = conn
	return conn, nil
}

func (d *DiscoveryGrpcClient) GetClient(serviceID constant.ServiceID) (interface{}, error) {
	client, ok := d.clientManager[serviceID]
	if ok {
		return client, nil
	}

	client, err := d.registerClient(serviceID)
	if err != nil {
		return nil, err
	}

	return client, nil
}
