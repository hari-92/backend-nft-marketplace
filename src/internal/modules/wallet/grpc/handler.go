package wallet_grpc

import (
	"context"

	walletConverters "gitlab.com/hari-92/nft-market-server/internal/modules/wallet/converters"
	walletServices "gitlab.com/hari-92/nft-market-server/internal/modules/wallet/services"
	pb "gitlab.com/hari-92/nft-market-server/pkg/grpc/proto_type"
)

type Handler struct {
	pb.UnimplementedWalletProtoServiceServer
	walletService walletServices.IWalletService
}

func NewGrpcHandler(walletService walletServices.IWalletService) *Handler {
	return &Handler{
		walletService: walletService,
	}
}

func (h *Handler) GetBalance(ctx context.Context, req *pb.GetBalanceRequest) (*pb.GetBalanceResponse, error) {
	wallets, err := h.walletService.GetBalance(req.UserId)
	if err != nil {
		return nil, err
	}
	return &pb.GetBalanceResponse{
		Wallets: walletConverters.ToWalletProtoList(wallets),
	}, nil
}
