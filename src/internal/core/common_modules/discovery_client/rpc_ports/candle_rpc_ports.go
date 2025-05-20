package rpc_ports

import (
	"context"

	pb "gitlab.com/hari-92/nft-market-server/pkg/grpc/proto_type"
)

type ICandleRpcPorts interface {
	GetCandles(ctx context.Context, request *pb.GetCandlesRequest) (*pb.GetCandlesResponse, error)
	GetCandleStats(ctx context.Context, request *pb.GetCandleStatsRequest) (*pb.GetCandleStatsResponse, error)
	GetCandleLatest(ctx context.Context, request *pb.GetCandleLatestRequest) (*pb.GetCandleLatestResponse, error)
	PostCandleValidate(ctx context.Context, request *pb.PostCandleValidateRequest) (*pb.PostCandleValidateResponse, error)
	GetCandlesAggregate(ctx context.Context, request *pb.GetCandlesAggregateRequest) (*pb.GetCandlesAggregateResponse, error)
}

type candleRpcPorts struct {
}

func NewCandleRpcPorts() ICandleRpcPorts {
	return &candleRpcPorts{}
}

func (p *candleRpcPorts) GetCandles(ctx context.Context, request *pb.GetCandlesRequest) (*pb.GetCandlesResponse, error) {
	// TODO: Implement this
	return nil, nil
}

func (p *candleRpcPorts) GetCandleStats(ctx context.Context, request *pb.GetCandleStatsRequest) (*pb.GetCandleStatsResponse, error) {
	// TODO: Implement this
	return nil, nil
}

func (p *candleRpcPorts) GetCandleLatest(ctx context.Context, request *pb.GetCandleLatestRequest) (*pb.GetCandleLatestResponse, error) {
	// TODO: Implement this
	return nil, nil
}

func (p *candleRpcPorts) PostCandleValidate(ctx context.Context, request *pb.PostCandleValidateRequest) (*pb.PostCandleValidateResponse, error) {
	// TODO: Implement this
	return nil, nil
}

func (p *candleRpcPorts) GetCandlesAggregate(ctx context.Context, request *pb.GetCandlesAggregateRequest) (*pb.GetCandlesAggregateResponse, error) {
	// TODO: Implement this
	return nil, nil
}
