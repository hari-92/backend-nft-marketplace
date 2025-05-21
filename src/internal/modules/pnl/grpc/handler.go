package pnl_grpc

import (
	"context"
	"fmt"

	pb "gitlab.com/hari-92/nft-market-server/pkg/grpc/proto_type"
)

type Handler struct {
	pb.UnimplementedPnlProtoServiceServer
}

func NewGrpcHandler() *Handler {
	return &Handler{}
}

func (h *Handler) GetRealized(ctx context.Context, in *pb.GetRealizedRequest) (*pb.GetRealizedResponse, error) {
	fmt.Println("pnl_grpc GetRealized")
	return nil, nil
}

func (h *Handler) GetUnrealized(ctx context.Context, in *pb.GetUnrealizedRequest) (*pb.GetUnrealizedResponse, error) {
	fmt.Println("pnl_grpc GetUnrealized")
	return nil, nil
}

func (h *Handler) GetSummary(ctx context.Context, in *pb.GetSummaryRequest) (*pb.GetSummaryResponse, error) {
	fmt.Println("pnl_grpc GetSummary")
	return nil, nil
}

func (h *Handler) GetPortfolio(ctx context.Context, in *pb.GetPortfolioRequest) (*pb.GetPortfolioResponse, error) {
	fmt.Println("pnl_grpc GetPortfolio")
	return nil, nil
}

func (h *Handler) PostRecalculate(ctx context.Context, in *pb.PostRecalculateRequest) (*pb.PostRecalculateResponse, error) {
	fmt.Println("pnl_grpc PostRecalculate")
	return nil, nil
}

func (h *Handler) GetPnlHistory(ctx context.Context, in *pb.GetPnlHistoryRequest) (*pb.GetPnlHistoryResponse, error) {
	fmt.Println("pnl_grpc GetPnlHistory")
	return nil, nil
}

func (h *Handler) PostValidate(ctx context.Context, in *pb.PostValidateRequest) (*pb.PostValidateResponse, error) {
	fmt.Println("pnl_grpc PostValidate")
	return nil, nil
}

func (h *Handler) GetPnlPair(ctx context.Context, in *pb.GetPnlPairRequest) (*pb.GetPnlPairResponse, error) {
	fmt.Println("pnl_grpc GetPnlPair")
	return nil, nil
}

func (h *Handler) GetExport(ctx context.Context, in *pb.GetExportRequest) (*pb.GetExportResponse, error) {
	fmt.Println("pnl_grpc GetExport")
	return nil, nil
}
