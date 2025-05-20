package rpc_ports

import (
	"context"

	pb "gitlab.com/hari-92/nft-market-server/pkg/grpc/proto_type"
)

type IOrderRpcPorts interface {
	GetOrders(ctx context.Context, request *pb.GetOrdersRequest) (*pb.GetOrdersResponse, error)
	GetOrder(ctx context.Context, request *pb.GetOrderRequest) (*pb.GetOrderResponse, error)
	PostOrder(ctx context.Context, request *pb.PostOrderRequest) (*pb.PostOrderResponse, error)
	DeleteOrder(ctx context.Context, request *pb.DeleteOrderRequest) (*pb.DeleteOrderResponse, error)
	DeleteOrders(ctx context.Context, request *pb.DeleteOrdersRequest) (*pb.DeleteOrdersResponse, error)
	PostValidateOrder(ctx context.Context, request *pb.PostValidateOrderRequest) (*pb.PostValidateOrderResponse, error)
	GetHistory(ctx context.Context, request *pb.GetHistoryRequest) (*pb.GetHistoryResponse, error)
	PostAdminCancel(ctx context.Context, request *pb.PostAdminCancelRequest) (*pb.PostAdminCancelResponse, error)
}

type orderRpcPorts struct {
}

func NewOrderRpcPorts() IOrderRpcPorts {
	return &orderRpcPorts{}
}

func (p *orderRpcPorts) GetOrders(ctx context.Context, request *pb.GetOrdersRequest) (*pb.GetOrdersResponse, error) {
	// TODO: Implement this
	return nil, nil
}

func (p *orderRpcPorts) GetOrder(ctx context.Context, request *pb.GetOrderRequest) (*pb.GetOrderResponse, error) {
	// TODO: Implement this
	return nil, nil
}

func (p *orderRpcPorts) PostOrder(ctx context.Context, request *pb.PostOrderRequest) (*pb.PostOrderResponse, error) {
	// TODO: Implement this
	return nil, nil
}

func (p *orderRpcPorts) DeleteOrder(ctx context.Context, request *pb.DeleteOrderRequest) (*pb.DeleteOrderResponse, error) {
	// TODO: Implement this
	return nil, nil
}

func (p *orderRpcPorts) DeleteOrders(ctx context.Context, request *pb.DeleteOrdersRequest) (*pb.DeleteOrdersResponse, error) {
	// TODO: Implement this
	return nil, nil
}

func (p *orderRpcPorts) PostValidateOrder(ctx context.Context, request *pb.PostValidateOrderRequest) (*pb.PostValidateOrderResponse, error) {
	// TODO: Implement this
	return nil, nil
}

func (p *orderRpcPorts) GetHistory(ctx context.Context, request *pb.GetHistoryRequest) (*pb.GetHistoryResponse, error) {
	// TODO: Implement this
	return nil, nil
}

func (p *orderRpcPorts) PostAdminCancel(ctx context.Context, request *pb.PostAdminCancelRequest) (*pb.PostAdminCancelResponse, error) {
	// TODO: Implement this
	return nil, nil
}
