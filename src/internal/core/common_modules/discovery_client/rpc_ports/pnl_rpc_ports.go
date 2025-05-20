package rpc_ports

import (
	"context"

	pb "gitlab.com/hari-92/nft-market-server/pkg/grpc/proto_type"
	"google.golang.org/grpc"
)

type IPnlRpcPorts interface {
	GetRealized(ctx context.Context, in *pb.GetRealizedRequest, opts ...grpc.CallOption) (*pb.GetRealizedResponse, error)
	GetUnrealized(ctx context.Context, in *pb.GetUnrealizedRequest, opts ...grpc.CallOption) (*pb.GetUnrealizedResponse, error)
	GetSummary(ctx context.Context, in *pb.GetSummaryRequest, opts ...grpc.CallOption) (*pb.GetSummaryResponse, error)
	GetPortfolio(ctx context.Context, in *pb.GetPortfolioRequest, opts ...grpc.CallOption) (*pb.GetPortfolioResponse, error)
	PostRecalculate(ctx context.Context, in *pb.PostRecalculateRequest, opts ...grpc.CallOption) (*pb.PostRecalculateResponse, error)
	GetHistory(ctx context.Context, in *pb.GetHistoryRequest, opts ...grpc.CallOption) (*pb.GetHistoryResponse, error)
	PostValidate(ctx context.Context, in *pb.PostValidateRequest, opts ...grpc.CallOption) (*pb.PostValidateResponse, error)
}

type pnlRpcPorts struct {
}

func NewPnlRpcPorts() IPnlRpcPorts {
	return &pnlRpcPorts{}
}

func (p *pnlRpcPorts) GetRealized(ctx context.Context, in *pb.GetRealizedRequest, opts ...grpc.CallOption) (*pb.GetRealizedResponse, error) {
	// TODO: Implement this
	return nil, nil
}

func (p *pnlRpcPorts) GetUnrealized(ctx context.Context, in *pb.GetUnrealizedRequest, opts ...grpc.CallOption) (*pb.GetUnrealizedResponse, error) {
	// TODO: Implement this
	return nil, nil
}

func (p *pnlRpcPorts) GetSummary(ctx context.Context, in *pb.GetSummaryRequest, opts ...grpc.CallOption) (*pb.GetSummaryResponse, error) {
	// TODO: Implement this
	return nil, nil
}

func (p *pnlRpcPorts) GetPortfolio(ctx context.Context, in *pb.GetPortfolioRequest, opts ...grpc.CallOption) (*pb.GetPortfolioResponse, error) {
	// TODO: Implement this
	return nil, nil
}

func (p *pnlRpcPorts) PostRecalculate(ctx context.Context, in *pb.PostRecalculateRequest, opts ...grpc.CallOption) (*pb.PostRecalculateResponse, error) {
	// TODO: Implement this
	return nil, nil
}

func (p *pnlRpcPorts) GetHistory(ctx context.Context, in *pb.GetHistoryRequest, opts ...grpc.CallOption) (*pb.GetHistoryResponse, error) {
	// TODO: Implement this
	return nil, nil
}

func (p *pnlRpcPorts) PostValidate(ctx context.Context, in *pb.PostValidateRequest, opts ...grpc.CallOption) (*pb.PostValidateResponse, error) {
	// TODO: Implement this
	return nil, nil
}

func (p *pnlRpcPorts) GetPnlPair(ctx context.Context, in *pb.GetPnlPairRequest, opts ...grpc.CallOption) (*pb.GetPnlPairResponse, error) {
	// TODO: Implement this
	return nil, nil
}

func (p *pnlRpcPorts) GetExport(ctx context.Context, in *pb.GetExportRequest, opts ...grpc.CallOption) (*pb.GetExportResponse, error) {
	// TODO: Implement this
	return nil, nil
}
