package rpc_ports

import (
	"context"

	discovery_client "gitlab.com/hari-92/nft-market-server/internal/core/common_modules/discovery_client"
	pb "gitlab.com/hari-92/nft-market-server/pkg/grpc/proto_type"
	"google.golang.org/grpc"
)

type IPnlRpcPorts interface {
	GetRealized(ctx context.Context, in *pb.GetRealizedRequest, opts ...grpc.CallOption) (*pb.GetRealizedResponse, error)
	GetUnrealized(ctx context.Context, in *pb.GetUnrealizedRequest, opts ...grpc.CallOption) (*pb.GetUnrealizedResponse, error)
	GetSummary(ctx context.Context, in *pb.GetSummaryRequest, opts ...grpc.CallOption) (*pb.GetSummaryResponse, error)
	GetPortfolio(ctx context.Context, in *pb.GetPortfolioRequest, opts ...grpc.CallOption) (*pb.GetPortfolioResponse, error)
	PostRecalculate(ctx context.Context, in *pb.PostRecalculateRequest, opts ...grpc.CallOption) (*pb.PostRecalculateResponse, error)
	GetHistory(ctx context.Context, in *pb.GetPnlHistoryRequest, opts ...grpc.CallOption) (*pb.GetPnlHistoryResponse, error)
	PostValidate(ctx context.Context, in *pb.PostValidateRequest, opts ...grpc.CallOption) (*pb.PostValidateResponse, error)
	GetPnlPair(ctx context.Context, in *pb.GetPnlPairRequest, opts ...grpc.CallOption) (*pb.GetPnlPairResponse, error)
	GetExport(ctx context.Context, in *pb.GetExportRequest, opts ...grpc.CallOption) (*pb.GetExportResponse, error)
}

type pnlRpcPorts struct {
}

func NewPnlRpcPorts() IPnlRpcPorts {
	return &pnlRpcPorts{}
}

func (p *pnlRpcPorts) GetRealized(ctx context.Context, in *pb.GetRealizedRequest, opts ...grpc.CallOption) (*pb.GetRealizedResponse, error) {
	res, err := discovery_client.GetPnlClient().GetRealized(ctx, in, opts...)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (p *pnlRpcPorts) GetUnrealized(ctx context.Context, in *pb.GetUnrealizedRequest, opts ...grpc.CallOption) (*pb.GetUnrealizedResponse, error) {
	res, err := discovery_client.GetPnlClient().GetUnrealized(ctx, in, opts...)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (p *pnlRpcPorts) GetSummary(ctx context.Context, in *pb.GetSummaryRequest, opts ...grpc.CallOption) (*pb.GetSummaryResponse, error) {
	res, err := discovery_client.GetPnlClient().GetSummary(ctx, in, opts...)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (p *pnlRpcPorts) GetPortfolio(ctx context.Context, in *pb.GetPortfolioRequest, opts ...grpc.CallOption) (*pb.GetPortfolioResponse, error) {
	res, err := discovery_client.GetPnlClient().GetPortfolio(ctx, in, opts...)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (p *pnlRpcPorts) PostRecalculate(ctx context.Context, in *pb.PostRecalculateRequest, opts ...grpc.CallOption) (*pb.PostRecalculateResponse, error) {
	res, err := discovery_client.GetPnlClient().PostRecalculate(ctx, in, opts...)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (p *pnlRpcPorts) GetHistory(ctx context.Context, in *pb.GetPnlHistoryRequest, opts ...grpc.CallOption) (*pb.GetPnlHistoryResponse, error) {
	res, err := discovery_client.GetPnlClient().GetHistory(ctx, in, opts...)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (p *pnlRpcPorts) PostValidate(ctx context.Context, in *pb.PostValidateRequest, opts ...grpc.CallOption) (*pb.PostValidateResponse, error) {
	res, err := discovery_client.GetPnlClient().PostValidate(ctx, in, opts...)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (p *pnlRpcPorts) GetPnlPair(ctx context.Context, in *pb.GetPnlPairRequest, opts ...grpc.CallOption) (*pb.GetPnlPairResponse, error) {
	res, err := discovery_client.GetPnlClient().GetPnlPair(ctx, in, opts...)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (p *pnlRpcPorts) GetExport(ctx context.Context, in *pb.GetExportRequest, opts ...grpc.CallOption) (*pb.GetExportResponse, error) {
	res, err := discovery_client.GetPnlClient().GetExport(ctx, in, opts...)
	if err != nil {
		return nil, err
	}
	return res, nil
}
