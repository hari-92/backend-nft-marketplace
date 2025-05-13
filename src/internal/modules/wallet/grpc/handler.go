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

func (h *Handler) GetBalances(ctx context.Context, req *pb.GetBalancesRequest) (*pb.GetBalancesResponse, error) {
	wallets, err := h.walletService.GetBalances(req.UserId)
	if err != nil {
		return nil, err
	}
	return &pb.GetBalancesResponse{
		Wallets: walletConverters.ToWalletProtoList(wallets),
	}, nil
}
