package rpc_ports

import (
	"context"

	pb "gitlab.com/hari-92/nft-market-server/pkg/grpc/proto_type"
)

type IPairRpcPorts interface {
	GetPairs(ctx context.Context, request *pb.GetPairsRequest) (*pb.GetPairsResponse, error)
	GetPair(ctx context.Context, request *pb.GetPairRequest) (*pb.GetPairResponse, error)
	PostPair(ctx context.Context, request *pb.PostPairRequest) (*pb.PostPairResponse, error)
	PutPair(ctx context.Context, request *pb.PutPairRequest) (*pb.PutPairResponse, error)
	DeletePair(ctx context.Context, request *pb.DeletePairRequest) (*pb.DeletePairResponse, error)
	PostValidatePair(ctx context.Context, request *pb.PostValidatePairRequest) (*pb.PostValidatePairResponse, error)
}

type pairRpcPorts struct {
}

func NewPairRpcPorts() IPairRpcPorts {
	return &pairRpcPorts{}
}

func (p *pairRpcPorts) GetPairs(ctx context.Context, request *pb.GetPairsRequest) (*pb.GetPairsResponse, error) {
	// TODO: Implement this
	return nil, nil
}

func (p *pairRpcPorts) GetPair(ctx context.Context, request *pb.GetPairRequest) (*pb.GetPairResponse, error) {
	// TODO: Implement this
	return nil, nil
}

func (p *pairRpcPorts) PostPair(ctx context.Context, request *pb.PostPairRequest) (*pb.PostPairResponse, error) {
	// TODO: Implement this
	return nil, nil
}

func (p *pairRpcPorts) PutPair(ctx context.Context, request *pb.PutPairRequest) (*pb.PutPairResponse, error) {
	// TODO: Implement this
	return nil, nil
}

func (p *pairRpcPorts) DeletePair(ctx context.Context, request *pb.DeletePairRequest) (*pb.DeletePairResponse, error) {
	// TODO: Implement this
	return nil, nil
}

func (p *pairRpcPorts) PostValidatePair(ctx context.Context, request *pb.PostValidatePairRequest) (*pb.PostValidatePairResponse, error) {
	// TODO: Implement this
	return nil, nil
}
