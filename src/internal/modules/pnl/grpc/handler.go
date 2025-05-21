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

func (h *Handler) GetExport(ctx context.Context, in *pb.GetExportRequest) (*pb.GetExportResponse, error) {
	fmt.Println("pnl_grpc GetExport")
	return nil, nil
}
