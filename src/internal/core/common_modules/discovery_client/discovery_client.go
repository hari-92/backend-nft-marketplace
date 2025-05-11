package discovery_client

import (
	"context"
	"fmt"

	discoveryClientProperties "gitlab.com/hari-92/nft-market-server/internal/core/common_modules/discovery_client/properties"
	"gitlab.com/hari-92/nft-market-server/internal/core/constant"
	pb "gitlab.com/hari-92/nft-market-server/pkg/grpc/proto_type"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type IDiscoveryGrpcClient interface {
	GetClient(serviceID constant.ServiceID) (interface{}, error)
	Discovery(serviceID constant.ServiceID) (host string, port uint64, err error)
}

type DiscoveryGrpcClient struct {
	client        pb.DiscoveryProtoServiceClient
	clientManager map[constant.ServiceID]interface{}
}

func NewDiscoveryClient(properties *discoveryClientProperties.GrpcServerDiscoveryProperties) IDiscoveryGrpcClient {
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

	if constant.RegisterGrpcClient[serviceID] == nil {
		return nil, fmt.Errorf("service %s not found", serviceID)
	}

	d.clientManager[serviceID] = constant.RegisterGrpcClient[serviceID](conn)
	return d.clientManager[serviceID], nil
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

func (d *DiscoveryGrpcClient) Discovery(serviceID constant.ServiceID) (host string, port uint64, err error) {
	fmt.Println("Discovering service...", serviceID, port)
	discovery, err := d.client.Discovery(context.Background(), &pb.DiscoveryServiceRequest{
		ServiceID: serviceID.String(),
	})
	if err != nil {
		return "", 0, err
	}

	return discovery.Host, discovery.Port, nil
}
